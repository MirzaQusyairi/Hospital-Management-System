package sesbooks

import (
	"Hospital-Management-System/business"
	"time"
)

type serviceSesbook struct {
	sesbookRepository Repository
}

func NewServiceSesbook(repoSesbook Repository) Service {
	return &serviceSesbook{
		sesbookRepository: repoSesbook,
	}
}

func (serv *serviceSesbook) AllSessionBook() ([]Domain, error) {

	result, err := serv.sesbookRepository.AllSessionBook()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceSesbook) AddSessionBook(domain *Domain) (Domain, error) {
	dateNow := time.Now()
	domain.Date = dateNow
	result, err := serv.sesbookRepository.AddSessionBook(domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSesbook) Update(sesID int, domain *Domain) (Domain, error) {

	result, err := serv.sesbookRepository.Update(sesID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSesbook) SessionBookByID(id int) (Domain, error) {

	result, err := serv.sesbookRepository.SessionBookByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSesbook) Delete(id int) (string, error) {

	result, err := serv.sesbookRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
