package simplelog

type errorLogger struct{}

func GetErrorLogger() errorLogger {
	//TODO: allow us to get different loggers (WarningLogger, NoticeLogger etc) to Print to different levels
	return errorLogger{}
}

func (l errorLogger) Print(message string) {
	Errorf("ERROR: %s", message)
}

func (l errorLogger) Println(a ...interface{}) {
	Errorf("ERROR: %s", a...)
}

func (l errorLogger) Printf(format string, a ...interface{}) {
	Errorf("ERROR: %s", a...)
}
