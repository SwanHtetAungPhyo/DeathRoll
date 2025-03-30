package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func init() {
	// Initialize the logger instance
	logger = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
			ForceColors:   true,
		},
	}
}

func NewLogger(name string) *logrus.Logger {
	return logger.WithField("service", name).Logger
}
