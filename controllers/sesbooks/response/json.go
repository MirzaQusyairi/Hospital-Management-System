package response

import (
	"Hospital-Management-System/business/sesbooks"
	"time"
)

type SesbookRegisterResponse struct {
	Message           string    `json:"message"`
	ID                int       `json:"id"`
	IDPatient         int       `json:"id_patient"`
	IDSessionSchedule int       `json:"id_session_schedule"`
	PatientQueue      int       `json:"patient_queue"`
	Date              time.Time `json:"date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
type SesbookResponse struct {
	ID                int       `json:"id"`
	IDPatient         int       `json:"id_patient"`
	IDSessionSchedule int       `json:"id_session_schedule"`
	PatientQueue      int       `json:"patient_queue"`
	Date              time.Time `json:"date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FromDomainRegister(domain sesbooks.Domain) SesbookRegisterResponse {
	return SesbookRegisterResponse{
		Message:           "Add Session Booking Success",
		ID:                domain.ID,
		IDPatient:         domain.IDPatient,
		IDSessionSchedule: domain.IDSessionSchedule,
		PatientQueue:      domain.PatientQueue,
		Date:              domain.Date,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func FromDomainAllSesbook(domain sesbooks.Domain) SesbookResponse {
	return SesbookResponse{
		ID:                domain.ID,
		IDPatient:         domain.IDPatient,
		IDSessionSchedule: domain.IDSessionSchedule,
		PatientQueue:      domain.PatientQueue,
		Date:              domain.Date,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func FromDomainUpdateSesbook(domain sesbooks.Domain) SesbookRegisterResponse {
	return SesbookRegisterResponse{
		Message:           "Update Session Booking Data Success",
		ID:                domain.ID,
		IDPatient:         domain.IDPatient,
		IDSessionSchedule: domain.IDSessionSchedule,
		PatientQueue:      domain.PatientQueue,
		Date:              domain.Date,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}

func FromSesbookListDomain(domain []sesbooks.Domain) []SesbookResponse {
	var response []SesbookResponse
	for _, value := range domain {
		response = append(response, FromDomainAllSesbook(value))
	}
	return response
}
