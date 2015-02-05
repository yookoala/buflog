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

// normal logger function
func (lr *Logger) Printf(msg string, params ...interface{}) {
	lr.logs = append(lr.logs, NewLog(msg, params...))
}

// playback
func (lr *Logger) Play() {
	for _, l := range lr.logs {
		l.Play()
	}
}
