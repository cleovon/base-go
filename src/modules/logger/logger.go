package logger

import (
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey, encoderCfg.EncodeTime = "dateTime", zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.Lock(os.Stdout), zap.DebugLevel)

	fields := []zap.Field{
		{Key: "traceId", Type: zapcore.StringType, String: uuid.NewString()},
	}

	logger := zap.New(core, zap.AddCaller(), zap.Fields(fields...))
	defer logger.Sync()

	return logger
}
