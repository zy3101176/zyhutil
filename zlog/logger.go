package zlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

// log level
const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

var sugarLogger *zap.SugaredLogger

func InitLogger(path string, level string) {
	writeSyncer := getLogWriter(path)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, convertLevel(level))
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	sugarLogger.Sync()
}

func convertLevel(level string) Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	default:
		return InfoLevel
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args)
}

func Panic(args ...interface{}) {
	sugarLogger.Panic(args)
}

func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args)
}

func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args)
}
