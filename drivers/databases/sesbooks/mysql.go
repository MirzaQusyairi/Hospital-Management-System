package sesbooks

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/sesbooks"
	"log"

	"gorm.io/gorm"
)

type MysqlSesbooksRepository struct {
	Conn *gorm.DB
}

func NewMysqlSesbooksRepository(conn *gorm.DB) sesbooks.Repository {
	return &MysqlSesbooksRepository{
		Conn: conn,
	}
}

func (rep *MysqlSesbooksRepository) UpdateStatus(sesID int, status string) (string, error) {
	log.Println(sesID, status, "sesID, status")
	result := rep.Conn.Where("id = ?", sesID).Model(Sesbooks{}).Update("status", "Checked")
	if result.Error != nil {
		return "failed", result.Error
	}
	log.Println(result)
	return "success", nil
}

func (rep *MysqlSesbooksRepository) AddSessionBook(domain *sesbooks.Domain) (sesbooks.Domain, error) {

	sesbookking := fromDomain(*domain)

	result := rep.Conn.Create(&sesbookking)
	if result.Error != nil {
		return sesbooks.Domain{}, result.Error
	}

	return toDomain(sesbookking), nil
}

func (rep *MysqlSesbooksRepository) Update(sesID int, domain *sesbooks.Domain) (sesbooks.Domain, error) {

	sesbooknUpdate := fromDomain(*domain)

	sesbooknUpdate.ID = sesID

	result := rep.Conn.Where("id = ?", sesID).Updates(&sesbooknUpdate)

	if result.Error != nil {
		return sesbooks.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(sesbooknUpdate), nil
}

func (rep *MysqlSesbooksRepository) AllSessionBook() ([]sesbooks.Domain, error) {

	var sesbooking []Sesbooks

	result := rep.Conn.Find(&sesbooking)

	if result.Error != nil {
		return []sesbooks.Domain{}, result.Error
	}

	return toDomainList(sesbooking), nil

}

func (rep *MysqlSesbooksRepository) SessionBookByID(id int) (sesbooks.Domain, error) {

	var sesbooking Sesbooks

	result := rep.Conn.Where("id = ?", id).First(&sesbooking)

	if result.Error != nil {
		return sesbooks.Domain{}, result.Error
	}

	return toDomain(sesbooking), nil
}

func (rep *MysqlSesbooksRepository) Delete(id int) (string, error) {
	rec := Sesbooks{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Session Booking has been delete", nil
}
