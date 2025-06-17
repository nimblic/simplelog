package legacy

type warningLogger struct{}

func GetWarningLogger() warningLogger {
	//TODO: allow us to get different loggers (WarningLogger, NoticeLogger etc) to Print to different levels
	return warningLogger{}
}

func (l warningLogger) Print(message string) {
	Print(logger.Info, "WARNING", message)
}

func (l warningLogger) Println(a ...interface{}) {
	Println(logger.Info, "WARNING", a...)
}

func (l warningLogger) Printf(format string, a ...interface{}) {
	Printf(logger.Info, "WARNING", format, a...)
}
