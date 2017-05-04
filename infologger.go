package simplelog

type infoLogger struct{}

func GetInfoLogger() infoLogger {
	//TODO: allow us to get different loggers (InfoLogger, NoticeLogger etc) to Print to different levels
	return infoLogger{}
}

func (l infoLogger) Print(message string) {
	Print(logger.Info, "INFO", message)
}

func (l infoLogger) Println(a ...interface{}) {
	Println(logger.Info, "INFO", a...)
}

func (l infoLogger) Printf(format string, a ...interface{}) {
	Printf(logger.Info, "INFO", format, a...)
}
