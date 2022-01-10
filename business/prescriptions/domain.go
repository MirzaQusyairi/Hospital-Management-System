package prescriptions

import "time"

type Domain struct {
	ID               int
	MedicineName     string
	MedicationRules  string
	IDPatient        string
	IDDoctor         string
	IDSessionBooking string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	AllPrescription() ([]Domain, error)
	AddPrescription(domain *Domain) (Domain, error)
	Update(schID int, domain *Domain) (Domain, error)
	PrescriptionByID(id int) (Domain, error)
	PrescriptionByIDSessionBooking(id int) ([]Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllPrescription() ([]Domain, error)
	AddPrescription(domain *Domain) (Domain, error)
	Update(schID int, domain *Domain) (Domain, error)
	PrescriptionByID(id int) (Domain, error)
	PrescriptionByIDSessionBooking(id int) ([]Domain, error)
	Delete(id int) (string, error)
}
