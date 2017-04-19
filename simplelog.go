package simplelog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync/atomic"
)

type Level int32

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
	log.SetPrefix("DEBUG")
	log.SetFlags(log.Ldate | log.Ltime)
}

// Start initializes simpleLog and only displays the specified logging level.
func Start(logLevel Level) {
	turnOnLogging(int32(logLevel), nil)
}

// LogLevel returns the configured logging level.
func LogLevel() int32 {
	return atomic.LoadInt32(&logger.LogLevel)
}

// turnOnLogging configures the logging writers.
func turnOnLogging(logLevel int32, fileHandle io.Writer) {
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

	logger.Debug = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime)
	logger.Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime)
	logger.Notice = log.New(noticeHandle, "NOTICE: ", log.Ldate|log.Ltime)
	logger.Warning = log.New(warnHandle, "WARNING: ", log.Ldate|log.Ltime)
	logger.Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime)

	atomic.StoreInt32(&logger.LogLevel, logLevel)
}
