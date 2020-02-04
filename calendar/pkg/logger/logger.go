package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Configuration структура для настройки логирования
type Configuration struct {
	Level      string
	JSONFormat bool
}

func getZapLevel(level string) zap.AtomicLevel {
	switch level {
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02/01/2006 3:4:5 PM"))
}

//NewLogger returns an instance of logger
func NewLogger(config Configuration) (*zap.Logger, error) {

	cfg := zap.Config{
		Encoding:         "console", //"json",
		Level:            getZapLevel(config.Level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey: "time",
			//EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeTime: syslogTimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	return logger, err
}

//InitLogger Вариант инициализации логера
func InitLogger() *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(config)
	//writerSyncer := getLogWriter()
	writerSyncer := zapcore.Lock(os.Stdout)
	atom := zap.NewAtomicLevel()
	logr := zap.New(zapcore.NewCore(encoder, writerSyncer, atom))

	return logr
}
