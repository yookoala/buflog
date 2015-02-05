package buflog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	// setup log to log to a buffer
	b := &bytes.Buffer{}
	log.SetOutput(b)

	// log something to logger and get result from buffer
	ts := time.Now().Format("2006/01/02 15:04:05")
	l := New()
	l.Printf("testing %d %d %s", 1, 2, "3")
	l.Printf("testing %d %d %s", 4, 5, "6")
	l.Printf("testing %d %d %s", 7, 8, "9")
	l.Play()
	r := b.String()

	// expected result
	e := fmt.Sprintf("%s testing 1 2 3\n"+
		"%s testing 4 5 6\n"+
		"%s testing 7 8 9\n",
		ts, ts, ts)

	// examine result
	if r != e {
		t.Errorf("Failed to playback logger.\nExptect: \"%s\"\nGetting: \"%s\"",
			e, r)
	}
	log.SetOutput(os.Stdout)
}
