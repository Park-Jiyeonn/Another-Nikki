package util

import (
	"Another-Nikki/config"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func UploadCos(path, name string) (string, error) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	rawURL := "https://jiyeon-1317677590.cos.ap-shanghai.myqcloud.com"
	u, _ := url.Parse(rawURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: config.SecretId,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: config.SecretKey,
		},
	})

	md5String, err := GetFileMd5(path)
	if err != nil {
		return "", err
	}

	key := "Another-Nikki" + "/" + md5String + "-" + name

	_, err = client.Object.PutFromFile(context.Background(), key, path, nil)
	if err != nil {
		return "", err
	}
	return rawURL + "/" + key, nil
}
