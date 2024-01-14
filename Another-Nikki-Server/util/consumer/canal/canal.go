package canal

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"reflect"
	"strings"
	"sync"
)

const (
	FullMatch   MatchMethod  = 1
	PrefixMatch MatchMethod  = 2
	Insert      BinlogAction = "insert"
	Update      BinlogAction = "update"
	Delete      BinlogAction = "delete"
	LogTemplate              = "[CanalBinlogConsumer][railgunId:%s][%s][Error], err:%+v"
)

type CanalConsumer struct {
	reader   *kafka.Reader
	config   *CanalConsumerConfig
	handlers map[string]*CanalConsumerHandler
}
type CanalConsumerConfig struct {
	Id     string // Group Id
	MsgLog bool
}
type CanalConsumerHandler struct {
	// 对该类表的匹配方式
	MatchMethod MatchMethod
	// 当有错误返回时，是否进行重试
	ErrRetry       bool
	database       string
	tableName      string
	unpackFunc     TableMsgUnpackFunc
	insertHandlers []*handlerInfo
	updateHandlers []*handlerInfo
	deleteHandlers []*handlerInfo
	allHandlers    []*handlerInfo
}
type handlerInfo struct {
	database string
	retry    bool
	doFunc   BinlogConsumerHandlerFunc
}

// MatchMethod 表名匹配类型: 全匹配、前缀匹配
type MatchMethod int8

// BinlogAction 表的action类型：insert/update/delete
type BinlogAction string

// BinlogConsumerHandlerFunc 统一binlog消费的函数定义格式
type BinlogConsumerHandlerFunc func(ctx context.Context, action string, new interface{}, old interface{}) error

// TableMsgUnpackFunc 单表数据解析的函数方法 jsonData为解析前数据，interface{}为解析后数据的结构地址
type TableMsgUnpackFunc func(ctx context.Context, jsonData []byte) (interface{}, error)

func New(config *CanalConsumerConfig) *CanalConsumer {
	if config == nil || config.Id == "" {
		panic("CanalConsumer Init error, Config nil")
	}
	kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "hello",
		MaxBytes: 1e6, // 10MB
	})
	consumer := new(CanalConsumer)
	consumer.config = config
	consumer.handlers = make(map[string]*CanalConsumerHandler)
	return consumer
}

func (canalConsumer *CanalConsumer) Group(table string, tableMatch MatchMethod, unpackFunc TableMsgUnpackFunc) *CanalConsumerHandler {
	if table == "" {
		panic("table is empty")
	}
	tableName := table
	if canalConsumer.handlers[tableName] != nil {
		group := canalConsumer.handlers[tableName]
		if group.MatchMethod != tableMatch || group.tableName != table || reflect.TypeOf(group.unpackFunc) != reflect.TypeOf(unpackFunc) {
			panic("multi init params not match")
		}
		return group
	}
	canalConsumer.handlers[tableName] = &CanalConsumerHandler{
		MatchMethod: tableMatch,
		tableName:   table,
		unpackFunc:  unpackFunc,
	}
	return canalConsumer.handlers[tableName]
}

func (handler *handlerInfo) DB(database string) *handlerInfo {
	handler.database = database
	return handler
}

func (handler *handlerInfo) Retry(retry bool) *handlerInfo {
	handler.retry = retry
	return handler
}

func defaultHandlerInfo(f BinlogConsumerHandlerFunc) *handlerInfo {
	return &handlerInfo{
		database: "",
		retry:    true,
		doFunc:   f,
	}
}

func (handler *CanalConsumerHandler) ALL(f BinlogConsumerHandlerFunc) *handlerInfo {
	newHandler := defaultHandlerInfo(f)
	handler.allHandlers = append(handler.allHandlers, newHandler)
	return newHandler
}

func (handler *CanalConsumerHandler) INSERT(f BinlogConsumerHandlerFunc) *handlerInfo {
	newHandler := defaultHandlerInfo(f)
	handler.insertHandlers = append(handler.insertHandlers, newHandler)
	return newHandler
}

func (handler *CanalConsumerHandler) UPDATE(f BinlogConsumerHandlerFunc) *handlerInfo {
	newHandler := defaultHandlerInfo(f)
	handler.updateHandlers = append(handler.updateHandlers, newHandler)
	return newHandler
}

func (handler *CanalConsumerHandler) DELETE(f BinlogConsumerHandlerFunc) *handlerInfo {
	newHandler := defaultHandlerInfo(f)
	handler.deleteHandlers = append(handler.deleteHandlers, newHandler)
	return newHandler
}

