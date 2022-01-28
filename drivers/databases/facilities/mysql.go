package facilities

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/facilities"

	"gorm.io/gorm"
)

type MysqlFaciltyRepository struct {
	Conn *gorm.DB
}

func NewMysqlFaciltyRepository(conn *gorm.DB) facilities.Repository {
	return &MysqlFaciltyRepository{
		Conn: conn,
	}
}

func (rep *MysqlFaciltyRepository) UpdateQueue(facID int, queue int) (int, error) {
	err := rep.Conn.Model(Facilities{}).Where("id = ?", facID).Update("queue", queue+1).Error
	if err != nil {
		return 0, err
	}
	return queue + 1, nil
}
func (rep *MysqlFaciltyRepository) AllFacilty() ([]facilities.Domain, error) {

	var facilty []Facilities

	result := rep.Conn.Find(&facilty)

	if result.Error != nil {
		return []facilities.Domain{}, result.Error
	}

	return toDomainList(facilty), nil

}
func (rep *MysqlFaciltyRepository) AddFacilty(domain *facilities.Domain) (facilities.Domain, error) {

	org := fromDomain(*domain)

	result := rep.Conn.Create(&org)

	if result.Error != nil {
		return facilities.Domain{}, result.Error
	}

	return toDomain(org), nil
}

func (rep *MysqlFaciltyRepository) Update(facID int, domain *facilities.Domain) (facilities.Domain, error) {

	facUpdate := fromDomain(*domain)

	facUpdate.ID = facID

	result := rep.Conn.Where("id = ?", facID).Updates(&facUpdate)

	if result.Error != nil {
		return facilities.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(facUpdate), nil
}
func (rep *MysqlFaciltyRepository) FacByID(id int) (facilities.Domain, error) {

	var facilty Facilities

	result := rep.Conn.Where("id = ?", id).First(&facilty)

	if result.Error != nil {
		return facilities.Domain{}, result.Error
	}

	return toDomain(facilty), nil
}
func (rep *MysqlFaciltyRepository) Delete(id int) (string, error) {
	rec := Facilities{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Facilty has been delete", nil

}
