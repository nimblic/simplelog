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

func PersistLog(_ bool) {
	// This function to make it easier to port existing code
}

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

func VerboseSlog(ctx context.Context, msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelVerbose, msg, args...)
	}
}

func VerboseDebugf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelVerbose, fmt.Sprintf(format, args...))
	}
}

func Verbose(format string) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelVerbose, fmt.Sprint(format))
	}
}

func Verbosef(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelVerbose, fmt.Sprintf(format, args...))
	}
}

func DebugSlog(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Debug(msg, args...)
	}
}

func Debug(format string) {
	if loggerSingleton != nil {
		loggerSingleton.Debug(fmt.Sprint(format))
	}
}

func Debugf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Debug(fmt.Sprintf(format, args...))
	}
}

func InfoSlog(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Info(msg, args...)
	}
}

func Info(format string) {
	if loggerSingleton != nil {
		loggerSingleton.Info(fmt.Sprint(format))
	}
}

func Infof(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Info(fmt.Sprintf(format, args...))
	}
}

func NoticeSlog(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelNotice, msg, args...)
	}
}

func Notice(format string) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelNotice, fmt.Sprint(format))
	}
}

func Noticef(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Log(context.TODO(), LevelNotice, fmt.Sprintf(format, args...))
	}
}

func WarningSlog(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Warn(msg, args...)
	}
}

func Warn(format string) {
	if loggerSingleton != nil {
		loggerSingleton.Warn(fmt.Sprint(format))
	}
}

func Warnf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Warn(fmt.Sprintf(format, args...))
	}
}

func Warningf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Warn(fmt.Sprintf(format, args...))
	}
}

func ErrorSlog(msg string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Error(msg, args...)
	}
}

func Error(format string) {
	if loggerSingleton != nil {
		loggerSingleton.Error(fmt.Sprint(format))
	}
}

func Errorf(format string, args ...any) {
	if loggerSingleton != nil {
		loggerSingleton.Error(fmt.Sprintf(format, args...))
	}
}
