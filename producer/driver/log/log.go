package log

import (
	"context"
	"fmt"
	cx "github/Shitomo/producer/driver/context"
	"time"

	"github/Shitomo/producer/domain/model"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

const format = "2006-01-02T15:04:05.000Z"

func InitLogger() {
	level := zap.InfoLevel

	if model.IsLocalEnv() {
		level = zap.DebugLevel
	}

	encoder := zap.NewDevelopmentEncoderConfig()

	if model.IsLocalEnv() {
		encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoder.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(format))
	})

	zapconf := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoder,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, _ = zapconf.Build()

	defer func() {
		_ = log.Sync()
	}()
}

func WithCtx(ctx context.Context) *zap.Logger {
	id := cx.GetReqCtx(ctx)

	return log.Named(id)
}

func Log() *zap.Logger {
	return log
}

func Access(ctx context.Context, path, method string, start time.Time) {
	elapsed := time.Since(start)
	info := fmt.Sprintf("%s %s %s", path, method, elapsed)
	WithCtx(ctx).Info(info)
}

func Elapsed(ctx context.Context, start time.Time, msg string) {
	elapsed := time.Since(start)
	info := fmt.Sprintf("%s elapsed %s", msg, elapsed)
	WithCtx(ctx).Info(info)
}
