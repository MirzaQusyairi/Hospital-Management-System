package patients

import (
	"Hospital-Management-System/business/patients"
	"time"

	"gorm.io/gorm"
)

type Patients struct {
	gorm.Model
	ID                    int `gorm:"primary_key"`
	MedicalPrescriptionID int
	MedicalRecordID       int
	Fullname              string
	Address               string
	Gender                string
	Age                   int
	NIK                   int
	No_Rm                 string
	DOB                   string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func toDomain(ex Patients) patients.Domain {
	return patients.Domain{
		ID:                    ex.ID,
		MedicalPrescriptionID: ex.MedicalPrescriptionID,
		MedicalRecordID:       ex.MedicalRecordID,
		Fullname:              ex.Fullname,
		Address:               ex.Address,
		Gender:                ex.Gender,
		Age:                   ex.Age,
		NIK:                   ex.NIK,
		No_Rm:                 ex.No_Rm,
		DOB:                   ex.DOB,
		CreatedAt:             ex.CreatedAt,
		UpdatedAt:             ex.UpdatedAt,
	}
}

func fromDomain(domain patients.Domain) Patients {
	return Patients{
		ID:                    domain.ID,
		MedicalPrescriptionID: domain.MedicalPrescriptionID,
		MedicalRecordID:       domain.MedicalRecordID,
		Fullname:              domain.Fullname,
		Address:               domain.Address,
		Gender:                domain.Gender,
		Age:                   domain.Age,
		NIK:                   domain.NIK,
		No_Rm:                 domain.No_Rm,
		DOB:                   domain.DOB,
		CreatedAt:             domain.CreatedAt,
		UpdatedAt:             domain.UpdatedAt,
	}
}

func toDomainUpdate(ex Patients) patients.Domain {
	return patients.Domain{
		ID:                    ex.ID,
		MedicalPrescriptionID: ex.MedicalPrescriptionID,
		MedicalRecordID:       ex.MedicalRecordID,
		Fullname:              ex.Fullname,
		Address:               ex.Address,
		Gender:                ex.Gender,
		Age:                   ex.Age,
		NIK:                   ex.NIK,
		No_Rm:                 ex.No_Rm,
		DOB:                   ex.DOB,
		CreatedAt:             ex.CreatedAt,
		UpdatedAt:             ex.UpdatedAt,
	}
}
func toDomainList(data []Patients) []patients.Domain {
	result := []patients.Domain{}

	for _, ev := range data {
		result = append(result, toDomain(ev))
	}
	return result
}
