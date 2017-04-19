package simplelog

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync/atomic"
)

type Level int32

var dateTimeFormat = "2006-01-02 15:04:05.000"

const (
	// LevelDebug logs everything
	LevelDebug Level = 1

	// LevelInfo logs Info, Notices, Warnings and Errors
	LevelInfo Level = 2

	// LevelNotice logs Notices, Warnings and Errors
	LevelNotice Level = 4

	// LevelWarn logs Warning and Errors
	LevelWarn Level = 8

	// LevelError logs just Errors
	LevelError Level = 16
)

type simpleLog struct {
	LogLevel int32
	Debug    *log.Logger
	Info     *log.Logger
	Notice   *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
}

//SetTimestampFormat sets for format for log entry timestamps
func SetTimestampFormat(f string) {
	dateTimeFormat = f
}

// ParseLevel takes a string level and returns the simplelog level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "error", "err":
		return LevelError, nil
	case "warn", "warning":
		return LevelWarn, nil
	case "notice":
		return LevelNotice, nil
	case "info":
		return LevelInfo, nil
	case "debug":
		return LevelDebug, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid Level: %q", lvl)
}

// log maintains a pointer to a singleton for the logging system.
var logger simpleLog

// Called to init the logging system.
func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

// Start initializes simpleLog and only displays the specified logging level.
func Start(logLevel Level) {
	turnOnLogging(int32(logLevel))
}

// LogLevel returns the configured logging level.
func LogLevel() int32 {
	return atomic.LoadInt32(&logger.LogLevel)
}

// turnOnLogging configures the logging writers.
func turnOnLogging(logLevel int32) {
	debugHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	noticeHandle := ioutil.Discard
	warnHandle := ioutil.Discard
	errorHandle := ioutil.Discard

	if logLevel&int32(LevelDebug) != 0 {
		debugHandle = os.Stdout
		infoHandle = os.Stdout
		noticeHandle = os.Stdout
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&int32(LevelInfo) != 0 {
		infoHandle = os.Stdout
		warnHandle = os.Stdout
		noticeHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&int32(LevelNotice) != 0 {
		warnHandle = os.Stdout
		noticeHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&int32(LevelWarn) != 0 {
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&int32(LevelError) != 0 {
		errorHandle = os.Stderr
	}

	logger.Debug = log.New(debugHandle, "", 0)
	logger.Info = log.New(infoHandle, "", 0)
	logger.Notice = log.New(noticeHandle, "", 0)
	logger.Warning = log.New(warnHandle, "", 0)
	logger.Error = log.New(errorHandle, "", 0)

	atomic.StoreInt32(&logger.LogLevel, logLevel)
}
