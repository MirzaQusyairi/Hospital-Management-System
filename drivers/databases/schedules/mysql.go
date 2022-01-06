package schedules

import (
	"Hospital-Management-System/business"
	"Hospital-Management-System/business/schedules"

	"gorm.io/gorm"
)

type MysqlScheduleRepository struct {
	Conn *gorm.DB
}

func NewMysqlScheduleRepository(conn *gorm.DB) schedules.Repository {
	return &MysqlScheduleRepository{
		Conn: conn,
	}
}

func (rep *MysqlScheduleRepository) AllSchedule() ([]schedules.Domain, error) {

	var schedule []Schedules

	result := rep.Conn.Find(&schedule)

	if result.Error != nil {
		return []schedules.Domain{}, result.Error
	}

	return toDomainList(schedule), nil

}

func (rep *MysqlScheduleRepository) AddSchedule(domain *schedules.Domain) (schedules.Domain, error) {

	sch := fromDomain(*domain)

	result := rep.Conn.Create(&sch)
	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}

	return toDomain(sch), nil
}

func (rep *MysqlScheduleRepository) Update(schID int, domain *schedules.Domain) (schedules.Domain, error) {

	scheduleUpdate := fromDomain(*domain)

	scheduleUpdate.ID = schID

	result := rep.Conn.Where("id = ?", schID).Updates(&scheduleUpdate)

	if result.Error != nil {
		return schedules.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(scheduleUpdate), nil
}

func (rep *MysqlScheduleRepository) ScheduleByID(id int) (schedules.Domain, error) {

	var schedule Schedules

	result := rep.Conn.Where("id = ?", id).First(&schedule)

	if result.Error != nil {
		return schedules.Domain{}, result.Error
	}

	return toDomain(schedule), nil
}

func (rep *MysqlScheduleRepository) Delete(id int) (string, error) {
	rec := Schedules{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Schedule has been delete", nil
}
