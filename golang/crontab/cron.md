```
package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 29 15 * * *", func() {
		fmt.Println("Every hour on the half hour")
	})
	c.Start()
	select {}
}

func packCronSchedule(endTime int64) string {
	tm := time.Unix(endTime, 0)
	month := int(tm.Month())
	day := tm.Day()
	hour := tm.Hour()
	minute := tm.Minute()
	second := tm.Second()
	schedule := fmt.Sprintf("%v %v %v %v %v *", second, minute, hour, day, month)

	return schedule
}
```
