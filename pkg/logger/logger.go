package logger

type Logger interface {
	Info(msg ...interface{})
	Errorf(format string, args ...interface{})
}
