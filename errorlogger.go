package simplelog

type errorLogger struct{}

func GetErrorLogger() errorLogger {
	//TODO: allow us to get different loggers (WarningLogger, NoticeLogger etc) to Print to different levels
	return errorLogger{}
}

func (l errorLogger) Print(message string) {
	Print(logger.Error, "ERROR", message)
}

func (l errorLogger) Println(a ...interface{}) {
	Println(logger.Error, "ERROR", a...)
}

func (l errorLogger) Printf(format string, a ...interface{}) {
	Printf(logger.Error, "ERROR", format, a...)
}
