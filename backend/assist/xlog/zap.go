package xlog

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	Logger, _ = zap.NewProduction()

	defer Logger.Sync() // flushes buffer, if any
}
