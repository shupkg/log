package log

// Logger is a generic logging interface
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

// LoggerProvider provider logger
type LoggerProvider interface {
	Prefix(prefix string) Logger
}
