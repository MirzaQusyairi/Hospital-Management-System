package schedules

import "time"

type Domain struct {
	ID        int
	Day       string
	Start     string
	End       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AllSchedule() ([]Domain, error)
	AddSchedule(domain *Domain) (Domain, error)
	Update(schID int, domain *Domain) (Domain, error)
	ScheduleByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllSchedule() ([]Domain, error)
	AddSchedule(domain *Domain) (Domain, error)
	Update(schID int, domain *Domain) (Domain, error)
	ScheduleByID(id int) (Domain, error)
	Delete(id int) (string, error)
}
