package seschedules

import "time"

type Domain struct {
	ID         int
	IDSchedule string
	IDDoctor   string
	IDFacilty  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	AllSessionSchedule() ([]Domain, error)
	AddSessionSchedule(domain *Domain) (Domain, error)
	Update(schID int, domain *Domain) (Domain, error)
	SessionScheduleByID(id int) (Domain, error)
	SessionScheduleByIDSchedule(id int) ([]Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllSessionSchedule() ([]Domain, error)
	AddSessionSchedule(domain *Domain) (Domain, error)
	Update(schID int, domain *Domain) (Domain, error)
	SessionScheduleByID(id int) (Domain, error)
	SessionScheduleByIDSchedule(id int) ([]Domain, error)
	Delete(id int) (string, error)
}
