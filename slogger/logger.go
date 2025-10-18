package slogger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

// Global singleton instance using atomic pointer for lock-free access
var (
	instance atomic.Pointer[slog.Logger]
	once     sync.Once
)

const (
	LevelNotice  = slog.Level(2)
	LevelVerbose = slog.Level(-5)
)

//	var LevelNames = map[slog.Leveler]string{
//		LevelVerbose: "VERBOSE",
//		LevelNotice:  "NOTICE",
//	}
var (
	verboseValue = slog.StringValue("VERBOSE")
	noticeValue  = slog.StringValue("NOTICE")
)

// Config holds logger configuration
type Config struct {
	Level               string `json:"level" yaml:"level"`
	Format              string `json:"format" yaml:"format"` // "json" or "text"
	AddSource           bool   `json:"add_source" yaml:"add_source"`
	SupportCustomLevels bool   `json:"support_custom_levels" yaml:"support_custom_levels"`
}

// Initialize sets up the logger singleton (call once at startup)
func Initialize(cfg Config) error {
	var initErr error

	once.Do(func() {
		level, err := parseLevel(cfg.Level)
		if err != nil {
			initErr = fmt.Errorf("invalid log level: %w", err)
			return
		}

		opts := &slog.HandlerOptions{
			Level:     level,
			AddSource: cfg.AddSource,
		}
		if cfg.SupportCustomLevels {
			opts.ReplaceAttr = func(_ []string, a slog.Attr) slog.Attr {
				if a.Key != slog.LevelKey {
					return a
				}
				if level, ok := a.Value.Any().(slog.Level); ok {
					switch level {
					case LevelNotice:
						a.Value = noticeValue
					case LevelVerbose:
						a.Value = verboseValue
					}
				}
				return a
			}
		}
		var handler slog.Handler
		output := os.Stdout

		if cfg.Format == "json" {
			handler = slog.NewJSONHandler(output, opts)
		} else {
			handler = slog.NewTextHandler(output, opts)
		}

		logger := slog.New(handler)
		instance.Store(logger)
		slog.SetDefault(logger)
	})

	return initErr
}

// get returns the singleton logger instance (internal use)
func get() *slog.Logger {
	if logger := instance.Load(); logger != nil {
		return logger
	}

	// Auto-initialize if needed
	Initialize(Config{
		Level:  "info",
		Format: "json",
	})

	return instance.Load()
}

// Direct logging methods - fastest possible access
func Debug(msg string, args ...any) {
	get().Debug(msg, args...)
}

func Verbose(msg string, args ...any) {
	get().Log(context.TODO(), LevelVerbose, msg, args...)
}

func Info(msg string, args ...any) {
	get().Info(msg, args...)
}

func Notice(msg string, args ...any) {
	get().Log(context.TODO(), LevelNotice, msg, args...)
}
func Warn(msg string, args ...any) {
	get().Warn(msg, args...)
}

func Error(msg string, args ...any) {
	get().Error(msg, args...)
}

func Errorf(format string, args ...any) {
	get().Error(fmt.Sprintf(format, args...))
}

// Context-aware methods
func DebugContext(ctx context.Context, msg string, args ...any) {
	get().DebugContext(ctx, msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	get().InfoContext(ctx, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	get().WarnContext(ctx, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	get().ErrorContext(ctx, msg, args...)
}

// With returns a new logger with additional fields
func With(args ...any) *slog.Logger {
	return get().With(args...)
}

// WithGroup returns a new logger with a group
func WithGroup(name string) *slog.Logger {
	return get().WithGroup(name)
}

// parseLevel converts string to slog.Level
func parseLevel(lvl string) (slog.Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	case "notice":
		return LevelNotice, nil
	case "verbose":
		return LevelVerbose, nil
	default:
		return slog.LevelDebug, nil
	}
}
