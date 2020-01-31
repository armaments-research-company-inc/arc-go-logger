package logger

import (
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

//Logger tem

var arclogger *logrus.Logger

var initOnce sync.Once

//Initialize logger
func Initialize(conf *Config) {

	initOnce.Do(func() {
		arclogger = logrus.New()
		arclogger.SetFormatter(&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%]: %msg% \n",
		})

		if conf == nil {
			arclogger.SetOutput(os.Stdout)
			return
		}

		arclogger.SetOutput(&lumberjack.Logger{
			Filename:   conf.Filename,
			MaxSize:    conf.MaxSize,    //file size in megabytes
			MaxBackups: conf.MaxBackups, //max number of files
			MaxAge:     conf.MaxAge,     //days
		})

	})
}
