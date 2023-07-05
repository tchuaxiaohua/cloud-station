package store

// Uploader 定义文件上传接口
// 所有云商都需要实现该接口签名
type Uploader interface {
	Upload(s *StorageReq) (string, error)
}

// StorageReq 存储请求参数
// BucketName 桶名称
// FilePrefix 文件前缀
// OssFilePath 上传云端后文件名称
// FileName 本地待上传文件路径
type StorageReq struct {
	BucketName    string
	FilePrefix    string
	CloudFileName string
	LocalFilePath string
}

// NewStorageReq 初始化函数
func NewStorageReq(bucketName, filePrefix, cloudFileName, uploadFileName string) *StorageReq {
	return &StorageReq{
		BucketName:    bucketName,
		FilePrefix:    filePrefix,
		CloudFileName: cloudFileName,
		LocalFilePath: uploadFileName,
	}
}

// Options 云商基础参数 结构体对象
// Endpoint bucket或者cos对应的 地域
// SecretID/SecretKey 对应的ak sk
type Options struct {
	Endpoint  string
	SecretID  string
	SecretKey string
}

func NewOptions(endpoint, secretAk, secretSk string) *Options {
	return &Options{
		Endpoint:  endpoint,
		SecretID:  secretAk,
		SecretKey: secretSk,
	}
}
