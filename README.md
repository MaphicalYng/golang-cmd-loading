# golang-cmd-loading

Show loading status easily in command-line app via Golang.

## Install

Run the command in your module:

```bash
go get -u github.com/MaphicalYng/golang-cmd-loading@latest
```

## Usage

Call the entry function with an anonymous function which contains your 
job logic:

```go
package main

import (
	"fmt"
	gcl "github.com/MaphicalYng/golang-cmd-loading"
	"os"
	"time"
)

func main() {
	gcl.WithLoading(func(cancelLoading func()) {
		// 1.Do your job here
		result := fmt.Sprintf("Parameters: %v", os.Args[1:])
		time.Sleep(5 * time.Second) // To see loading icon clearly, not required

		// 2.Call function `cancelLoading` to cancel loading status,
		// so that you can print content which the job generated
		cancelLoading()
		
		// 3.Print content if you need
		fmt.Println(result)
	})
}
```

You can also show loading status with customized message:

```go
package main

import (
	"fmt"
	gcl "github.com/MaphicalYng/golang-cmd-loading"
	"os"
	"time"
)

func main() {
	// Change WithLoading to WithLoadingMessage
	gcl.WithLoadingMessage(func(cancelLoading func()) {
		// 1.Do your job here
		result := fmt.Sprintf("Parameters: %v", os.Args[1:])
		time.Sleep(5 * time.Second) // To see loading icon clearly, not required

		// 2.Call function `cancelLoading` to cancel loading status,
		// so that you can print content which the job generated
		cancelLoading()

		// 3.Print content if you need
		fmt.Println(result)
	}, "Computing...") // Add your message as the second parameter
}
```

You can adjust the speed of rotation of the loading icon:

```go
gcl.SetDelayRate(gcl.DelayRateSlow)
```

The `DelayRate` is period of the time between every frame of rotation of the icon.

You can use the given values, including:
* `DelayRateDefault`(180 ms, default)
* `DelayRateSlow`(250 ms)
* `DelayRateFast`(130 ms)

Or you can just set a customized value of type `time.Duration`.

## Output

Loading with a customized message: "Computing..."

![](https://github.com/MaphicalYng/readme-res/raw/main/github_readme_golang_cmd_loading.gif)
