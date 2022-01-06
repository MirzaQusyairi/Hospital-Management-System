package schedules

import (
	"Hospital-Management-System/business"
)

type serviceSchedule struct {
	scheduleRepository Repository
}

func NewServiceSchedule(repoSchedule Repository) Service {
	return &serviceSchedule{
		scheduleRepository: repoSchedule,
	}
}

func (serv *serviceSchedule) AllSchedule() ([]Domain, error) {

	result, err := serv.scheduleRepository.AllSchedule()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *serviceSchedule) AddSchedule(domain *Domain) (Domain, error) {

	result, err := serv.scheduleRepository.AddSchedule(domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSchedule) Update(schID int, domain *Domain) (Domain, error) {

	result, err := serv.scheduleRepository.Update(schID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSchedule) ScheduleByID(id int) (Domain, error) {

	result, err := serv.scheduleRepository.ScheduleByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *serviceSchedule) Delete(id int) (string, error) {

	result, err := serv.scheduleRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
