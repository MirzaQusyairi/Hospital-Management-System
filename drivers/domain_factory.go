package drivers

import (
	adminDomain "Hospital-Management-System/business/admins"
	doctorDomain "Hospital-Management-System/business/doctors"
	faciltyDomain "Hospital-Management-System/business/facilities"
	nurseDomain "Hospital-Management-System/business/nurses"
	patientDomain "Hospital-Management-System/business/patients"
	prescriptionDomain "Hospital-Management-System/business/prescriptions"
	scheduleDomain "Hospital-Management-System/business/schedules"

	adminDB "Hospital-Management-System/drivers/databases/admins"
	doctorDB "Hospital-Management-System/drivers/databases/doctors"
	faciltyDB "Hospital-Management-System/drivers/databases/facilities"
	nurseDB "Hospital-Management-System/drivers/databases/nurses"
	patientDB "Hospital-Management-System/drivers/databases/patients"
	prescriptionDB "Hospital-Management-System/drivers/databases/prescriptions"
	scheduleDB "Hospital-Management-System/drivers/databases/schedules"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}

func NewDoctorRepository(conn *gorm.DB) doctorDomain.Repository {
	return doctorDB.NewMysqlDoctorRepository(conn)
}

func NewPatientRepository(conn *gorm.DB) patientDomain.Repository {
	return patientDB.NewMysqlPatientRepository(conn)
}

func NewNurseRepository(conn *gorm.DB) nurseDomain.Repository {
	return nurseDB.NewMysqlNurseRepository(conn)
}

func NewFaciltyRepository(conn *gorm.DB) faciltyDomain.Repository {
	return faciltyDB.NewMysqlFaciltyRepository(conn)
}

func NewScheduleRepository(conn *gorm.DB) scheduleDomain.Repository {
	return scheduleDB.NewMysqlScheduleRepository(conn)
}

func NewPrescriptionRepository(conn *gorm.DB) prescriptionDomain.Repository {
	return prescriptionDB.NewMysqlPrescriptionRepository(conn)
}
