package request

import (
	"Hospital-Management-System/business/schedules"
	"time"
)

type Schedules struct {
	Day   string    `json:"day"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (req *Schedules) ToDomain() *schedules.Domain {
	return &schedules.Domain{

		Day:   req.Day,
		Start: req.Start,
		End:   req.End,
	}
}
