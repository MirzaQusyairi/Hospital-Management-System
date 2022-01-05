package drivers

import (
	adminDomain "Hospital-Management-System/business/admins"
	doctorDomain "Hospital-Management-System/business/doctors"
	faciltyDomain "Hospital-Management-System/business/facilities"
	patientDomain "Hospital-Management-System/business/patients"

	nurseDomain "Hospital-Management-System/business/nurses"
	adminDB "Hospital-Management-System/drivers/databases/admins"
	doctorDB "Hospital-Management-System/drivers/databases/doctors"
	faciltyDB "Hospital-Management-System/drivers/databases/facilities"
	patientDB "Hospital-Management-System/drivers/databases/patients"

	nurseDB "Hospital-Management-System/drivers/databases/nurses"

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
