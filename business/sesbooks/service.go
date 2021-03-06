package sesbooks

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/facilities"
	"Hospital-Management-System/business/patients"
	"Hospital-Management-System/business/seschedules"
	"time"
)

type serviceSesbook struct {
	sesbookRepository        Repository
	PatientsRepository       patients.Repository
	SessionScheduleRepositry seschedules.Repository
	FacilitiesRepository     facilities.Repository
}

func NewServiceSesbook(repoSesbook Repository, repoPatients patients.Repository, repoSessionSchedule seschedules.Repository, repoFacility facilities.Repository) Service {
	return &serviceSesbook{
		sesbookRepository:        repoSesbook,
		PatientsRepository:       repoPatients,
		SessionScheduleRepositry: repoSessionSchedule,
		FacilitiesRepository:     repoFacility,
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
	resultPatients, err := serv.PatientsRepository.PatientByID(domain.IDPatient)
	if err != nil {
		return Domain{}, err
	}
	domain.IDPatient = resultPatients.ID
	resultSessionSchedule, err := serv.SessionScheduleRepositry.SessionScheduleByID(domain.IDSessionSchedule)
	if err != nil {
		return Domain{}, err
	}
	findLatestQueue, err := serv.FacilitiesRepository.FacByID(resultSessionSchedule.IDFacilty)
	if err != nil {
		return Domain{}, err
	}
	if findLatestQueue.Queue+1 > findLatestQueue.Capacity {
		return Domain{}, err
	}
	lastQueue := findLatestQueue.Queue
	resultUpdateQueue, err := serv.FacilitiesRepository.UpdateQueue(resultSessionSchedule.IDFacilty, lastQueue)
	if err != nil {
		return Domain{}, err
	}
	domain.PatientQueue = resultUpdateQueue
	domain.IDSessionSchedule = resultSessionSchedule.ID
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
