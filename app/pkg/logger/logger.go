package logger

type Logger interface {
	Info(args ...interface{})
	Fatal(args ...interface{})
}
