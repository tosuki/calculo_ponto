package model

import "time"

type Clock struct {
	Time        int64
	EndTime     int64
	Description string
}

func (this *Clock) End() {
	this.EndTime = time.Now().Unix()
}

func NewClock(time int64, description string) *Clock {
	return &Clock{
		Time:        time,
		Description: description,
		EndTime:     -1,
	}
}
