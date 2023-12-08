package pkg

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LogLevelStringToZapLevel = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

var (
	// Logger is the (logger) of zap
	logger *zap.Logger
	// Level is the logger level of zap
	level zap.AtomicLevel
)

// logLayout is the layout of log time
const logLayout = "2006-01-02 15:04:05.000"

func InitLogger(logLevel string) {
	var err error

	l, ok := LogLevelStringToZapLevel[logLevel]
	if !ok {
		log.Println("invalid log level. set to info")
		l = zapcore.InfoLevel
	}

	level = zap.NewAtomicLevel()
	level.SetLevel(l)
	logger, err = zap.Config{
		Level:             level,
		Development:       false,
		Encoding:          "json",
		DisableStacktrace: true,
		DisableCaller:     true,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			EncodeTime:     zapcore.TimeEncoderOfLayout(logLayout),
			EncodeDuration: zapcore.StringDurationEncoder,

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			NameKey:     "key",
			FunctionKey: zapcore.OmitKey,

			MessageKey: "msg",
			LineEnding: zapcore.DefaultLineEnding,
		},
	}.Build()

	if err != nil {
		panic(err)
	}
}
