package golang_cmd_loading

import "time"

// SetDelayRate Set the delay time between every frame of the loading animation, use constants DelayRateDefault,
// DelayRateSlow, DelayRateFast or just give one.
func SetDelayRate(delayRateNew time.Duration) {
	delayRate = delayRateNew
}

// WithLoading Start a job in a new goroutine, and display a loading icon during the execution of the job.
func WithLoading(job func(cancelLoading func())) {
	WithLoadingMessage(job, "")
}

// WithLoadingMessage Start a job in a new goroutine, and display a loading icon with the provided message
// during the execution of the job.
func WithLoadingMessage(job func(cancelLoading func()), message string) {
	var (
		cancelLoading = false
		joinChannel   = make(chan int)
	)
	// Start the job first
	go func() {
		job(func() {
			notifyNoLoading(&cancelLoading, delayRate)
		})
		// Notify main goroutine all the jobs are done, it's time to end
		joinChannel <- 1
	}()
	// Display loading status
	displayLoadingWithMessage(message, delayRate, &cancelLoading)
	<-joinChannel
}
