package model

type LunchPeriods struct {
	Start int64
	End   int64
}

type GlobalState struct {
	StartPeriod  int64
	LunchPeriods []*LunchPeriods
	JourneyHours int

	Config
}

func NewGlobalState(startPeriod int64, journeyHours int, lunchPeriods []*LunchPeriods, config *Config) *GlobalState {
	return &GlobalState{
		StartPeriod:  startPeriod,
		LunchPeriods: lunchPeriods,
		JourneyHours: journeyHours,
		Config:       *config,
	}
}
