package log

import (
	"github.com/navono/go-logger"
	"log"
)

func init() {
	config := logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logger.Debug,
		ConsoleJSONFormat: true,
		EnableFile:        true,
		FileLevel:         logger.Info,
		FileJSONFormat:    true,
		FileLocation:      "log.log",
	}
	zapLogger, err := logger.NewLogger(config, logger.InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	zapLogger.Infof("Starting with zap")

	contextLogger := zapLogger.WithFields(logger.Fields{"key1": "value1"})
	contextLogger.Infof("Zap is awesome")

	logrusLogger, err := logger.NewLogger(config, logger.InstanceLogrusLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	logrusLogger.Debugf("Starting with logrus")
	logrusLogger.Infof("Logrus is awesome")

	contextLogger = logrusLogger.WithFields(logger.Fields{"key1": "value1"})
	contextLogger.Debugf("Starting with context logrus")
	contextLogger.Infof("Logrus is awesome")
}
