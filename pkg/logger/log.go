package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogInterface interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	SetLevel(level logrus.Level)
}

type log struct {
	logger *logrus.Logger
}

func NewLog() LogInterface {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	return &log{
		logger: logger,
	}
}

func (l *log) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *log) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *log) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *log) SetLevel(level logrus.Level) {
	l.logger.SetLevel(level)
}
