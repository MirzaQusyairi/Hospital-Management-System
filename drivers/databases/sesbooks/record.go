package sesbooks

import (
	"Hospital-Management-System/business/sesbooks"
	"Hospital-Management-System/drivers/databases/patients"
	"Hospital-Management-System/drivers/databases/seschedules"
	"time"

	"gorm.io/gorm"
)

type Sesbooks struct {
	gorm.Model
	ID                int `gorm:"primary_key"`
	PatientsID        int
	Patients          patients.Patients `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	SessionScheduleID int
	SessionSchedule   seschedules.Sschedules `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	PatientQueue      int
	Date              time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func toDomain(ev Sesbooks) sesbooks.Domain {
	return sesbooks.Domain{
		ID:                ev.ID,
		IDPatient:         ev.PatientsID,
		IDSessionSchedule: ev.SessionScheduleID,
		PatientQueue:      ev.PatientQueue,
		Date:              ev.Date,
		CreatedAt:         ev.CreatedAt,
		UpdatedAt:         ev.UpdatedAt,
	}
}

func fromDomain(domain sesbooks.Domain) Sesbooks {
	return Sesbooks{
		ID:                domain.ID,
		PatientsID:        domain.IDPatient,
		SessionScheduleID: domain.IDSessionSchedule,
		PatientQueue:      domain.PatientQueue,
		Date:              domain.Date,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Sesbooks) sesbooks.Domain {
	return sesbooks.Domain{
		ID:                ev.ID,
		IDPatient:         ev.PatientsID,
		IDSessionSchedule: ev.SessionScheduleID,
		PatientQueue:      ev.PatientQueue,
		Date:              ev.Date,
		CreatedAt:         ev.CreatedAt,
		UpdatedAt:         ev.UpdatedAt,
	}
}

func toDomainList(data []Sesbooks) []sesbooks.Domain {
	result := []sesbooks.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
