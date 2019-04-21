package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.SugaredLogger
}

func NewLogger() *Logger {
	logger := &Logger{}
	logger.log = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			os.Stdout,
			zapcore.DebugLevel,
		),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel)).Sugar()

	return logger
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.log.Debug(args)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.log.Info(args)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.log.Warn(args)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.log.Error(args)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.log.Fatal(args)
}
