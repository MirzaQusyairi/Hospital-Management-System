package seschedules

import (
	"Hospital-Management-System/business"
)

type serviceSessionSchedule struct {
	SessionScheduleRepository Repository
}

func NewServiceSessionSchedule(repoSessionSchedule Repository) Service {
	return &serviceSessionSchedule{
		SessionScheduleRepository: repoSessionSchedule,
	}
}

func (serv *serviceSessionSchedule) AllSessionSchedule() ([]Domain, error) {

	result, err := serv.SessionScheduleRepository.AllSessionSchedule()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceSessionSchedule) AddSessionSchedule(domain *Domain) (Domain, error) {

	result, err := serv.SessionScheduleRepository.AddSessionSchedule(domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSessionSchedule) Update(schID int, domain *Domain) (Domain, error) {

	result, err := serv.SessionScheduleRepository.Update(schID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSessionSchedule) SessionScheduleByID(id int) (Domain, error) {

	result, err := serv.SessionScheduleRepository.SessionScheduleByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSessionSchedule) SessionScheduleByIDSchedule(id int) ([]Domain, error) {

	result, err := serv.SessionScheduleRepository.SessionScheduleByIDSchedule(id)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceSessionSchedule) Delete(id int) (string, error) {

	result, err := serv.SessionScheduleRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
