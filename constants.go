package golang_cmd_loading

import "time"

// Delay time between every change of loading icon in milliseconds
const (
	DelayRateDefault = 180 * time.Millisecond
	DelayRateSlow    = 250 * time.Millisecond
	DelayRateFast    = 130 * time.Millisecond
)
