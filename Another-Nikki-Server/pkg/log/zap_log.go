package log

import (
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
)

func Init(env, serviceName string) log.Logger {

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder.MessageKey = "message"

	zapLogger := NewZapLogger(
		encoder,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCallerSkip(1),
		zap.Development(),
	)

	logger := log.With(zapLogger,
		//"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", env,
		"service.name", serviceName,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	//log.SetLogger(logger)

	return logger
}

type Logger struct {
	log  *zap.Logger
	Sync func() error
}

func NewZapLogger(encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *Logger {
	client, err := elastic.NewTypedClient(elastic.Config{
		Addresses: []string{"http://47.116.20.160:9200"},
	})
	if err != nil {
		panic(err)
	}
	elasticHook, err := NewElasticsearchHook(client, "test-1")
	// 设置 zapcore
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(elasticHook),
		), level)
	//  new 一个 *zap.Logger
	zapLogger := zap.New(core, opts...)
	return &Logger{log: zapLogger, Sync: zapLogger.Sync}
}

// Log 方法实现了 kratos/log/log.go 中的 Logger interface
func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// 按照 KV 传入的时候,使用的 zap.Field
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}

func Info(ctx context.Context, msg string, filed ...interface{}) {
	log.Context(ctx).Infof(msg, filed...)
}

func Warn(ctx context.Context, msg string, filed ...interface{}) {
	log.Context(ctx).Warnf(msg, filed...)
}

func Error(ctx context.Context, msg string, filed ...interface{}) {
	log.Context(ctx).Errorf(msg, filed...)
}
