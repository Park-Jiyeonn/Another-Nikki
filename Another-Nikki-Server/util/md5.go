package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func GetFileMd5(path string) (string, error) {
	// 文件全路径名
	pFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer pFile.Close()
	md5h := md5.New()
	_, err = io.Copy(md5h, pFile)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(md5h.Sum(nil)), nil
}
