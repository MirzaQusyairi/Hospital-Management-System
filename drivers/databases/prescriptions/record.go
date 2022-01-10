package prescriptions

import (
	"Hospital-Management-System/business/prescriptions"
	"time"

	"gorm.io/gorm"
)

type Prescriptions struct {
	gorm.Model
	ID               int `gorm:"primary_key"`
	MedicineName     string
	MedicationRules  string
	IDPatient        string
	IDDoctor         string
	IDSessionBooking string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func toDomain(ev Prescriptions) prescriptions.Domain {
	return prescriptions.Domain{
		ID:               ev.ID,
		MedicineName:     ev.MedicineName,
		MedicationRules:  ev.MedicationRules,
		IDPatient:        ev.IDPatient,
		IDDoctor:         ev.IDDoctor,
		IDSessionBooking: ev.IDSessionBooking,
		CreatedAt:        ev.CreatedAt,
		UpdatedAt:        ev.UpdatedAt,
	}
}

func fromDomain(domain prescriptions.Domain) Prescriptions {
	return Prescriptions{
		ID:               domain.ID,
		MedicineName:     domain.MedicineName,
		MedicationRules:  domain.MedicationRules,
		IDPatient:        domain.IDPatient,
		IDDoctor:         domain.IDDoctor,
		IDSessionBooking: domain.IDSessionBooking,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Prescriptions) prescriptions.Domain {
	return prescriptions.Domain{
		ID:               ev.ID,
		MedicineName:     ev.MedicineName,
		MedicationRules:  ev.MedicationRules,
		IDPatient:        ev.IDPatient,
		IDDoctor:         ev.IDDoctor,
		IDSessionBooking: ev.IDSessionBooking,
		CreatedAt:        ev.CreatedAt,
		UpdatedAt:        ev.UpdatedAt,
	}
}

func toDomainList(data []Prescriptions) []prescriptions.Domain {
	result := []prescriptions.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
