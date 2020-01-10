package game

import "time"

type Stats struct {
	Clicks_ int       `json:"clicks"`
	Start_  time.Time `json:"start"`
}

func (s *Stats) Click() {
	s.Clicks_++
}

func (s *Stats) Clicks() int {
	return s.Clicks_
}

func (s *Stats) Time() int {
	return s.Start_.Second()
}
