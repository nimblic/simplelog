// simplelog is a super pared-back version of github.com/goingo/tracelog, modified for Nimblic's very basic requirements:
// 	- Support for a number of logging levels, including a level in between info and warning
// 	- Output logs in format \"level\":"LEVEL" msg="LOG_MSG"
package simplelog

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const ErrNotInitialized = "simplelog logger not initialized"

//Printf outputs a formatted log message to the error output
func Println(logger *log.Logger, level string, a ...interface{}) {
	logger.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":%q, \"msg\":%q}\n", time.Now().Format(dateTimeFormat), level, fmt.Sprintln(a...)))
}

//Println outputs a formatted log message to the error output
func Printf(logger *log.Logger, level string, format string, a ...interface{}) {
	logger.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":%q, \"msg\":%q}\n", time.Now().Format(dateTimeFormat), level, fmt.Sprintf(format, a...)))
}

//Print outputs a log message to the error output
func Print(logger *log.Logger, level string, message string) {
	logger.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":%q, \"msg\":%q}\n", time.Now().Format(dateTimeFormat), level, message))
}

//VerboseDebugf outputs a formatted log message to the VerboseDebug output
func VerboseDebugf(format string, a ...interface{}) {
	if logger.Verbose == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Verbose.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"DEBUG\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//VerboseDebug outputs a log message to the VerboseDebug output
func VerboseDebug(message string) {
	if logger.Verbose == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Verbose.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"DEBUG\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), message))
}

//VerboseDebugm outputs a set of key-value pairs to the VerboseDebug output
func VerboseDebugm(message string, vals map[string]string) {
	if logger.Verbose == nil {
		panic(errors.New(ErrNotInitialized))
	}
	s := fmt.Sprintf("{\"time\":%q, \"level\":\"DEBUG\", \"msg\":%q", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf(", \"%s\":%q", k, v)
	}
	s += "}\n"
	logger.Verbose.Output(2, s)
}

//Debugf outputs a formatted log message to the debug output
func Debugf(format string, a ...interface{}) {
	if logger.Debug == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Debug.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"DEBUG\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Debug outputs a log message to the debug output
func Debug(message string) {
	if logger.Debug == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Debug.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"DEBUG\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), message))
}

//Debugm outputs a set of key-value pairs to the debug output
func Debugm(message string, vals map[string]string) {
	if logger.Debug == nil {
		panic(errors.New(ErrNotInitialized))
	}
	s := fmt.Sprintf("{\"time\":%q, \"level\":\"DEBUG\", \"msg\":%q", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf(", \"%s\":%q", k, v)
	}
	s += "}\n"
	logger.Debug.Output(2, s)
}

//Infof outputs a formatted log message to the info output
func Infof(format string, a ...interface{}) {
	if logger.Info == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Info.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"INFO\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Info outputs a log message to the info output
func Info(message string) {
	if logger.Info == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Info.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"INFO\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), message))
}

//Infom outputs a set of key-value pairs to the Info output
func Infom(message string, vals map[string]string) {
	if logger.Info == nil {
		panic(errors.New(ErrNotInitialized))
	}
	s := fmt.Sprintf("{\"time\":%q, \"level\":\"INFO\", \"msg\":%q", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf(", \"%s\":%q", k, v)
	}
	s += "}\n"
	logger.Info.Output(2, s)
}

//Noticef outputs a formatted log message to the notice output
func Noticef(format string, a ...interface{}) {
	if logger.Notice == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Notice.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"NOTICE\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Notice outputs a message to the notice output
func Notice(message string) {
	if logger.Notice == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Notice.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"NOTICE\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), message))
}

//Noticem outputs a set of key-value pairs to the notice output
func Noticem(message string, vals map[string]string) {
	if logger.Notice == nil {
		panic(errors.New(ErrNotInitialized))
	}
	s := fmt.Sprintf("{\"time\":%q, \"level\":\"NOTICE\", \"msg\":%q", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf(", \"%s\":%q", k, v)
	}
	s += "}\n"
	logger.Notice.Output(2, s)
}

//Warningf outputs a formatted log message to the warning output
func Warningf(format string, a ...interface{}) {
	if logger.Warning == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Warning.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"WARNING\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Warning outputs a message to the warning output
func Warning(message string) {
	if logger.Warning == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Warning.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"WARNING\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), message))
}

func Warnf(format string, a ...interface{}) {
	if logger.Warning == nil {
		panic(errors.New(ErrNotInitialized))
	}
	Warningf(format, a)
}

func Warn(message string) {
	if logger.Warning == nil {
		panic(errors.New(ErrNotInitialized))
	}
	Warning(message)
}

//Warningm outputs a set of key-value pairs to the warning output
func Warningm(message string, vals map[string]string) {
	if logger.Warning == nil {
		panic(errors.New(ErrNotInitialized))
	}
	s := fmt.Sprintf("{\"time\":%q, \"level\":\"WARNING\", \"msg\":%q", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf(", \"%s\":%q", k, v)
	}
	s += "}\n"
	logger.Warning.Output(2, s)
}

//Errorf outputs a formatted log message to the error output
func Errorf(format string, a ...interface{}) {
	if logger.Error == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Error.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"ERROR\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Error outputs a message to the error output
func Error(message string) {
	if logger.Error == nil {
		panic(errors.New(ErrNotInitialized))
	}
	logger.Error.Output(2, fmt.Sprintf("{\"time\":%q, \"level\":\"ERROR\", \"msg\":%q}\n", time.Now().Format(dateTimeFormat), message))
}

//Errorm outputs a set of key-value pairs to the error output
func Errorm(message string, vals map[string]string) {
	if logger.Error == nil {
		panic(errors.New(ErrNotInitialized))
	}
	s := fmt.Sprintf("{\"time\":%q, \"level\":\"ERROR\", \"msg\":%q", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf(", \"%s\":%q", k, v)
	}
	s += "}\n"
	logger.Error.Output(2, s)
}
