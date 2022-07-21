package main

import (
	"fmt"
	"time"
)

type roomLog struct {
	log         string
	clearBefore bool
}

func newRoomLog(msg string) roomLog {
	ts := time.Now().Format("3:01")
	log := fmt.Sprintf("[%s] %s", ts, msg)

	return roomLog{
		log:         log,
		clearBefore: false,
	}
}
