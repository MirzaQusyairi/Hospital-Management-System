package prescriptions

import (
	"Hospital-Management-System/business/prescriptions"
	"Hospital-Management-System/drivers/databases/doctors"
	"Hospital-Management-System/drivers/databases/patients"
	"Hospital-Management-System/drivers/databases/sesbooks"
	"time"

	"gorm.io/gorm"
)

type Prescriptions struct {
	gorm.Model
	ID               int `gorm:"primary_key"`
	MedicineName     string
	MedicationRules  string
	PatientsID       int
	Patients         patients.Patients `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	DoctorsID        int
	Doctors          doctors.Doctors `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	SessionBookingID int
	SessionBooking   sesbooks.Sesbooks `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func toDomain(ev Prescriptions) prescriptions.Domain {
	return prescriptions.Domain{
		ID:               ev.ID,
		MedicineName:     ev.MedicineName,
		MedicationRules:  ev.MedicationRules,
		PatientsID:       ev.PatientsID,
		DoctorID:         ev.DoctorsID,
		SessionBookingID: ev.SessionBookingID,
		CreatedAt:        ev.CreatedAt,
		UpdatedAt:        ev.UpdatedAt,
	}
}

func fromDomain(domain prescriptions.Domain) Prescriptions {
	return Prescriptions{
		ID:               domain.ID,
		MedicineName:     domain.MedicineName,
		MedicationRules:  domain.MedicationRules,
		PatientsID:       domain.PatientsID,
		DoctorsID:        domain.DoctorID,
		SessionBookingID: domain.SessionBookingID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func toDomainUpdate(ev Prescriptions) prescriptions.Domain {
	return prescriptions.Domain{
		ID:               ev.ID,
		MedicineName:     ev.MedicineName,
		MedicationRules:  ev.MedicationRules,
		PatientsID:       ev.PatientsID,
		DoctorID:         ev.DoctorsID,
		SessionBookingID: ev.SessionBookingID,
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
