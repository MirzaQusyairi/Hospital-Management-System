package request

import (
	"Hospital-Management-System/business/sesbooks"
)

type Sesbooks struct {
	IDPatient         int `json:"id_patient"`
	IDSessionSchedule int `json:"id_session_schedule"`
	PatientQueue      int `json:"patient_queue"`
}

func (req *Sesbooks) ToDomain() *sesbooks.Domain {
	return &sesbooks.Domain{
		IDPatient:         req.IDPatient,
		IDSessionSchedule: req.IDSessionSchedule,
		PatientQueue:      req.PatientQueue,
	}
}
