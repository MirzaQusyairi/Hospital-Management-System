package facilities

import (
	"Hospital-Management-System/business/facilities"
	"time"

	"gorm.io/gorm"
)

type Facilities struct {
	gorm.Model
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"unique"`
	Queue     int
	Location  string
	Capacity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(org Facilities) facilities.Domain {
	return facilities.Domain{
		ID:        org.ID,
		Name:      org.Name,
		Queue:     org.Queue,
		Location:  org.Location,
		Capacity:  org.Capacity,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}
}
func toDomainUpdate(upd Facilities) facilities.Domain {
	return facilities.Domain{
		ID:        upd.ID,
		Name:      upd.Name,
		Queue:     upd.Queue,
		Location:  upd.Location,
		Capacity:  upd.Capacity,
		CreatedAt: upd.CreatedAt,
		UpdatedAt: upd.UpdatedAt,
	}
}
func fromDomain(domain facilities.Domain) Facilities {
	return Facilities{
		ID:        domain.ID,
		Name:      domain.Name,
		Queue:     domain.Queue,
		Location:  domain.Location,
		Capacity:  domain.Capacity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func toDomainList(data []Facilities) []facilities.Domain {
	result := []facilities.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
