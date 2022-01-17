package seschedules

import (
	"Hospital-Management-System/business/seschedules"
	"Hospital-Management-System/drivers/databases/doctors"
	"Hospital-Management-System/drivers/databases/facilities"
	"Hospital-Management-System/drivers/databases/schedules"
	"time"

	"gorm.io/gorm"
)

type Sschedules struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	SchedulesID int
	Schedules   schedules.Schedules `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DoctorsID   int
	Doctors     doctors.Doctors `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FacilityID  int
	Facility    facilities.Facilities `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(ev Sschedules) seschedules.Domain {
	return seschedules.Domain{
		ID:         ev.ID,
		IDSchedule: ev.SchedulesID,
		IDDoctor:   ev.DoctorsID,
		IDFacilty:  ev.FacilityID,
		CreatedAt:  ev.CreatedAt,
		UpdatedAt:  ev.UpdatedAt,
	}
}

func fromDomain(domain seschedules.Domain) Sschedules {
	return Sschedules{
		ID: domain.ID,

		SchedulesID: domain.IDSchedule,
		DoctorsID:   domain.IDDoctor,
		FacilityID:  domain.IDFacilty,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Sschedules) seschedules.Domain {
	return seschedules.Domain{
		ID:         ev.ID,
		IDSchedule: ev.SchedulesID,
		IDDoctor:   ev.DoctorsID,
		IDFacilty:  ev.FacilityID,
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
