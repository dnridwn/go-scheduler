
# Go Scheduler

Package for register scheduler in your Go project

#### THIS PACKAGE IS STILL UNDER DEVELOPMENT AND NOT READY FOR PRODUCTION USE!
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

    scheduler.Add("* * * * * *", func() {
        fmt.Println("Run this every second")
    })

    scheduler.Add("0 0 * * *", func() {
        fmt.Println("Run this every day at 00:00")
    })
}
```

