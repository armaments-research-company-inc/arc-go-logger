package logger

//Warn warn
func Warn(msg ...interface{}) {
	arclogger.Warn(msg...)
}

//Info info
func Info(msg ...interface{}) {
	arclogger.Info(msg...)
}

//Debug debug
func Debug(msg ...interface{}) {
	arclogger.Debug(msg...)
}

//Error error
func Error(msg ...interface{}) {
	arclogger.Error(msg...)
}

//Fatal fatal
func Fatal(msg ...interface{}) {
	arclogger.Fatal(msg...)
}

//Infof infof
func Infof(format string, msg ...interface{}) {
	arclogger.Infof(format, msg...)
}

//Debugf debugf
func Debugf(format string, msg ...interface{}) {
	arclogger.Debugf(format, msg...)
}

//Errorf errorf
func Errorf(format string, msg ...interface{}) {
	arclogger.Errorf(format, msg...)
}

//Fatalf fatalf
func Fatalf(format string, msg ...interface{}) {
	arclogger.Fatalf(format, msg...)
}
