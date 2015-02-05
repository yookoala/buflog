package buflog

func New() *Logger {
	return &Logger{
		logs: make([]*L, 0, 100),
	}
}

// logger
type Logger struct {
	logs []*L
}

// implements standard logger function
func (lr *Logger) Print(msg string) {
	lr.Printf(msg)
}

// implements standard logger function
func (lr *Logger) Printf(msg string, params ...interface{}) {
	lr.logs = append(lr.logs, NewLog(msg, params...))
}

// implements standard logger function
func (lr *Logger) Panic(msg string) {
	lr.Panicf(msg)
}

// implements standard logger function
func (lr *Logger) Panicf(msg string, params ...interface{}) {
	lr.logs = append(lr.logs, NewPanic(msg, params...))
	lr.Play() // play back immediately
}

// implements standard logger function
func (lr *Logger) Fatal(msg string) {
	lr.Fatalf(msg)
}

// implements standard logger function
func (lr *Logger) Fatalf(msg string, params ...interface{}) {
	lr.logs = append(lr.logs, NewFatal(msg, params...))
	lr.Play() // play back immediately
}

// playback
func (lr *Logger) Play() {
	for _, l := range lr.logs {
		l.Play()
	}
}
