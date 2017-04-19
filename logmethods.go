// simplelog is a super pared-back version of github.com/goingo/tracelog, modified for Nimblic's very basic requirements:
// 	- Support for a number of logging levels, including a level in between info and warning
// 	- Output logs in format level="LEVEL" message="LOG_MSG"
package simplelog

import (
	"fmt"
)

//Debugf output a formatted log message to the debug output
func Debugf(format string, a ...interface{}) {
	logger.Debug.Output(2, fmt.Sprintf("level=\"DEBUG\" message=\"%s\"\n", fmt.Sprintf(format, a...)))
}

//Debug outputs a log message to the debug output
func Debug(message string) {
	logger.Debug.Output(2, fmt.Sprintf("level=\"DEBUG\" message=\"%s\"\n", message))
}

//Infof output a formatted log message to the info output
func Infof(format string, a ...interface{}) {
	logger.Info.Output(2, fmt.Sprintf("level=\"INFO\" message=\"%s\"\n", fmt.Sprintf(format, a...)))
}

//Info outputs a log message to the info output
func Info(message string) {
	logger.Info.Output(2, fmt.Sprintf("level=\"INFO\" message=\"%s\"\n", message))
}

//Noticef output a formatted log message to the notice output
func Noticef(format string, a ...interface{}) {
	logger.Notice.Output(2, fmt.Sprintf("level=\"NOTICE\" message=\"%s\"\n", fmt.Sprintf(format, a...)))
}

//Notice outputs a message to the notice output
func Notice(message string) {
	logger.Notice.Output(2, fmt.Sprintf("level=\"NOTICE\" message=\"%s\"\n", message))
}

//Warningf output a formatted log message to the warning output
func Warningf(format string, a ...interface{}) {
	logger.Warning.Output(2, fmt.Sprintf("level=\"WARNING\" message=\"%s\"\n", fmt.Sprintf(format, a...)))
}

//Warning outputs a message to the warning output
func Warning(message string) {
	logger.Warning.Output(2, fmt.Sprintf("level=\"WARNING\" message=\"%s\"\n", message))
}

//Errorf output a formatted log message to the error output
func Errorf(format string, a ...interface{}) {
	logger.Error.Output(2, fmt.Sprintf("level=\"ERROR\" message=\"%s\"\n", fmt.Sprintf(format, a...)))
}

//Error outputs a message to the error output
func Error(message string) {
	logger.Error.Output(2, fmt.Sprintf("level=\"ERROR\" message=\"%s\"\n", message))
}
