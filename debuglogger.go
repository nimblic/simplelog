package simplelog

type debugLogger struct{}

func GetDebugLogger() debugLogger {
	//TODO: allow us to get different loggers (WarningLogger, NoticeLogger etc) to Print to different levels
	return debugLogger{}
}

func (l debugLogger) Print(message string) {
	Print(logger.Debug, "DEBUG", message)
}

func (l debugLogger) Println(a ...interface{}) {
	Println(logger.Debug, "DEBUG", a...)
}

func (l debugLogger) Printf(format string, a ...interface{}) {
	Printf(logger.Debug, "DEBUG", format, a...)
}
