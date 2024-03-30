package config

var (
	Secret_key = "secret-key-token"
)

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
