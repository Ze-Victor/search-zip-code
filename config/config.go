package config

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
