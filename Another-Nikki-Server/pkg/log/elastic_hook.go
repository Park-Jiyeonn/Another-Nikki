package log

import (
	"bytes"
	"context"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// ElasticsearchHook 实现了zapcore.WriteSyncer接口，用于将日志消息发送到Elasticsearch
type ElasticsearchHook struct {
	client    *elastic.TypedClient
	indexName string
}

// NewElasticsearchHook 创建一个新的ElasticsearchHook
func NewElasticsearchHook(client *elastic.TypedClient, indexName string) (*ElasticsearchHook, error) {
	return &ElasticsearchHook{
		client:    client,
		indexName: indexName,
	}, nil
}

// Write 方法将日志消息写入Elasticsearch
func (hook *ElasticsearchHook) Write(p []byte) (n int, err error) {
	req := esapi.IndexRequest{
		Index:      hook.indexName,
		DocumentID: "", // Elasticsearch会自动生成文档ID
		Body:       bytes.NewReader(p),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), hook.client)
	defer res.Body.Close()
	//fmt.Println(res, err)
	//fmt.Println("haha", string(p))
	return len(p), err
}

// Sync 方法实现zapcore.WriteSyncer接口中的Sync方法
func (hook *ElasticsearchHook) Sync() error {
	// 在这里可以添加同步逻辑，如果有的话
	return nil
}
