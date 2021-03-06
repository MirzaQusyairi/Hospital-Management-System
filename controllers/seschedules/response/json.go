package response

import (
	"Hospital-Management-System/business/seschedules"
	"time"
)

type SessionScheduleRegisterResponse struct {
	Message    string    `json:"message"`
	ID         int       `json:"id"`
	IDSchedule int       `json:"id_schedule"`
	IDDoctor   int       `json:"id_doctor"`
	IDFacilty  int       `json:"id_facilty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type SessionScheduleResponse struct {
	ID         int       `json:"id"`
	IDSchedule int       `json:"id_schedule"`
	IDDoctor   int       `json:"id_doctor"`
	IDFacilty  int       `json:"id_facilty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomainRegister(domain seschedules.Domain) SessionScheduleRegisterResponse {
	return SessionScheduleRegisterResponse{
		Message:    "Add Session Schedule Success",
		ID:         domain.ID,
		IDSchedule: domain.IDSchedule,
		IDDoctor:   domain.IDDoctor,
		IDFacilty:  domain.IDFacilty,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainAllSessionSchedule(domain seschedules.Domain) SessionScheduleResponse {
	return SessionScheduleResponse{
		ID:         domain.ID,
		IDSchedule: domain.IDSchedule,
		IDDoctor:   domain.IDDoctor,
		IDFacilty:  domain.IDFacilty,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainUpdateSessionSchedule(domain seschedules.Domain) SessionScheduleRegisterResponse {
	return SessionScheduleRegisterResponse{
		Message:    "Update Sesscion Schedule Data Success",
		ID:         domain.ID,
		IDSchedule: domain.IDSchedule,
		IDDoctor:   domain.IDDoctor,
		IDFacilty:  domain.IDFacilty,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromSessionScheduleListDomain(domain []seschedules.Domain) []SessionScheduleResponse {
	var response []SessionScheduleResponse
	for _, value := range domain {
		response = append(response, FromDomainAllSessionSchedule(value))
	}
	return response
}
