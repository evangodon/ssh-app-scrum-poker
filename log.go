package main

import (
	"fmt"
	"time"
)

type roomLog struct {
	log string
}

func newRoomLog(msg string) roomLog {
	ts := time.Now().Format("3:01:05")
	log := fmt.Sprintf("[%s] %s", ts, msg)

	return roomLog{
		log: log,
	}
}
