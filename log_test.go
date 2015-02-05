package buflog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

// testing log without parameters
func TestNewLog0(t *testing.T) {
	now := time.Now()
	l := NewLog("testing")

	// test log time
	if l.t.Unix() != now.Unix() && l.t.Unix() > now.Unix()+1 {
		t.Errorf("The timestamp of log is not within acceptable range")
	}

	// test log content
	s := fmt.Sprintf(l.m, l.p...)
	if s != "testing" {
		t.Errorf("The log message generated is not as expected")
	}
}

// testing log with parameters
func TestNewLog1(t *testing.T) {
	now := time.Now()
	l := NewLog("testing %d %d %s", 1, 2, "3")

	// test log time
	if l.t.Unix() != now.Unix() && l.t.Unix() > now.Unix()+1 {
		t.Errorf("The timestamp of log is not within acceptable range")
	}

	// test log content
	s := fmt.Sprintf(l.m, l.p...)
	if s != "testing 1 2 3" {
		t.Errorf("The log message generated is not as expected")
	}
}

// test to play back a log entry
func TestLogPlay(t *testing.T) {
	b := &bytes.Buffer{}
	log.SetOutput(b)
	e := fmt.Sprintf("%s testing %d %d %s\n",
		time.Now().Format("2006/01/02 15:04:05"),
		1, 2, "3")
	l := NewLog("testing %d %d %s", 1, 2, "3")
	l.Play()
	r := b.String()
	if r != e {
		t.Errorf("Failed to playback log.\nExptect: \"%s\"\nGetting: \"%s\"",
			e, r)
	}
	log.SetOutput(os.Stdout)
}
