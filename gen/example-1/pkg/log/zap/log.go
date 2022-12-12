package zap

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLog 初始化 log
func InitLog(path string, level string) *Logger {

	encoder := getEncoder()
	writer := getLogWriter(path)

	var lev zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		lev = zap.DebugLevel
	case "info":
		lev = zap.InfoLevel
	case "error":
		lev = zap.ErrorLevel
	default:
		lev = zap.InfoLevel
	}

	atomicLevel := zap.NewAtomicLevelAt(lev)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer),
		atomicLevel)

	logger := zap.New(core,
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.DPanicLevel)),
		zap.AddCaller(), // 开启开发模式，堆栈跟踪
		zap.AddCallerSkip(2),
		zap.Development(), // 开启文件及行号
	)

	return &Logger{logger: logger}
}

func getEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",   // 时间的 key
		LevelKey:       "level",  // 日志等级的 key
		NameKey:        "logger", // 日志名称的 key
		CallerKey:      "caller", // 调用函数的 key
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder}
}

func getLogWriter(path string) zapcore.WriteSyncer {
	if path == "" {
		path = "./grpc.log" // 默认文件名，当前路径
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,  // 保存的路径
		MaxSize:    50,    // 保存大小，单位MB
		MaxBackups: 5,     // 日志的最大保存数量
		MaxAge:     30,    // 日志文件存储最大天数
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Logger grpc 日志实现
type Logger struct {
	logger *zap.Logger
}

// Debug .
func (g *Logger) Debug(args ...interface{}) {
	g.logger.Sugar().Debug(args...)
}

// Debugf .
func (g *Logger) Debugf(format string, args ...interface{}) {
	g.logger.Sugar().Debugf(format, args...)
}

// DebugContextf .
func (g *Logger) DebugContextf(ctx context.Context, format string, args ...interface{}) {
	g.logger.Sugar().Debugf(format, args...)
}

// Info .
func (g *Logger) Info(args ...interface{}) {
	g.logger.Sugar().Info(args...)
}

// Infof .
func (g *Logger) Infof(format string, args ...interface{}) {
	g.logger.Sugar().Infof(format, args...)
}

// InfoContextf .
func (g *Logger) InfoContextf(ctx context.Context, format string, args ...interface{}) {
	g.logger.Sugar().Infof(format, args...)
}

// Warn .
func (g *Logger) Warn(args ...interface{}) {
	g.logger.Sugar().Warn(args...)
}

// Warnf .
func (g *Logger) Warnf(format string, args ...interface{}) {
	g.logger.Sugar().Warnf(format, args...)
}

// WarnContextf .
func (g *Logger) WarnContextf(ctx context.Context, format string, args ...interface{}) {
	g.logger.Sugar().Warnf(format, args...)
}

// Error .
func (g *Logger) Error(args ...interface{}) {
	g.logger.Sugar().Error(args...)
}

// Errorf .
func (g *Logger) Errorf(format string, args ...interface{}) {
	g.logger.Sugar().Errorf(format, args...)
}

// ErrorContextf .
func (g *Logger) ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	g.logger.Sugar().Errorf(format, args)
}

// Fatal .
func (g *Logger) Fatal(args ...interface{}) {
	g.logger.Sugar().Fatal(args...)
}

// Fatalf .
func (g *Logger) Fatalf(format string, args ...interface{}) {
	g.logger.Sugar().Fatalf(format, args...)
}

// FatalContextf .
func (g *Logger) FatalContextf(ctx context.Context, format string, args ...interface{}) {
	g.logger.Sugar().Fatalf(format, args...)
}

// Sync 资源关闭
func (g *Logger) Sync() error {
	return g.logger.Sugar().Sync()
}
