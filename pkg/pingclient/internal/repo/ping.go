package repo

import "time"

type Ping struct {
	ID        int64     `json:"id"`
	StartTime time.Time `json:"start_time"`
	Start     int64     `json:"start"`
	EndTime   time.Time `json:"end_time"`
	End       int64     `json:"end"`

	Delay int64 `json:"delay"`

	Error error `json:"error"`
}

type PingRepo interface {
	InsertRecord(*Ping) error
}
