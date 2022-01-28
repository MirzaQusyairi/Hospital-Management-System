package facilities

import "time"

type Domain struct {
	ID        int
	Name      string
	Queue     int
	Location  string
	Capacity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AllFacilty() ([]Domain, error)
	AddFacilty(domain *Domain) (Domain, error)
	Update(facID int, domain *Domain) (Domain, error)
	FacByID(id int) (Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	AllFacilty() ([]Domain, error)
	AddFacilty(domain *Domain) (Domain, error)
	Update(facID int, domain *Domain) (Domain, error)
	UpdateQueue(facID int, queue int) (int, error)
	FacByID(id int) (Domain, error)
	Delete(id int) (string, error)
}
