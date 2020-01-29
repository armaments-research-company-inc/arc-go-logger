package logger

//Warn warn
func Warn(msg ...interface{}) {
	logger.Warn(msg)
}

//Info info
func Info(msg ...interface{}) {
	logger.Info(msg)
}

//Debug debug
func Debug(msg ...interface{}) {
	logger.Debug(msg)
}

//Error error
func Error(msg ...interface{}) {
	logger.Error(msg)
}

//Fatal fatal
func Fatal(msg ...interface{}) {
	logger.Fatal(msg)
}
