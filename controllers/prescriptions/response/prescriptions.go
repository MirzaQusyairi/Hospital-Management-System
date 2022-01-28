package response

import (
	"Hospital-Management-System/business/prescriptions"
	"time"
)

type PrescriptionRegisterResponse struct {
	Message          string    `json:"message"`
	ID               int       `json:"id"`
	MedicineName     string    `json:"medicine_name"`
	MedicationRules  string    `json:"medication_rules"`
	PatientsID       int       `json:"id_patient"`
	DoctorID         int       `json:"id_doctor"`
	SessionBookingID int       `json:"id_sessionbooking"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
type PrescriptionResponse struct {
	ID               int       `json:"id"`
	MedicineName     string    `json:"medicine_name"`
	MedicationRules  string    `json:"medication_rules"`
	PatientsID       int       `json:"id_patient"`
	DoctorID         int       `json:"id_doctor"`
	SessionBookingID int       `json:"id_sessionbooking"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FromDomainRegister(domain prescriptions.Domain) PrescriptionRegisterResponse {
	return PrescriptionRegisterResponse{
		Message:          "Add Prescription Success",
		ID:               domain.ID,
		MedicineName:     domain.MedicineName,
		MedicationRules:  domain.MedicationRules,
		PatientsID:       domain.PatientsID,
		DoctorID:         domain.DoctorID,
		SessionBookingID: domain.SessionBookingID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func FromDomainAllPrescription(domain prescriptions.Domain) PrescriptionResponse {
	return PrescriptionResponse{
		ID:               domain.ID,
		MedicineName:     domain.MedicineName,
		MedicationRules:  domain.MedicationRules,
		PatientsID:       domain.PatientsID,
		DoctorID:         domain.DoctorID,
		SessionBookingID: domain.SessionBookingID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func FromDomainUpdatePrescription(domain prescriptions.Domain) PrescriptionRegisterResponse {
	return PrescriptionRegisterResponse{
		Message:          "Update Prescription Data Success",
		ID:               domain.ID,
		MedicineName:     domain.MedicineName,
		MedicationRules:  domain.MedicationRules,
		PatientsID:       domain.PatientsID,
		DoctorID:         domain.DoctorID,
		SessionBookingID: domain.SessionBookingID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func FromPrescriptionListDomain(domain []prescriptions.Domain) []PrescriptionResponse {
	var response []PrescriptionResponse
	for _, value := range domain {
		response = append(response, FromDomainAllPrescription(value))
	}
	return response
}
