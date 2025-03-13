package repository

import (
	"context"

	"github.com/google/uuid"
)

// ExpressionRepository defines the interface for expression-related database operations
type ExpressionRepository interface {
	// GetExpressionByID retrieves an expression by its ID
	GetExpressionByID(ctx context.Context, id int32) (Expression, error)

	// GetExpressionByUUID retrieves an expression by its UUID
	GetExpressionByUUID(ctx context.Context, uuid uuid.UUID) (Expression, error)

	// FilterExpressions retrieves expressions based on filter parameters
	FilterExpressions(ctx context.Context, params FilterExpressionsParams) ([]Expression, error)
}

// Ensure Queries implements ExpressionRepository
var _ ExpressionRepository = (*Queries)(nil)