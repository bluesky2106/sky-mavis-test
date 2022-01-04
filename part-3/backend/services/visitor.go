package services

import (
	"context"
	"net"
	"time"

	"github.com/bluesky2106/sky-mavis-test/part-3/backend/config"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/daos"
	"github.com/bluesky2106/sky-mavis-test/part-3/backend/models"
	"github.com/jinzhu/gorm"
	"github.com/oschwald/geoip2-golang"
)

// VisitorService : struct
type VisitorService struct {
	conf       *config.Config
	visitorDAO *daos.Visitor
}

// NewVisitor : conf
func NewVisitorService(conf *config.Config, visitorDAO *daos.Visitor) *VisitorService {
	return &VisitorService{
		conf:       conf,
		visitorDAO: visitorDAO,
	}
}

func (svc *VisitorService) GetCurrentVisitor(ctx context.Context, ip string) (*models.Visitor, error) {
	visitor, _ := svc.visitorDAO.FindOneByQuery(map[string]interface{}{
		"ip_address = ?": ip,
	})
	if visitor == nil {
		db, err := geoip2.Open("./resources/GeoLite2-City.mmdb")
		if err != nil {
			return nil, err
		}
		defer db.Close()

		ipAddr := net.ParseIP(ip)
		record, err := db.City(ipAddr)
		if err != nil {
			return nil, err
		}

		city := record.City.Names["en"]
		country := record.Country.Names["en"]
		location := country
		if city != "" {
			location = record.City.Names["en"] + "," + record.Country.Names["en"]
		}
		visitor = &models.Visitor{
			IpAddress: ip,
			Location:  location,
			Timezone:  record.Location.TimeZone,
			LastVisit: time.Now().UTC(),
			Visits:    1,
		}
		if err := daos.WithTransaction(func(tx *gorm.DB) error {
			return svc.visitorDAO.Create(tx, visitor)
		}); err != nil {
			return nil, err
		}
	} else {
		visitor.LastVisit = time.Now().UTC()
		visitor.Visits++
		if err := daos.WithTransaction(func(tx *gorm.DB) error {
			return svc.visitorDAO.Update(tx, visitor)
		}); err != nil {
			return nil, err
		}
	}

	return visitor, nil
}

func (svc *VisitorService) GetLast100Visitors(ctx context.Context) ([]*models.Visitor, error) {
	return svc.visitorDAO.FindLast100Visitors()
}

func (svc *VisitorService) GetTop100Visitors(ctx context.Context) ([]*models.Visitor, error) {
	return svc.visitorDAO.FindTop100Visitors()
}
