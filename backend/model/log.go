package model

import "time"

type Log struct {
	ID       *int      `json:"id"`
	Quadrant string    `json:"quadrant"`
	Ctime    time.Time `json:"ctime"`
	Location *string   `json:"location"`
	Atype    string    `json:"atype"`
	Title    string    `json:"title"`
	Detail   *string   `json:"detail"`
	Review   *string   `json:"review"`
}
