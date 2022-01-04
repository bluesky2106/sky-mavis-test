package interfaces

import (
	"context"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/models"
)

// IVisitorService : interface
type IVisitorService interface {
	GetCurrentVisitor(ctx context.Context, ip string) (*models.Visitor, error)
	CreateNewVisitor(ctx context.Context, visitor *models.Visitor) (*models.Visitor, error)
	GetLast100Visitors(ctx context.Context) ([]*models.Visitor, error)
	GetTop100Visitors(ctx context.Context) ([]*models.Visitor, error)
}
