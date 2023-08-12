
# Go Scheduler

Package for register scheduler command in your Go project


## Installation

Run command below to install Go Scheduler

```bash
  go get github.com/dnridwn/go-scheduler
```
    
## Usage/Examples

```go
import (
    goscheduler "github.com/dnridwn/go-scheduler"
)

func SchedulerKernel() {
    scheduler := goscheduler.NewScheduler()
    defer scheduler.Run()

    scheduler.Add(context.Background(), "* * * * * *", func() {
        fmt.Println("Run this every second")
    })

    scheduler.Add(context.Background(), "0 0 * * *", func() {
        fmt.Println("Run this every day at 00:00")
    })
}
```

