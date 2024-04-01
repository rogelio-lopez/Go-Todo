package main

import (
	"time"
)

func timestamp() string {
	var timestamp time.Time = time.Now()
	return timestamp.Format(time.Stamp)
}
