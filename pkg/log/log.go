package log

import (
	"context"
	"runtime"
	"strings"
)

var log Logger

// SetLogger .
func SetLogger(logger Logger) {
	log = logger
}

// Debug .
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf .
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// DebugfWithFuncName 将函数名封装在format前
func DebugfWithFuncName(format string, args ...interface{}) {
	callerName := callerFunc()
	log.Debugf(strings.Join([]string{"[%s]", format}, ""), append([]interface{}{callerName}, args...)...)
}

// DebugContextf .
func DebugContextf(ctx context.Context, format string, args ...interface{}) {
	log.DebugContextf(ctx, format, args...)
}

// Info .
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof .
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// callerFunc 获取 log 调用的位置
func callerFunc() string {
	pc, _, _, ok := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if ok && f != nil {
		//return f.Name()
		name := f.Name()
		name = name[strings.LastIndex(name, "/")+1:]
		return name[strings.Index(name, ".")+1:]
	} else {
		return ""
	}
}

// InfofWithFuncName 将函数名封装在format前
func InfofWithFuncName(format string, args ...interface{}) {
	callerName := callerFunc()
	log.Infof(strings.Join([]string{"[%s]", format}, ""), append([]interface{}{callerName}, args...)...)
}

// InfoContextf .
func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	log.InfoContextf(ctx, format, args...)
}

// Warn .
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warnf .
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// WarnContextf .
func WarnContextf(ctx context.Context, format string, args ...interface{}) {
	log.WarnContextf(ctx, format, args...)
}

// Error .
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf .
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// ErrorfWithFuncName 将函数名封装在format前
func ErrorfWithFuncName(format string, args ...interface{}) {
	callerName := callerFunc()
	log.Errorf(strings.Join([]string{"[%s]", format}, ""), append([]interface{}{callerName}, args...)...)
}

// ErrorContextf .
func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	log.ErrorContextf(ctx, format, args...)
}

// Fatal .
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf .
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// FatalContextf .
func FatalContextf(ctx context.Context, format string, args ...interface{}) {
	log.FatalContextf(ctx, format, args...)
}

// Sync ...
func Sync() error {
	return log.Sync()
}
