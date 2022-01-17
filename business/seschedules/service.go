package seschedules

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/doctors"
	"Hospital-Management-System/business/facilities"
	"Hospital-Management-System/business/schedules"
)

type serviceSessionSchedule struct {
	SessionScheduleRepository Repository
	FacilitiesRepository      facilities.Repository
	DoctorsRepository         doctors.Repository
	SchedulesRepository       schedules.Repository
}

func NewServiceSessionSchedule(repoSessionSchedule Repository, repoFacilities facilities.Repository, repoDoctors doctors.Repository, repoSchedules schedules.Repository) Service {
	return &serviceSessionSchedule{
		SessionScheduleRepository: repoSessionSchedule,
		FacilitiesRepository:      repoFacilities,
		DoctorsRepository:         repoDoctors,
		SchedulesRepository:       repoSchedules,
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
	resultFacilities, err := serv.FacilitiesRepository.FacByID(domain.IDFacilty)
	if err != nil {
		return Domain{}, err
	}
	domain.IDFacilty = resultFacilities.ID
	resultDoctors, err := serv.DoctorsRepository.DoctorByID(domain.IDDoctor)
	if err != nil {
		return Domain{}, err
	}
	domain.IDDoctor = resultDoctors.ID
	resultSchedules, err := serv.SchedulesRepository.ScheduleByID(domain.IDSchedule)
	if err != nil {
		return Domain{}, err
	}
	domain.IDSchedule = resultSchedules.ID
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
