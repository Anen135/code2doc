package logger

func Info(msg string) {
	Log(INFO, msg)
}

func Error(msg string) {
	Log(ERROR, msg)
}

func ErrorV(err error) {
	Log(ERROR, err.Error())
}

func Warning(msg string) {
	Log(WARNING, msg)
}

func Debug(msg string) {
	Log(DEBUG, msg)
}
