package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	*zap.SugaredLogger
}

var (
	DefaultLogger Logger
)

func init() {
	DefaultLogger, err := NewZapLogger()
	if err != nil {
		panic(err)
	}
}

func NewZapLogger() (*ZapLogger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 小写编码器
		EncodeTime:     TimeEncoder,                      // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}

	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.DebugLevel)
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	config := zap.Config{
		Level:         atom,
		Development:   true,
		Encoding:      "console",
		DisableCaller: true,
		EncoderConfig: encoderConfig,
		Sampling: &zap.SamplingConfig{
			Initial:    1,
			Thereafter: 100,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return &ZapLogger{logger.Sugar()}, nil
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05"))
}

func (logger *ZapLogger) Debugln(args ...interface{}) {
	logger.Debug(args...)
}

func (logger *ZapLogger) Infoln(args ...interface{}) {
	logger.Info(args...)
}

func (logger *ZapLogger) Warnln(args ...interface{}) {
	logger.Warn(args...)
}

func (logger *ZapLogger) Errorln(args ...interface{}) {
	logger.Error(args...)
}
