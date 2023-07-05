package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/tchuaxiaohua/cloud-station/store"
)

// NewAliOssStore AliOssStore对象构造函数
func NewAliOssStore(opts *store.Options) (*AliOssStore, error) {
	c, err := oss.New(opts.Endpoint, opts.SecretID, opts.SecretKey)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
	}, nil
}

type AliOssStore struct {
	client *oss.Client
}

func (a *AliOssStore) Upload(obj *store.StorageReq) (string, error) {
	bucket, err := a.client.Bucket(obj.BucketName)
	if err != nil {
		// HandleError(err)
		return "", err
	}
	// PutObjectFromFile 2个参数
	// 第一个 上传到oss后文件路径
	// 第二个 本地文件路径
	var cosFileKey string
	if obj.FilePrefix == "" {
		cosFileKey = obj.CloudFileName
	} else {
		cosFileKey = fmt.Sprintf("%s/%s", obj.FilePrefix, obj.CloudFileName)
	}
	err = bucket.PutObjectFromFile(cosFileKey, obj.LocalFilePath)
	if err != nil {
		// HandleError(err)
		return "", err
	}

	downUrl, err := bucket.SignURL(cosFileKey, oss.HTTPGet, 60*60*72)
	if err != nil {
		fmt.Printf("获取下载链接失败:%s\n", err)
		return "", err
	}
	return downUrl, nil
}
