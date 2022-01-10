package request

import (
	"Hospital-Management-System/business/prescriptions"
)

type Prescriptions struct {
	MedicineName     string `json:"medicine_name"`
	MedicationRules  string `json:"medication_rules"`
	IDPatient        string `json:"id_patient"`
	IDDoctor         string `json:"id_doctor"`
	IDSessionBooking string `json:"id_sessionbooking"`
}

func (req *Prescriptions) ToDomain() *prescriptions.Domain {
	return &prescriptions.Domain{
		MedicineName:     req.MedicineName,
		MedicationRules:  req.MedicationRules,
		IDPatient:        req.IDPatient,
		IDDoctor:         req.IDDoctor,
		IDSessionBooking: req.IDSessionBooking,
	}
}
