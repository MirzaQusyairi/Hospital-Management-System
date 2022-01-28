package prescriptions

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/doctors"
	"Hospital-Management-System/business/patients"
	"Hospital-Management-System/business/sesbooks"
	"log"
)

type servicePrescription struct {
	prescriptionRepository Repository
	patientRepository      patients.Repository
	doctorRepository       doctors.Repository
	sesbookRepository      sesbooks.Repository
}

func NewServicePrescription(repoPrescription Repository, repoPatients patients.Repository, repoDoctors doctors.Repository, repoSesbook sesbooks.Repository) Service {
	return &servicePrescription{
		prescriptionRepository: repoPrescription,
		patientRepository:      repoPatients,
		doctorRepository:       repoDoctors,
		sesbookRepository:      repoSesbook,
	}
}

func (serv *servicePrescription) AllPrescription() ([]Domain, error) {

	result, err := serv.prescriptionRepository.AllPrescription()

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *servicePrescription) AddPrescription(domain *Domain) (Domain, error) {
	resultPatients, err := serv.patientRepository.PatientByID(domain.PatientsID)
	if err != nil {
		return Domain{}, err
	}
	domain.PatientsID = resultPatients.ID
	resultDoctors, err := serv.doctorRepository.DoctorByID(domain.DoctorID)
	if err != nil {
		return Domain{}, err
	}
	domain.DoctorID = resultDoctors.ID
	resultSesbook, err := serv.sesbookRepository.SessionBookByID(domain.SessionBookingID)
	if err != nil {
		return Domain{}, err
	}
	domain.SessionBookingID = resultSesbook.ID
	updateStatus, err := serv.sesbookRepository.UpdateStatus(domain.SessionBookingID, "Checked")
	if err != nil {
		return Domain{}, err
	}
	log.Println(updateStatus)
	result, err := serv.prescriptionRepository.AddPrescription(domain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (serv *servicePrescription) Update(schID int, domain *Domain) (Domain, error) {

	result, err := serv.prescriptionRepository.Update(schID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *servicePrescription) PrescriptionByID(id int) (Domain, error) {

	result, err := serv.prescriptionRepository.PrescriptionByID(id)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *servicePrescription) PrescriptionByIDSessionBooking(id int) ([]Domain, error) {

	result, err := serv.prescriptionRepository.PrescriptionByIDSessionBooking(id)

	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (serv *servicePrescription) Delete(id int) (string, error) {

	result, err := serv.prescriptionRepository.Delete(id)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}
