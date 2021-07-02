package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	initial *zap.Logger

	// NOTE: sugared presents *zap.Logger clone, that is any changes maden
	// to the *zap.Logger since initialization would not be reflected at
	// *zap.SugaredLogger
	sugared *zap.SugaredLogger
}

func New() *Logger {
	var (
		writeSyncer = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			// TODO: add output to log file
		)
		checkIsLevelLoggable = zap.LevelEnablerFunc(
			func(l zapcore.Level) bool {
				return true
			},
		)
		consoleEncoder = defaultConsoleEncoder()

		core = zapcore.NewCore(consoleEncoder, writeSyncer, checkIsLevelLoggable)
	)
	logger := zap.New(core, zap.ErrorOutput(os.Stderr), zap.AddCaller()) // TODO: options

	return &Logger{
		initial: logger,
		sugared: logger.Sugar(),
	}
}

func defaultConsoleEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stack_trace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *Logger) Initial() *zap.Logger {
	return l.initial
}
