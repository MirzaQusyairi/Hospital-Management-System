package request

import (
	"Hospital-Management-System/business/seschedules"
)

type Seschedules struct {
	IDSchedule int `json:"id_schedule"`
	IDDoctor   int `json:"id_doctor"`
	IDFacilty  int `json:"id_facilty"`
}

func (req *Seschedules) ToDomain() *seschedules.Domain {
	return &seschedules.Domain{

		IDSchedule: req.IDSchedule,
		IDDoctor:   req.IDDoctor,
		IDFacilty:  req.IDFacilty,
	}
}