func (canalConsumer *CanalConsumer) Start() {
	if canalConsumer.config == nil || canalConsumer.config.Id == "" {
		panic("canal consumer start err because railgun Id nil")
	}
	//processor := single.New(actMsgUnpack, canalConsumer.actMsgHandler)
	var err error
	// 指定的处理器 ID
	canalConsumer.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  canalConsumer.config.Id,
		Topic:    "hello",
		MaxBytes: 1e6, // 10MB
	})
	if err != nil {
		panic(fmt.Sprintf("canal consumer start err: %v", err))
	}
}

type binlogMsg struct {
	Database string          `json:"schema"`
	Action   string          `json:"action"`
	Table    string          `json:"table"`
	New      json.RawMessage `json:"new"`
	Old      json.RawMessage `json:"old"`
}

//func actMsgUnpack(msg message.Message) (*single.UnpackMessage, error) {
//	msgUnpack := new(binlogMsg)
//	if err := json.Unmarshal(msg.Payload(), &msgUnpack); err != nil {
//		return nil, err
//	}
//	return &single.UnpackMessage{
//		Group: time.Now().Unix(),
//		Item:  msgUnpack,
//	}, nil
//}
//
//func (canalConsumer *CanalConsumer) actMsgHandler(ctx context.Context, msg interface{}, _ *single.Extra) message.Policy {
//	notifyMsg, ok := msg.(*binlogMsg)
//	if !ok || notifyMsg.Table == "" {
//		return message.Ignore
//	}
//	//group := canalConsumer.config.RailgunConfig.Databus.Group
//
//	if canalConsumer.config.MsgLog {
//		jsonV, err := json.Marshal(notifyMsg)
//		if err != nil {
//			log.Errorc(ctx, LogTemplate, canalConsumer.config.Id, "Marsh", err)
//		}
//		log.Infoc(ctx, "[CanalBinlogConsumer] group:%s, msg:%+v", canalConsumer.config.Id, string(jsonV))
//	}
//	tableForConfig, err := canalConsumer.tableMatchCheck(notifyMsg.Table)
//	if err != nil {
//		// 此处不记录日志，因为可能存在表不被消费
//		// 此处的规则后续再探讨
//		return message.Ignore
//	}
//	handlersConfig := canalConsumer.handlers[tableForConfig]
//	if handlersConfig == nil {
//		return message.Ignore
//	}
//	handlers := handlersFind(handlersConfig, notifyMsg.Action, notifyMsg.Database)
//	_ = canalConsumer.handlersDo(ctx, handlers, handlersConfig.unpackFunc, notifyMsg.Action, notifyMsg.New, notifyMsg.Old)
//	return message.Success
//}

func (canalConsumer *CanalConsumer) handlersDo(ctx context.Context, handlers []*handlerInfo, unpackFunc TableMsgUnpackFunc, action string, newData []byte, old []byte) (err error) {
	if len(handlers) == 0 {
		return
	}
	//group := canalConsumer.config.RailgunConfig.Databus.Group

	var newInterface, oldInterface interface{}
	newInterface, err = unpackFunc(ctx, newData)
	if err != nil {
		//log.Errorc(ctx, LogTemplate, canalConsumer.config.Id, "handlerDo new, unpackFunc doError", err)
		return
	}
	if len(old) > 0 {
		oldInterface, err = unpackFunc(ctx, old)
		if err != nil {
			//log.Errorc(ctx, LogTemplate, canalConsumer.config.Id, "handlerDo old, unpackFunc doError", err)
			return
		}
	}
	//defaultAttempts := 2
	//if canalConsumer.config != nil && canalConsumer.config.RailgunConfig != nil && canalConsumer.config.RailgunConfig.Railgun != nil &&
	//	canalConsumer.config.RailgunConfig.Railgun.Attempts != 0 {
	//	defaultAttempts = canalConsumer.config.RailgunConfig.Railgun.Attempts
	//}
	funcs := make([]BinlogConsumerHandlerFunc, 0)
	for _, handler := range handlers {
		funcs = append(funcs, handler.doFunc)
	}
	// TODO 该重试逻辑 可删除，因为railgun 后台可配置， 如果该动作重试了，会导致raingun平台的重试无效
	//backOff := &_defaultBackoffConfig
	//if canalConsumer.config != nil && canalConsumer.config.RailgunConfig != nil && canalConsumer.config.RailgunConfig.Railgun != nil &&
	//	canalConsumer.config.RailgunConfig.Railgun.Backoff != nil {
	//	backOff = canalConsumer.config.RailgunConfig.Railgun.Backoff
	//}
	//handlerMapping := formatHandlerMapping(handlers)
	//err = retry.WithAttempts(ctx, fmt.Sprintf("TOOL_Canal_railgunId:%s", canalConsumer.config.Id), defaultAttempts, *backOff, func(c context.Context) (err error) {
	//	if len(funcs) == 0 {
	//		return
	//	}
	//	_, failFunc, err := oneTimeDo(ctx, funcs, action, newInterface, oldInterface)
	//	if err != nil {
	//		funcs = formatNewDoFunctions(handlerMapping, failFunc)
	//	}
	//	return
	//})
	//if err != nil {
	//	log.Errorc(ctx, LogTemplate, canalConsumer.config.Id, "handlerDo", err)
	//	return
	//}
	if len(funcs) == 0 {
		return
	}
	var wg sync.WaitGroup
	lock := new(sync.Mutex)
	for _, handler := range funcs {
		f := handler
		wg.Add(1)
		go func() {
			_ = f(ctx, action, newInterface, oldInterface)
			defer func() {
				wg.Done()
			}()
			lock.Lock()
		}()
	}
	wg.Wait()
	return
}

