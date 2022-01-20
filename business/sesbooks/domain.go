package sesbooks

import "time"

type Domain struct {
	ID                int
	IDPatient         int
	IDSessionSchedule int
	PatientQueue      int
	Date              time.Time
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Service interface {
	AllSessionBook() ([]Domain, error)
	AddSessionBook(domain *Domain) (Domain, error)
	Update(sesID int, domain *Domain) (Domain, error)
	SessionBookByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllSessionBook() ([]Domain, error)
	AddSessionBook(domain *Domain) (Domain, error)
	Update(sesID int, domain *Domain) (Domain, error)
	SessionBookByID(id int) (Domain, error)
	Delete(id int) (string, error)
}
