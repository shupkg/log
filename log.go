package log

// Logger is a generic logging interface
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

// Provider provider logger
type Provider interface {
	Prefix(prefix string) Logger
}
