package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vishwaszadte/numinaut-be/internal/repository"
)

// ExpressionService handles business logic for expressions
type ExpressionService struct {
	repo repository.ExpressionRepository
}

// NewExpressionService creates a new expression service instance
func NewExpressionService(repo repository.ExpressionRepository) *ExpressionService {
	return &ExpressionService{repo: repo}
}

// GetByID retrieves an expression by its ID
func (s *ExpressionService) GetByID(ctx context.Context, id int32) (*repository.Expression, error) {
	expr, err := s.repo.GetExpressionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &expr, nil
}

// GetByUUID retrieves an expression by its UUID
func (s *ExpressionService) GetByUUID(ctx context.Context, uuid uuid.UUID) (*repository.Expression, error) {
	expr, err := s.repo.GetExpressionByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return &expr, nil
}

// Filter retrieves expressions based on filter parameters
func (s *ExpressionService) Filter(ctx context.Context, params repository.FilterExpressionsParams) ([]repository.Expression, error) {
	return s.repo.FilterExpressions(ctx, params)
}
