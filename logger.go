package resozyme

// Logger is a logger interface.
// This interface is compatible with zap.SugaredLogger.
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

// NilLogger is a nil logger implementation.
type NilLogger struct {
}

// Debug implements resource.Logger.
func (logger *NilLogger) Debug(args ...interface{}) {}

// Info implements resource.Logger.
func (logger *NilLogger) Info(args ...interface{}) {}

// Warn implements resource.Logger.
func (logger *NilLogger) Warn(args ...interface{}) {}

// Error implements resource.Logger.
func (logger *NilLogger) Error(args ...interface{}) {}

// Fatal implements resource.Logger.
func (logger *NilLogger) Fatal(args ...interface{}) {}

// Debugf implements resource.Logger.
func (logger *NilLogger) Debugf(template string, args ...interface{}) {}

// Infof implements resource.Logger.
func (logger *NilLogger) Infof(template string, args ...interface{}) {}

// Warnf implements resource.Logger.
func (logger *NilLogger) Warnf(template string, args ...interface{}) {}

// Errorf implements resource.Logger.
func (logger *NilLogger) Errorf(template string, args ...interface{}) {}

// Fatalf implements resource.Logger.
func (logger *NilLogger) Fatalf(template string, args ...interface{}) {}
