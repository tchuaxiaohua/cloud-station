package tx

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/tchuaxiaohua/cloud-station/store"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// TxCosStore 腾讯cos客户端结构体对象
type TxCosStore struct {
	client *cos.Client
}

func NewTxCosStore(opts *store.Options) (*TxCosStore, error) {
	u, _ := url.Parse(opts.Endpoint)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: opts.SecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: opts.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return &TxCosStore{
		client: c,
	}, nil
}

func (t *TxCosStore) Upload(obj *store.StorageReq) (string, error) {
	var cosFileKey string
	if obj.FilePrefix == "" {
		cosFileKey = obj.CloudFileName
	} else {
		cosFileKey = fmt.Sprintf("%s/%s", obj.FilePrefix, obj.CloudFileName)
	}
	res, _, err := t.client.Object.Upload(context.Background(), cosFileKey, obj.LocalFilePath, nil)
	if err != nil {
		log.Print("Error uploading", err)
	}
	return res.Location, nil
}
