package resozyme

import "testing"

func TestNilLogger(t *testing.T) {
	logger := &NilLogger{}
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
	logger.Fatal("Fatal")
	logger.Debugf("%s", "Debug")
	logger.Infof("%s", "Info")
	logger.Warnf("%s", "Warn")
	logger.Errorf("%s", "Error")
	logger.Fatalf("%s", "Fatal")
}
