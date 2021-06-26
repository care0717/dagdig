package openapi

import (
	"fmt"
	"time"
)

func ToCreatedFormat(d time.Duration) string {
	sec := d - d.Truncate(time.Minute)
	min := d - d.Truncate(time.Hour) - sec
	hour := d - min - sec
	return fmt.Sprintf("%02.0f:%02.0f:%02.0f", hour.Hours(), min.Minutes(), sec.Seconds())
}
