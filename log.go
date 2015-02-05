package buflog

import (
	"log"
	"time"
)

const (
	NORMAL = iota
	PANIC  = iota
	FATAL  = iota
)

// log entry
type L struct {
	t time.Time
	m string
	p []interface{}
	f int
}

// playback the log entry
func (l *L) Play() {
	switch l.f {
	case NORMAL:
		log.Printf(l.m, l.p...)
	case PANIC:
		log.Panicf(l.m, l.p...)
	case FATAL:
		log.Fatalf(l.m, l.p...)
	}
}

func NewLog(msg string, p ...interface{}) *L {
	l := L{
		t: time.Now(),
		m: msg,
		p: p,
	}
	return &l
}

func NewPanic(msg string, p ...interface{}) *L {
	l := NewLog(msg, p...)
	l.f = PANIC
	return l
}

func NewFatal(msg string, p ...interface{}) *L {
	l := NewLog(msg, p...)
	l.f = FATAL
	return l
}
