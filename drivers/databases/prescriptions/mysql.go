package prescriptions

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/prescriptions"

	"gorm.io/gorm"
)

type MysqlPrescriptionRepository struct {
	Conn *gorm.DB
}

func NewMysqlPrescriptionRepository(conn *gorm.DB) prescriptions.Repository {
	return &MysqlPrescriptionRepository{
		Conn: conn,
	}
}

func (rep *MysqlPrescriptionRepository) AddPrescription(domain *prescriptions.Domain) (prescriptions.Domain, error) {

	pres := fromDomain(*domain)

	result := rep.Conn.Create(&pres)
	if result.Error != nil {
		return prescriptions.Domain{}, result.Error
	}

	return toDomain(pres), nil
}

func (rep *MysqlPrescriptionRepository) Update(presID int, domain *prescriptions.Domain) (prescriptions.Domain, error) {

	prescriptionUpdate := fromDomain(*domain)

	prescriptionUpdate.ID = presID

	result := rep.Conn.Where("id = ?", presID).Updates(&prescriptionUpdate)

	if result.Error != nil {
		return prescriptions.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(prescriptionUpdate), nil
}

func (rep *MysqlPrescriptionRepository) AllPrescription() ([]prescriptions.Domain, error) {

	var prescription []Prescriptions

	result := rep.Conn.Find(&prescription)

	if result.Error != nil {
		return []prescriptions.Domain{}, result.Error
	}

	return toDomainList(prescription), nil

}

func (rep *MysqlPrescriptionRepository) PrescriptionByID(id int) (prescriptions.Domain, error) {

	var prescription Prescriptions

	result := rep.Conn.Where("id = ?", id).First(&prescription)

	if result.Error != nil {
		return prescriptions.Domain{}, result.Error
	}

	return toDomain(prescription), nil
}

func (rep *MysqlPrescriptionRepository) PrescriptionByIDSessionBooking(id int) ([]prescriptions.Domain, error) {

	var prescription []Prescriptions

	result := rep.Conn.Where("id_session_booking = ?", id).Find(&prescription)

	if result.Error != nil {
		return []prescriptions.Domain{}, result.Error
	}

	return toDomainList(prescription), nil
}

func (rep *MysqlPrescriptionRepository) Delete(id int) (string, error) {
	rec := Prescriptions{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Prescription has been delete", nil
}
