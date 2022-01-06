package response

import (
	"Hospital-Management-System/business/schedules"
	"time"
)

type ScheduleRegisterResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id"`
	Day       string    `json:"day"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ScheduleResponse struct {
	ID        int       `json:"id:"`
	Day       string    `json:"day"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain schedules.Domain) ScheduleRegisterResponse {
	return ScheduleRegisterResponse{
		Message:   "Add Schedule Success",
		ID:        domain.ID,
		Day:       domain.Day,
		Start:     domain.Start,
		End:       domain.End,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainAllSchedule(domain schedules.Domain) ScheduleResponse {
	return ScheduleResponse{
		ID:        domain.ID,
		Day:       domain.Day,
		Start:     domain.Start,
		End:       domain.End,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdateSchedule(domain schedules.Domain) ScheduleRegisterResponse {
	return ScheduleRegisterResponse{
		Message:   "Update Schedule Data Success",
		ID:        domain.ID,
		Day:       domain.Day,
		Start:     domain.Start,
		End:       domain.End,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromScheduleListDomain(domain []schedules.Domain) []ScheduleResponse {
	var response []ScheduleResponse
	for _, value := range domain {
		response = append(response, FromDomainAllSchedule(value))
	}
	return response
}
