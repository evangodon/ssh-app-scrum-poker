package main

import (
	"fmt"
	"time"
)

type roomLog struct {
	log         string
	clearBefore bool
}

func faint(s string) string {
	return style().Faint(true).Render(s)
}

func newRoomLog(msg string) roomLog {
	ts := fmt.Sprintf("[%s]", time.Now().Format(time.Kitchen))
	log := fmt.Sprintf("%s %s", faint(ts), msg)

	return roomLog{
		log:         log,
		clearBefore: false,
	}
}
