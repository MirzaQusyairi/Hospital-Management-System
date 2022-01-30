package request

import "Hospital-Management-System/business/facilities"

type Facilities struct {
	Name     string `json:"name"`
	Queue    int    `json:"queue"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}

func (req *Facilities) ToDomain() *facilities.Domain {
	return &facilities.Domain{
		Name:     req.Name,
		Queue:    req.Queue,
		Location: req.Location,
		Capacity: req.Capacity,
	}
}
