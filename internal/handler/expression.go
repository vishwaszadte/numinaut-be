package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vishwaszadte/numinaut-be/internal/repository"
	"github.com/vishwaszadte/numinaut-be/internal/service"
)

// ExpressionHandler handles HTTP requests for expressions
type ExpressionHandler struct {
	service *service.ExpressionService
}

// NewExpressionHandler creates a new expression handler instance
func NewExpressionHandler(service *service.ExpressionService) *ExpressionHandler {
	return &ExpressionHandler{service: service}
}

// RegisterRoutes registers all expression routes to the router
func (h *ExpressionHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/expressions/{id:[0-9]+}", h.GetByID).Methods("GET")
	r.HandleFunc("/expressions/uuid/{uuid}", h.GetByUUID).Methods("GET")
	r.HandleFunc("/expressions", h.Filter).Methods("GET")
}

// GetByID handles requests to get an expression by ID
func (h *ExpressionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid expression ID")
		return
	}

	expr, err := h.service.GetByID(r.Context(), int32(id))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Expression not found")
		return
	}

	respondWithJSON(w, http.StatusOK, expr)
}

// GetByUUID handles requests to get an expression by UUID
func (h *ExpressionHandler) GetByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuidStr := vars["uuid"]

	uuidVal, err := uuid.Parse(uuidStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	expr, err := h.service.GetByUUID(r.Context(), uuidVal)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Expression not found")
		return
	}

	respondWithJSON(w, http.StatusOK, expr)
}

// Filter handles requests to filter expressions
func (h *ExpressionHandler) Filter(w http.ResponseWriter, r *http.Request) {
	var params repository.FilterExpressionsParams

	// Parse query parameters
	queryParams := r.URL.Query()

	// Set filter parameters based on query string
	if expr := queryParams.Get("expression"); expr != "" {
		params.ExpressionFilterOp = "="
		params.Expression = expr
	}

	if result := queryParams.Get("result"); result != "" {
		params.ResultFilterOp = "="
		if val, err := strconv.ParseFloat(result, 32); err == nil {
			params.Result = float32(val)
		}
	}

	if numOp := queryParams.Get("num_operands"); numOp != "" {
		params.NumOperandsFilterOp = "="
		if val, err := strconv.ParseInt(numOp, 10, 32); err == nil {
			params.NumOperands = int32(val)
		}
	}

	if diff := queryParams.Get("difficulty"); diff != "" {
		params.DifficultyFilterOp = "="
		if val, err := strconv.ParseInt(diff, 10, 32); err == nil {
			params.Difficulty = int32(val)
		}
	}

	// Set ordering parameters
	params.OrderBy = queryParams.Get("order_by")
	if params.OrderBy == "" {
		params.OrderBy = "id"
	}

	params.OrderDirection = queryParams.Get("order_direction")
	if params.OrderDirection == "" {
		params.OrderDirection = "asc"
	}

	if limit := queryParams.Get("limit"); limit != "" {
		if val, err := strconv.Atoi(limit); err == nil {
			params.Limit = pgtype.Int4{Int32: int32(val), Valid: true}
		}
	} else {
		params.Limit = pgtype.Int4{Int32: 10, Valid: true}
	}

	if offset := queryParams.Get("offset"); offset != "" {
		if val, err := strconv.Atoi(offset); err == nil {
			params.Offset = pgtype.Int4{Int32: int32(val), Valid: true}
		}
	} else {
		params.Offset = pgtype.Int4{Int32: 0, Valid: true}
	}

	expressions, err := h.service.Filter(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error filtering expressions")
		return
	}

	respondWithJSON(w, http.StatusOK, expressions)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling JSON"`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
