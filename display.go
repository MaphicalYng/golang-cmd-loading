package golang_cmd_loading

import (
	"fmt"
	"strings"
	"time"
)

var (
	delayRate = DelayRateDefault
)

// displayLoadingWithMessage Display a loading icon forever until the `notify` is changed to true
func displayLoadingWithMessage(message string, delayRate time.Duration, notifyNoLoadingBool *bool) {
	message = " " + message
	for {
		if *notifyNoLoadingBool {
			break
		}
		fmt.Printf("\r|%s", message)
		if *notifyNoLoadingBool {
			break
		}
		time.Sleep(delayRate)
		if *notifyNoLoadingBool {
			break
		}
		fmt.Printf("\r/%s", message)
		if *notifyNoLoadingBool {
			break
		}
		time.Sleep(delayRate)
		if *notifyNoLoadingBool {
			break
		}
		fmt.Printf("\r-%s", message)
		if *notifyNoLoadingBool {
			break
		}
		time.Sleep(delayRate)
		if *notifyNoLoadingBool {
			break
		}
		fmt.Printf("\r\\%s", message)
		if *notifyNoLoadingBool {
			break
		}
		time.Sleep(delayRate)
	}
}

func clearCurrentLineSub(messageLength int) {
	fmt.Printf("\r  %s\r", strings.Repeat(" ", messageLength))
}

// notifyNoLoading Cancel loading status
func notifyNoLoading(notifyNoLoadingBool *bool, delayRate time.Duration, messageLength int) {
	*notifyNoLoadingBool = true
	time.Sleep(delayRate + 50*time.Millisecond)
	clearCurrentLineSub(messageLength)
}
