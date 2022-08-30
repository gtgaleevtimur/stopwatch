package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	StartTime     time.Time
	StopPointList []time.Time
}

func (s *Stopwatch) Start() {
	s.StartTime = time.Now()
	s.StopPointList = nil
}

func (s *Stopwatch) SavePoint() {
	s.StopPointList = append(s.StopPointList, time.Now())
}

func (s *Stopwatch) GetResults() (result []time.Duration) {
	for _, stops := range s.StopPointList {
		result = append(result, stops.Sub(s.StartTime))
	}
	return
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SavePoint()

	time.Sleep(500 * time.Millisecond)
	sw.SavePoint()

	time.Sleep(300 * time.Millisecond)
	sw.SavePoint()

	fmt.Println(sw.GetResults())
}
