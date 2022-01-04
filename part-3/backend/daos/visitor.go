package daos

import (
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/models"
	"github.com/jinzhu/gorm"
)

const (
	visitorTable = "visitors"
)

// Visitor : struct
type Visitor struct {
}

// NewVisitor :
func NewVisitor() *Visitor {
	return &Visitor{}
}

// Create : tx, mod
func (dao *Visitor) Create(tx *gorm.DB, mod *models.Visitor) error {
	return tx.Create(mod).Error
}

// Update : tx, mod
func (dao *Visitor) Update(tx *gorm.DB, mod *models.Visitor) error {
	return tx.Save(mod).Error
}

// FindByID : id
func (dao *Visitor) FindByID(id uint64) (*models.Visitor, error) {
	var mod models.Visitor
	tx := GetDB().Table(visitorTable)
	if err := tx.Where("id = ?", id).First(&mod).Error; err != nil {
		return nil, err
	}
	return &mod, nil
}

// FindOneByQuery :
func (dao *Visitor) FindOneByQuery(filters map[string]interface{}) (*models.Visitor, error) {
	var mod models.Visitor

	query := GetDB().Table(visitorTable)
	query = where(query, filters)

	if err := query.First(&mod).Error; err != nil {
		return nil, err
	}
	return &mod, nil
}

// FindAllByQuery :
func (dao *Visitor) FindAllByQuery(filters map[string]interface{}) ([]*models.Visitor, error) {
	var (
		models []*models.Visitor
	)

	query := GetDB().Table(visitorTable)
	query = where(query, filters)

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
