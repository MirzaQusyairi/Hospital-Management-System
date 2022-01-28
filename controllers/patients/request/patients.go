package request

import (
	"Hospital-Management-System/business/patients"
	"time"
)

type Patients struct {
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	NIK      int    `json:"nik"`
	Address  string `json:"address"`
	DOB      string `json:"dob"`
	No_Rm    string `json:"no_rm"`
}
type PatientResponse struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	NIK       int       `json:"nik"`
	Address   string    `json:"address"`
	DOB       string    `json:"dob"`
	No_Rm     string    `json:"no_rm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (req *Patients) ToDomain() *patients.Domain {
	return &patients.Domain{

		Fullname: req.Fullname,
		Gender:   req.Gender,
		Age:      req.Age,
		NIK:      req.NIK,
		Address:  req.Address,
		DOB:      req.DOB,
		No_Rm:    req.No_Rm,
	}
}
func FromDomainAllPatient(domain patients.Domain) PatientResponse {
	return PatientResponse{
		ID:        domain.ID,
		Fullname:  domain.Fullname,
		Gender:    domain.Gender,
		Age:       domain.Age,
		NIK:       domain.NIK,
		Address:   domain.Address,
		DOB:       domain.DOB,
		No_Rm:     domain.No_Rm,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
