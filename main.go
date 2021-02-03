package main

import (
	"math"
	"time"
)


type SW struct {
	bytes int64
	startTime time.Time
}

func (s *SW) SetStartTime() {
	s.startTime = time.Now()
}

func (s *SW) Read(n []byte) {
	s.bytes += int64(len(n))
}
func (s *SW) Write(n []byte) {
	s.bytes += int64(len(n))
}

func (s *SW) GetSpeed() float64 {
	return float64(s.bytes)/time.Since(s.startTime).Seconds()
}
func (s *SW) GetSpeedRound() int64 {
	return int64(math.Round(s.GetSpeed()))
}
