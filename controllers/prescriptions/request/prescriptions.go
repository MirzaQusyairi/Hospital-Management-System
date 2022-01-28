package request

import (
	"Hospital-Management-System/business/prescriptions"
)

type Prescriptions struct {
	MedicineName     string `json:"medicine_name"`
	MedicationRules  string `json:"medication_rules"`
	PatientsID       int    `json:"id_patient"`
	DoctorID         int    `json:"id_doctor"`
	SessionBookingID int    `json:"id_sessionbooking"`
}

func (req *Prescriptions) ToDomain() *prescriptions.Domain {
	return &prescriptions.Domain{
		MedicineName:     req.MedicineName,
		MedicationRules:  req.MedicationRules,
		PatientsID:       req.PatientsID,
		DoctorID:         req.DoctorID,
		SessionBookingID: req.SessionBookingID,
	}
}
