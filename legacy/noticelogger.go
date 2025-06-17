package legacy

type noticeLogger struct{}

func GetNoticeLogger() noticeLogger {
	//TODO: allow us to get different loggers (NoticeLogger, NoticeLogger etc) to Print to different levels
	return noticeLogger{}
}

func (l noticeLogger) Print(message string) {
	Print(logger.Info, "NOTICE", message)
}

func (l noticeLogger) Println(a ...interface{}) {
	Println(logger.Info, "NOTICE", a...)
}

func (l noticeLogger) Printf(format string, a ...interface{}) {
	Printf(logger.Info, "NOTICE", format, a...)
}
