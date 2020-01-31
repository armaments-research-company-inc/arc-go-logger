package logger

//Warn warn
func Warn(msg ...interface{}) {
	arclogger.Warn(msg...)
}

//Info info
func Info(msg ...interface{}) {
	arclogger.Info(msg...)
}

//Infof infof
func Infof(msg ...interface{}) {
	arclogger.Infof(msg...)
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

//Debugf debugf
func Debugf(msg ...interface{}) {
	arclogger.Debugf(msg...)
}

//Errorf errorf
func Errorf(msg ...interface{}) {
	arclogger.Errorf(msg...)
}

//Fatalf fatalf
func Fatalf(msg ...interface{}) {
	arclogger.Fatalf(msg...)
}
