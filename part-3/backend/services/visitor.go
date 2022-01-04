package services

import (
	"context"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/config"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/models"
)

// VisitorService : struct
type VisitorService struct {
	conf *config.Config
}

// NewVisitor : conf
func NewVisitorService(conf *config.Config) *VisitorService {
	return &VisitorService{
		conf: conf,
	}
}

func (svc *VisitorService) GetCurrentVisitor(ctx context.Context, ip string) (*models.Visitor, error) {
	return nil, nil
}

func (svc *VisitorService) CreateNewVisitor(ctx context.Context, visitor *models.Visitor) (*models.Visitor, error) {
	return nil, nil
}

func (svc *VisitorService) GetLast100Visitors(ctx context.Context) ([]*models.Visitor, error) {
	return nil, nil
}

func (svc *VisitorService) GetTop100Visitors(ctx context.Context) ([]*models.Visitor, error) {
	return nil, nil
}
