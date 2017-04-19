package simplelog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync/atomic"
)

const (
	// LevelDebug logs everything
	LevelDebug int32 = 1

	// LevelInfo logs Info, Notices, Warnings and Errors
	LevelInfo int32 = 2

	// LevelNotice logs Notices, Warnings and Errors
	LevelNotice int32 = 4

	// LevelWarn logs Warning and Errors
	LevelWarn int32 = 8

	// LevelError logs just Errors
	LevelError int32 = 16
)

type simpleLog struct {
	LogLevel int32
	Debug    *log.Logger
	Info     *log.Logger
	Notice   *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
}

// log maintains a pointer to a singleton for the logging system.
var logger simpleLog

// Called to init the logging system.
func init() {
	log.SetPrefix("DEBUG")
	log.SetFlags(log.Ldate | log.Ltime)
}

// Start initializes simpleLog and only displays the specified logging level.
func Start(logLevel int32) {
	turnOnLogging(logLevel, nil)
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

	if logLevel&LevelDebug != 0 {
		debugHandle = os.Stdout
		infoHandle = os.Stdout
		noticeHandle = os.Stdout
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&LevelInfo != 0 {
		infoHandle = os.Stdout
		warnHandle = os.Stdout
		noticeHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&LevelNotice != 0 {
		warnHandle = os.Stdout
		noticeHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&LevelWarn != 0 {
		warnHandle = os.Stdout
		errorHandle = os.Stderr
	}

	if logLevel&LevelError != 0 {
		errorHandle = os.Stderr
	}

	logger.Debug = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime)
	logger.Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime)
	logger.Notice = log.New(noticeHandle, "NOTICE: ", log.Ldate|log.Ltime)
	logger.Warning = log.New(warnHandle, "WARNING: ", log.Ldate|log.Ltime)
	logger.Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime)

	atomic.StoreInt32(&logger.LogLevel, logLevel)
}
