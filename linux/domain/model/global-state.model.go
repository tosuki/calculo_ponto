package model

type GlobalState struct {
	StartTime    int64
	JourneyHours int

	CurrentClock *Clock
	Clocks       []*Clock

	Config
}

func NewGlobalState(startPeriod int64, journeyHours int, clocks []*Clock, config *Config) *GlobalState {
	return &GlobalState{
		StartTime:    startPeriod,
		JourneyHours: journeyHours,
		Clocks:       clocks,
		CurrentClock: nil,
		Config:       *config,
	}
}
