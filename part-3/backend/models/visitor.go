package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Visitor struct
type Visitor struct {
	gorm.Model
	IpAddress string    `gorm:"unique_index;column:ip_address" json:"ip_address"`
	Location  string    `gorm:"column:location" json:"location"`
	Timezone  string    `gorm:"column:timezone" json:"timezone"`
	LastVisit time.Time `gorm:"index;column:last_visit" json:"last_visit"`
	Visits    uint      `gorm:"index;column:visits" json:"visits"`
}
