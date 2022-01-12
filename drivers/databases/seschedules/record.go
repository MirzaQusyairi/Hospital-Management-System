package seschedules

import (
	"Hospital-Management-System/business/seschedules"
	"time"

	"gorm.io/gorm"
)

type Sschedules struct {
	gorm.Model
	ID         int `gorm:"primary_key"`
	IDSchedule string
	IDDoctor   string
	IDFacilty  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func toDomain(ev Sschedules) seschedules.Domain {
	return seschedules.Domain{
		ID:         ev.ID,
		IDSchedule: ev.IDSchedule,
		IDDoctor:   ev.IDDoctor,
		IDFacilty:  ev.IDFacilty,
		CreatedAt:  ev.CreatedAt,
		UpdatedAt:  ev.UpdatedAt,
	}
}

func fromDomain(domain seschedules.Domain) Sschedules {
	return Sschedules{
		ID: domain.ID,

		IDSchedule: domain.IDSchedule,
		IDDoctor:   domain.IDDoctor,
		IDFacilty:  domain.IDFacilty,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Sschedules) seschedules.Domain {
	return seschedules.Domain{
		ID:         ev.ID,
		IDSchedule: ev.IDSchedule,
		IDDoctor:   ev.IDDoctor,
		IDFacilty:  ev.IDFacilty,
		CreatedAt:  ev.CreatedAt,
		UpdatedAt:  ev.UpdatedAt,
	}
}

func toDomainList(data []Sschedules) []seschedules.Domain {
	result := []seschedules.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
