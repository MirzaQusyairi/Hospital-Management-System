package request

import (
	"Hospital-Management-System/business/seschedules"
)

type Seschedules struct {
	IDSchedule string `json:"id_schedule"`
	IDDoctor   string `json:"id_doctor"`
	IDFacilty  string `json:"id_facilty"`
}

func (req *Seschedules) ToDomain() *seschedules.Domain {
	return &seschedules.Domain{

		IDSchedule: req.IDSchedule,
		IDDoctor:   req.IDDoctor,
		IDFacilty:  req.IDFacilty,
	}
}
