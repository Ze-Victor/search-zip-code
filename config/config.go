package config

var (
	Secret_key = "secret-key-token"
	Base_Path  = "http://localhost:8001/api/v1"
)

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
