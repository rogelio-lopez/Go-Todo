package main

import (
	"time"
)

func timestamp() string {
	var timestamp time.Time = time.Now()
	return timestamp.Format(time.Stamp)
}

func argText(args []string) string {
	var txt string = ""
	for i, s := range args {
		if i < 1 {
			txt = s
		} else {
			txt = txt + " " + s
		}
	}
	return txt
}
