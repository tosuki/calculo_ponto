package src

type AppClock struct {
	Hours   int
	Minutes int
	Seconds int
}

func NewAppClock(hours, minutes, seconds int) *AppClock {
	return &AppClock{
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}
