package model

import (
	"sync"
	"time"
)

type Stats struct {
	statsLock sync.RWMutex
	Requests  int       `json:"requests"`
	Keys      int       `json:"keys"`
	StartTime time.Time `json:"-"`
}

func NewStats() *Stats {
	return &Stats{
		StartTime: time.Now(),
	}
}

func (s *Stats) IncRequests() {
	s.statsLock.Lock()
	defer s.statsLock.Unlock()
	s.Requests++
}

func (s *Stats) SetKeys(count int) {
	s.statsLock.Lock()
	defer s.statsLock.Unlock()
	s.Keys = count
}

func (s *Stats) Snapshot() map[string]interface{} {
	s.statsLock.RLock()
	defer s.statsLock.RUnlock()

	return map[string]interface{}{
		"requests":       s.Requests,
		"keys":           s.Keys,
		"uptime_seconds": int(time.Since(s.StartTime).Seconds()),
	}
}
