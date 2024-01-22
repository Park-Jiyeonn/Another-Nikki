package log

import (
	"context"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func Init() {
	// TODO: 此处的逻辑太丑了，毕竟我不是真的几台机器上的微服务，都是本地直接跑的，没办法，会重复创建 Client
	//if zap.L() != nil {
	//	return
	//}
	client, err := elastic.NewTypedClient(elastic.Config{
		Addresses: []string{"http://47.116.20.160:9200"},
	})
	if err != nil {
		panic(err)
	}
	elasticHook, err := NewElasticsearchHook(client, "test-1")
	NewWithHook(elasticHook)
}

func NewWithHook(w io.Writer) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	// 配置Zap core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(zapcore.AddSync(w))),
		zap.DebugLevel, // 设置日志级别
	)

	// 创建Zap logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	defer logger.Sync()

	zap.ReplaceGlobals(logger)
}

func Info(ctx context.Context, msg string, field ...zap.Field) {
	if traceId, ok := ExtractTraceIDFromContext(ctx); ok {
		field = append(field, zap.String("trace_id", traceId))
	}
	zap.L().Info(msg, field...)
}

func Warn(ctx context.Context, msg string, field ...zap.Field) {
	if traceId, ok := ExtractTraceIDFromContext(ctx); ok {
		field = append(field, zap.String("trace_id", traceId))
	}
	zap.L().Warn(msg, field...)
}

func Error(ctx context.Context, msg string, field ...zap.Field) {
	if traceId, ok := ExtractTraceIDFromContext(ctx); ok {
		field = append(field, zap.String("trace_id", traceId))
	}
	zap.L().Error(msg, field...)
}