//func oneTimeDo(ctx context.Context, handlers []BinlogConsumerHandlerFunc, action string, newInterface, oldInterface interface{}) (successFunc []BinlogConsumerHandlerFunc, failFunc []BinlogConsumerHandlerFunc, err error) {
//	successFunc = make([]BinlogConsumerHandlerFunc, 0)
//	failFunc = make([]BinlogConsumerHandlerFunc, 0)
//	errGroup := errgroupV2.WithContext(ctx)
//	lock := new(sync.Mutex)
//	for _, handler := range handlers {
//		f := handler
//		errGroup.Go(func(ctx context.Context) error {
//			gErr := f(ctx, action, newInterface, oldInterface)
//			lock.Lock()
//			if gErr != nil {
//				failFunc = append(failFunc, f)
//			} else {
//				successFunc = append(successFunc, f)
//			}
//			lock.Unlock()
//			return gErr
//		})
//	}
//	err = errGroup.Wait()
//	return
//}

func formatHandlerMapping(handlers []*handlerInfo) (handlersMap map[reflect.Type]*handlerInfo) {
	handlersMap = make(map[reflect.Type]*handlerInfo)
	for _, handler := range handlers {
		handlersMap[reflect.TypeOf(handler.doFunc)] = handler
	}
	return
}

func formatNewDoFunctions(handlers map[reflect.Type]*handlerInfo, failFuncs []BinlogConsumerHandlerFunc) (funcs []BinlogConsumerHandlerFunc) {
	funcs = make([]BinlogConsumerHandlerFunc, 0)
	for _, function := range failFuncs {
		if handlers[reflect.TypeOf(function)] != nil && handlers[reflect.TypeOf(function)].retry {
			funcs = append(funcs, function)
		}
	}
	return
}

func (canalConsumer *CanalConsumer) tableMatchCheck(tableName string) (tableForConfig string, err error) {
	handlers := canalConsumer.handlers
	if len(handlers) == 0 {
		err = errors.Errorf("config error")
		return
	}
	for tableC, handler := range handlers {
		switch handler.MatchMethod {
		case FullMatch:
			if tableC == tableName {
				tableForConfig = tableC
				return
			}
		case PrefixMatch:
			if strings.HasPrefix(tableName, tableC) {
				tableForConfig = tableC
				return
			}
		}
	}
	err = errors.Errorf("table no match")
	return
}

func handlersFind(handlersConfig *CanalConsumerHandler, action string, database string) (handlers []*handlerInfo) {
	handlers = make([]*handlerInfo, 0)
	midHandlers := make([]*handlerInfo, 0)
	switch action {
	case string(Insert):
		midHandlers = append(append(midHandlers, handlersConfig.insertHandlers...), handlersConfig.allHandlers...)
	case string(Update):
		midHandlers = append(append(midHandlers, handlersConfig.updateHandlers...), handlersConfig.allHandlers...)
	case string(Delete):
		midHandlers = append(append(midHandlers, handlersConfig.deleteHandlers...), handlersConfig.allHandlers...)
	}
	for _, handler := range midHandlers {
		if handler.database != "" && handler.database != database {
			continue
		}
		handlers = append(handlers, handler)
	}
	return
}

func (canalConsumer *CanalConsumer) Close() {
	if canalConsumer != nil && canalConsumer.reader != nil {
		canalConsumer.reader.Close()
	}
}
