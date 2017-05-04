// simplelog is a super pared-back version of github.com/goingo/tracelog, modified for Nimblic's very basic requirements:
// 	- Support for a number of logging levels, including a level in between info and warning
// 	- Output logs in format level="LEVEL" msg="LOG_MSG"
package simplelog

import (
	"fmt"
	"log"
	"time"
)

//Printf outputs a formatted log message to the error output
func Println(logger *log.Logger, level string, a ...interface{}) {
	logger.Output(2, fmt.Sprintf("time=%q level=%q msg=%q\n", time.Now().Format(dateTimeFormat), level, fmt.Sprintln(a...)))
}

//Println outputs a formatted log message to the error output
func Printf(logger *log.Logger, level string, format string, a ...interface{}) {
	logger.Output(2, fmt.Sprintf("time=%q level=%q msg=%q\n", time.Now().Format(dateTimeFormat), level, fmt.Sprintf(format, a...)))
}

//Print outputs a log message to the error output
func Print(logger *log.Logger, level string, message string) {
	logger.Output(2, fmt.Sprintf("time=%q level=%q msg=%q\n", time.Now().Format(dateTimeFormat), level, message))
}

//Debugf outputs a formatted log message to the debug output
func Debugf(format string, a ...interface{}) {
	logger.Debug.Output(2, fmt.Sprintf("time=%q level=\"DEBUG\" msg=%q\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Debug outputs a log message to the debug output
func Debug(message string) {
	logger.Debug.Output(2, fmt.Sprintf("time=%q level=\"DEBUG\" msg=%q\n", time.Now().Format(dateTimeFormat), message))
}

//Debugm outputs a set of key-value pairs to the debug output
func Debugm(message string, vals map[string]string) {
	s := fmt.Sprintf("time=%q level=\"DEBUG\" msg=%q\n", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf("%s=%q", k, v)
	}
	logger.Debug.Output(2, s)
}

//Infof outputs a formatted log message to the info output
func Infof(format string, a ...interface{}) {
	logger.Info.Output(2, fmt.Sprintf("time=%q level=\"INFO\" msg=%q\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Info outputs a log message to the info output
func Info(message string) {
	logger.Info.Output(2, fmt.Sprintf("time=%q level=\"INFO\" msg=%q\n", time.Now().Format(dateTimeFormat), message))
}

//Infom outputs a set of key-value pairs to the Info output
func Infom(message string, vals map[string]string) {
	s := fmt.Sprintf("time=%q level=\"INFO\" msg=%q\n", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf("%s=%q", k, v)
	}
	logger.Info.Output(2, s)
}

//Noticef outputs a formatted log message to the notice output
func Noticef(format string, a ...interface{}) {
	logger.Notice.Output(2, fmt.Sprintf("time=%q level=\"NOTICE\" msg=%q\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Notice outputs a message to the notice output
func Notice(message string) {
	logger.Notice.Output(2, fmt.Sprintf("time=%q level=\"NOTICE\" msg=%q\n", time.Now().Format(dateTimeFormat), message))
}

//Noticem outputs a set of key-value pairs to the notice output
func Noticem(message string, vals map[string]string) {
	s := fmt.Sprintf("time=%q level=\"NOTICE\" msg=%q\n", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf("%s=%q", k, v)
	}
	logger.Notice.Output(2, s)
}

//Warningf outputs a formatted log message to the warning output
func Warningf(format string, a ...interface{}) {
	logger.Warning.Output(2, fmt.Sprintf("time=%q level=\"WARNING\" msg=%q\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Warning outputs a message to the warning output
func Warning(message string) {
	logger.Warning.Output(2, fmt.Sprintf("time=%q level=\"WARNING\" msg=%q\n", time.Now().Format(dateTimeFormat), message))
}

func Warnf(format string, a ...interface{}) {
	Warningf(format, a)
}

func Warn(message string) {
	Warning(message)
}

//Warningm outputs a set of key-value pairs to the warning output
func Warningm(message string, vals map[string]string) {
	s := fmt.Sprintf("time=%q level=\"WARNING\" msg=%q\n", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf("%s=%q", k, v)
	}
	logger.Warning.Output(2, s)
}

//Errorf outputs a formatted log message to the error output
func Errorf(format string, a ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("time=%q level=\"ERROR\" msg=%q\n", time.Now().Format(dateTimeFormat), fmt.Sprintf(format, a...)))
}

//Error outputs a message to the error output
func Error(message string) {
	logger.Error.Output(2, fmt.Sprintf("time=%q level=\"ERROR\" msg=%q\n", time.Now().Format(dateTimeFormat), message))
}

//Errorm outputs a set of key-value pairs to the error output
func Errorm(message string, vals map[string]string) {
	s := fmt.Sprintf("time=%q level=\"ERROR\" msg=%q\n", time.Now().Format(dateTimeFormat), message)
	for k, v := range vals {
		s += fmt.Sprintf("%s=%q", k, v)
	}
	logger.Error.Output(2, s)
}
