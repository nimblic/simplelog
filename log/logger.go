package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

const (
	LevelDebug   = slog.LevelDebug
	LevelInfo    = slog.LevelInfo
	LevelWarn    = slog.LevelWarn
	LevelError   = slog.LevelError
	LevelNotice  = slog.Level(2)
	LevelVerbose = slog.Level(-5)
)

var LevelNames = map[slog.Leveler]string{
	LevelVerbose: "VERBOSE",
	LevelNotice:  "NOTICE",
}

type MedLog struct {
	*slog.Logger
}

var loggerSingleton *MedLog
var once sync.Once
var usePrettyPrintLogger = false

func Init(level slog.Level) *MedLog {
	var defaultLogLevel slog.Level
	defaultLogLevel = level
	// switch strings.ToLower(level) {
	// case "verbose":
	// 	defaultLogLevel = LevelVerbose
	// case "debug":
	// 	defaultLogLevel = slog.LevelDebug
	// case "info":
	// 	defaultLogLevel = slog.LevelInfo
	// case "notice":
	// 	defaultLogLevel = LevelNotice
	// case "warning":
	// 	defaultLogLevel = slog.LevelWarn
	// case "error":
	// 	defaultLogLevel = slog.LevelError
	// }
	once.Do(func() {
		opts := &slog.HandlerOptions{
			// AddSource: true,
			Level: defaultLogLevel,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.LevelKey {
					level := a.Value.Any().(slog.Level)
					levelLabel, exists := LevelNames[level]
					if !exists {
						levelLabel = level.String()
					}

					a.Value = slog.StringValue(levelLabel)
				}

				return a
			},
		}
		var logger *slog.Logger
		if usePrettyPrintLogger {
			logger = slog.New(NewHandler(opts))
		} else {
			logger = slog.New(slog.NewJSONHandler(os.Stderr, opts))
		}
		loggerSingleton = &MedLog{logger}
	})

	return loggerSingleton
}

func Verbose(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelVerbose, msg, args...)
	}
}

func Verbosef(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelVerbose, fmt.Sprintf(format, args...))
	}
}

func Debug(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Debug(msg, args...)
	}
}

func Debugf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Debug(fmt.Sprintf(format, args...))
	}
}

func Info(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Info(msg, args...)
	}
}
func Infof(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Info(fmt.Sprintf(format, args...))
	}
}

func Notice(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelNotice, msg, args...)
	}
}

func Noticef(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelNotice, fmt.Sprintf(format, args...))
	}
}

func Warning(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Warn(msg, args...)
	}
}
func Warningf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Warn(fmt.Sprintf(format, args...))
	}
}

func Error(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Error(msg, args...)
	}
}

func Errorf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Error(fmt.Sprintf(format, args...))
	}
}
