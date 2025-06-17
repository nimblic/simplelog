package legacy

type verboseLogger struct{}

func GetVerboseLogger() verboseLogger {
	//TODO: allow us to get different loggers (WarningLogger, NoticeLogger etc) to Print to different levels
	return verboseLogger{}
}

func (l verboseLogger) Print(message string) {
	Print(logger.Debug, "DEBUG", message)
}

func (l verboseLogger) Println(a ...interface{}) {
	Println(logger.Debug, "DEBUG", a...)
}

func (l verboseLogger) Printf(format string, a ...interface{}) {
	Printf(logger.Debug, "DEBUG", format, a...)
}
