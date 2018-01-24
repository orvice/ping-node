package main

import (
	"github.com/sparrc/go-ping"
)

func Ping(addr string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return nil, err
	}
	pinger.Count = 10
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()
	return stats, nil
}
