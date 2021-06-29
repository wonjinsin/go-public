package utils

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func (m *Logger) Logging() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger.Sugar()
}
