package request

import (
	"Hospital-Management-System/business/schedules"
)

type Schedules struct {
	Day   string `json:"day"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (req *Schedules) ToDomain() *schedules.Domain {
	return &schedules.Domain{

		Day:   req.Day,
		Start: req.Start,
		End:   req.End,
	}
}
