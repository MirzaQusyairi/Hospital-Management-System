package schedules

import (
	"Hospital-Management-System/business/schedules"
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	Day       string
	Start     string
	End       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(ev Schedules) schedules.Domain {
	return schedules.Domain{
		ID:        ev.ID,
		Day:       ev.Day,
		Start:     ev.Start,
		End:       ev.End,
		CreatedAt: ev.CreatedAt,
		UpdatedAt: ev.UpdatedAt,
	}
}

func fromDomain(domain schedules.Domain) Schedules {
	return Schedules{
		ID:        domain.ID,
		Day:       domain.Day,
		Start:     domain.Start,
		End:       domain.End,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Schedules) schedules.Domain {
	return schedules.Domain{
		ID:        ev.ID,
		Day:       ev.Day,
		Start:     ev.Start,
		End:       ev.End,
		CreatedAt: ev.CreatedAt,
		UpdatedAt: ev.UpdatedAt,
	}
}

func toDomainList(data []Schedules) []schedules.Domain {
	result := []schedules.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
