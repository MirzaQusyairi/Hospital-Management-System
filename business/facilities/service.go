package facilities

import (
	"Hospital-Management-System/business"
)

type serviceFacilty struct {
	faciltyRepository Repository
}

func NewServiceFacilty(repoFacilty Repository) Service {
	return &serviceFacilty{
		faciltyRepository: repoFacilty,
	}
}
func (serv *serviceFacilty) AllFacilty() ([]Domain, error) {

	result, err := serv.faciltyRepository.AllFacilty()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
func (serv *serviceFacilty) AddFacilty(domain *Domain) (Domain, error) {

	result, err := serv.faciltyRepository.AddFacilty(domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceFacilty) Update(facID int, domain *Domain) (Domain, error) {

	result, err := serv.faciltyRepository.Update(facID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceFacilty) FacByID(id int) (Domain, error) {

	result, err := serv.faciltyRepository.FacByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
func (serv *serviceFacilty) Delete(id int) (string, error) {

	result, err := serv.faciltyRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
