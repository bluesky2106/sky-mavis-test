package daos

import (
	"fmt"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/config"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/libs/mysql"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/models"
	"github.com/jinzhu/gorm"
)

var (
	tables = []interface{}{(*models.Visitor)(nil)}
	db     *gorm.DB
)

// GetDB : getter
func GetDB() *gorm.DB {
	return db
}

func emptyDBError() error {
	return fmt.Errorf("empty database")
}

// AutoMigrate :
func AutoMigrate() error {
	if db == nil {
		return emptyDBError()
	}

	return db.AutoMigrate(tables...).Error
}

// WithTransaction :
func WithTransaction(callback func(*gorm.DB) error) error {
	tx := db.Begin()

	if err := callback(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func where(db *gorm.DB, filters map[string]interface{}) *gorm.DB {
	query := db
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v)
		} else {
			query = query.Where(k)
		}
	}
	return query
}

// Init : initialize DB
func Init(conf *config.Config) error {
	var err error
	db, err = mysql.Init(conf.MySQLConnURL, conf.Env)
	if err != nil {
		return err
	}
	return nil
}
