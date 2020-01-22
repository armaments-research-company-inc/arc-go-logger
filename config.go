package logger

//Config Config
type Config struct {
	Filename   string
	MaxSize    int //file size in megabytes
	MaxBackups int //max number of files
	MaxAge     int //days
}

//NewConfig for logger
func NewConfig(file string, maxSize int, maxBackups int, maxAge int) *Config {
	return &Config{
		Filename:   file,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxBackups,
	}
}
