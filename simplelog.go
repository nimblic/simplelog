package simplelog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync/atomic"
)

type Level int32

var (
	dateTimeFormat = "2006-01-02 15:04:05.000"
	persistLog     = false
	messageLog     []string
)

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

//ArrayWriter appends writes to an array of strings
type ArrayWriter struct{}

func (a ArrayWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	messageLog = append(messageLog, s)
	fmt.Print(s)
	return len(p), nil
}

var arrayWriter = ArrayWriter{}

//SetTimestampFormat sets for format for log entry timestamps
func SetTimestampFormat(f string) {
	dateTimeFormat = f
}

//PersistLog will store messages in an array for testing purposes
func PersistLog(p bool) {
	if p {
		logger.Debug.SetOutput(arrayWriter)
		logger.Info.SetOutput(arrayWriter)
		logger.Notice.SetOutput(arrayWriter)
		logger.Warning.SetOutput(arrayWriter)
		logger.Error.SetOutput(arrayWriter)
	} else {
		setLoggingLevel(logger.LogLevel)
	}
}

//return the array of all logged messages
func GetMessages() []string {
	return messageLog
}

type logEntry struct {
	Time  string `json:"time"`
	Msg   string `json:"msg"`
	Level string `json:"level"`
}

//return true if the a message has been logged
func LogContainsMessage(message string) bool {
	return LogContains(message, "")
}

//return true if the a message with a given level has been logged
func LogContains(message string, level string) bool {
	for _, m := range messageLog {
		var e logEntry
		err := json.Unmarshal([]byte(m), &e)
		if err != nil {
			panic(fmt.Errorf("Invalid log entry: %s", m))
		}
		if e.Msg == message {
			if level == "" || e.Level == level {
				return true
			}
		}
	}
	return false
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

// Used to retrieve the current logger for sending to external libs (look at fhir_client.go in hl7processor)
func Logger() *simpleLog {
	return &logger
}

// Init initializes simplelog to only display logs at or above the specified logging level, and returns a pointer to the logger
func Init(logLevel Level) *simpleLog {
	setLoggingLevel(int32(logLevel))
	return &logger
}

// LogLevel returns the configured logging level.
func LogLevel() int32 {
	return atomic.LoadInt32(&logger.LogLevel)
}

// turnOnLogging configures the logging writers.
func setLoggingLevel(logLevel int32) {
	debugHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	noticeHandle := ioutil.Discard
	warnHandle := ioutil.Discard
	errorHandle := os.Stderr

	if logLevel&int32(LevelDebug) != 0 {
		debugHandle = os.Stdout
		infoHandle = os.Stdout
		noticeHandle = os.Stdout
		warnHandle = os.Stdout
	}

	if logLevel&int32(LevelInfo) != 0 {
		infoHandle = os.Stdout
		warnHandle = os.Stdout
		noticeHandle = os.Stdout
	}

	if logLevel&int32(LevelNotice) != 0 {
		warnHandle = os.Stdout
		noticeHandle = os.Stdout
	}

	if logLevel&int32(LevelWarn) != 0 {
		warnHandle = os.Stdout
	}

	logger.Debug = log.New(debugHandle, "", 0)
	logger.Info = log.New(infoHandle, "", 0)
	logger.Notice = log.New(noticeHandle, "", 0)
	logger.Warning = log.New(warnHandle, "", 0)
	logger.Error = log.New(errorHandle, "", 0)

	atomic.StoreInt32(&logger.LogLevel, logLevel)
}
