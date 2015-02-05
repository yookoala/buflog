package buflog

import (
	"log"
	"time"
)

// log entry
type L struct {
	t time.Time
	m string
	p []interface{}
}

// playback the log entry
func (l *L) Play() {
	log.Printf(l.m, l.p...)
}

func NewLog(msg string, p ...interface{}) *L {
	l := L{
		t: time.Now(),
		m: msg,
		p: p,
	}
	return &l
}
