package log

import (
	"os"
	"time"
)

type LocalHook struct {
}

func (hook *LocalHook) Write(p []byte) (n int, err error) {
	filename := time.Now().Format(time.DateOnly) + ".log"
	_ = appendToFile(filename, string(p))
	return len(p), err
}
func appendToFile(filename string, content string) error {
	// 打开文件，如果文件不存在则创建，使用追加模式
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入内容到文件
	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}
