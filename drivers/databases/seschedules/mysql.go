package seschedules

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/seschedules"

	"gorm.io/gorm"
)

type MysqlSscheduleRepository struct {
	Conn *gorm.DB
}

func NewMysqlSscheduleRepository(conn *gorm.DB) seschedules.Repository {
	return &MysqlSscheduleRepository{
		Conn: conn,
	}
}

func (rep *MysqlSscheduleRepository) AddSessionSchedule(domain *seschedules.Domain) (seschedules.Domain, error) {

	pres := fromDomain(*domain)

	result := rep.Conn.Create(&pres)
	if result.Error != nil {
		return seschedules.Domain{}, result.Error
	}

	return toDomain(pres), nil
}

func (rep *MysqlSscheduleRepository) Update(presID int, domain *seschedules.Domain) (seschedules.Domain, error) {

	sescheduleUpdate := fromDomain(*domain)

	sescheduleUpdate.ID = presID

	result := rep.Conn.Where("id = ?", presID).Updates(&sescheduleUpdate)

	if result.Error != nil {
		return seschedules.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(sescheduleUpdate), nil
}

func (rep *MysqlSscheduleRepository) AllSessionSchedule() ([]seschedules.Domain, error) {

	var sschedule []Sschedules

	result := rep.Conn.Find(&sschedule)

	if result.Error != nil {
		return []seschedules.Domain{}, result.Error
	}

	return toDomainList(sschedule), nil

}

func (rep *MysqlSscheduleRepository) SessionScheduleByID(id int) (seschedules.Domain, error) {

	var sschedule Sschedules

	result := rep.Conn.Where("id = ?", id).First(&sschedule)

	if result.Error != nil {
		return seschedules.Domain{}, result.Error
	}

	return toDomain(sschedule), nil
}

func (rep *MysqlSscheduleRepository) SessionScheduleByIDSchedule(id int) ([]seschedules.Domain, error) {

	var sschedule []Sschedules

	result := rep.Conn.Where("id_schedule = ?", id).Find(&sschedule)

	if result.Error != nil {
		return []seschedules.Domain{}, result.Error
	}

	return toDomainList(sschedule), nil
}

func (rep *MysqlSscheduleRepository) Delete(id int) (string, error) {
	rec := Sschedules{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Session Schedule has been delete", nil
}
