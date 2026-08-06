package src

import (
	"fmt"
	"log"
	"time"
)

type TimerMode = int

var (
	TimerModeClock      TimerMode = 0
	TimerModeRegressive TimerMode = 1
	TimerModePaused     TimerMode = 2
)

type Timer struct {
	Mode     TimerMode
	LastMode TimerMode
	IsPaused bool

	PausedAt     time.Time
	Journey      time.Duration
	StartedAt    time.Time
	IntervalTime time.Duration
}

func (t *Timer) Pause() {
	if !t.IsPaused {
		t.IsPaused = true
		t.PausedAt = time.Now()
		t.LastMode = t.Mode
		t.Mode = TimerModePaused
	}
}

func (t *Timer) Resume() {
	if t.IsPaused {
		t.IsPaused = false

		now := time.Now()
		intervalTime := now.Sub(t.PausedAt)

		t.IntervalTime += intervalTime
		t.Mode = t.LastMode
	}
}

func (t *Timer) SetMode(mode TimerMode) {
	if mode < 0 || mode > 2 {
		log.Fatalf("Mode não suportado pelo timer %d\n", mode)
	}

	t.Mode = mode
}

func (t *Timer) GetOutput() string {
	switch t.Mode {
	case TimerModeClock:
		now := time.Now()
		return now.Format("15:04:05")

	case TimerModeRegressive:
		now := time.Now()
		elapsed := now.Sub(t.StartedAt)
		left := t.Journey - (elapsed - t.IntervalTime)

		return formatDuration(left)

	case TimerModePaused:
		if t.LastMode == TimerModeRegressive {
			elapsed := t.PausedAt.Sub(t.StartedAt)
			left := t.Journey - (elapsed - t.IntervalTime)
			return formatDuration(left)
		}

		return t.PausedAt.Format("15:04:05")
	}

	return ""
}

func formatDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}

	d = d.Round(time.Second)

	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func NewTimer(mode TimerMode, startedAt time.Time, journey time.Duration) *Timer {
	return &Timer{
		StartedAt:    startedAt,
		Journey:      journey,
		Mode:         mode,
		IsPaused:     false,
		IntervalTime: 0,
	}
}
