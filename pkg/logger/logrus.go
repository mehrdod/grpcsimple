package logger

import "github.com/sirupsen/logrus"

type LogrusLogger struct{}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}

func (l *LogrusLogger) Info(msg ...interface{}) {
	logrus.Info(msg...)
}
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
