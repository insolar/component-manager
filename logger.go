package component

import "log"

// Logger interface provides methods for debug logging
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

// DefaultLogger logs to std out
type DefaultLogger struct{}

func (l *DefaultLogger) Debug(v ...interface{}) {
	log.Println(v...)
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// NoLogger skips log messages
type NoLogger struct{}

func (l *NoLogger) Debug(v ...interface{}) {
}

func (l *NoLogger) Debugf(format string, v ...interface{}) {
}
