package prescriptions

import (
	"Hospital-Management-System/business"
)

type servicePrescription struct {
	prescriptionRepository Repository
}

func NewServicePrescription(repoPrescription Repository) Service {
	return &servicePrescription{
		prescriptionRepository: repoPrescription,
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
