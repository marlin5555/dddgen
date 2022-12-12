package log

import "context"

// Logger 日志接口
type Logger interface {
	// Debug logs to DEBUG log. Arguments are handled in the manner of fmt.Print.
	Debug(args ...interface{})
	// Debugf logs to DEBUG log. Arguments are handled in the manner of fmt.Print.
	Debugf(format string, args ...interface{})
	// DebugContextf logs to DEBUG log. Arguments are handled in the manner of fmt.Printf.
	DebugContextf(ctx context.Context, format string, args ...interface{})
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Infof(format string, args ...interface{})
	// InfoContextf logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	InfoContextf(ctx context.Context, format string, args ...interface{})
	// Warn logs to WARN log. Arguments are handled in the manner of fmt.Print.
	Warn(args ...interface{})
	// Warnf logs to WARN log. Arguments are handled in the manner of fmt.Printf.
	Warnf(format string, args ...interface{})
	// WarnContextf logs to WARN log. Arguments are handled in the manner of fmt.Printf.
	WarnContextf(ctx context.Context, format string, args ...interface{})
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
	// ErrorContextf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	ErrorContextf(ctx context.Context, format string, args ...interface{})
	// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...interface{})
	// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalf(format string, args ...interface{})
	// FatalContextf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	FatalContextf(ctx context.Context, format string, args ...interface{})
	// Sync resource close
	Sync() error
}
