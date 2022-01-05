package response

import (
	"Hospital-Management-System/business/facilities"
	"time"
)

type AddResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	Queue     int       `json:"queue"`
	Location  string    `json:"location"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	Queue     int       `json:"queue"`
	Location  string    `json:"location"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type FaciltyResponse struct {
	ID        int       `json:"id:"`
	Name      string    `json:"name"`
	Queue     int       `json:"queue"`
	Location  string    `json:"location"`
	Capacity  int       `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain facilities.Domain) AddResponse {
	return AddResponse{
		Message:   "Add Facilty Success",
		ID:        domain.ID,
		Name:      domain.Name,
		Queue:     domain.Queue,
		Location:  domain.Location,
		Capacity:  domain.Capacity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromDomainAllFacilty(domain facilities.Domain) FaciltyResponse {
	return FaciltyResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Queue:     domain.Queue,
		Location:  domain.Location,
		Capacity:  domain.Capacity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdateDoctor(domain facilities.Domain) UpdateResponse {
	return UpdateResponse{
		Message:   "Update Facilty Success",
		ID:        domain.ID,
		Name:      domain.Name,
		Queue:     domain.Queue,
		Location:  domain.Location,
		Capacity:  domain.Capacity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func FromFaciltyListDomain(domain []facilities.Domain) []FaciltyResponse {
	var response []FaciltyResponse
	for _, value := range domain {
		response = append(response, FromDomainAllFacilty(value))
	}
	return response
}
