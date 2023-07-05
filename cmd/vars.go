package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	secretID      = ""
	secretKey     = ""
	endPoint      = ""
	bucketName    = ""
	bucketPrefix  = ""
	localFileName = ""
)
var (
	config    = ""
	token     = ""
	cloudtype = ""
	title     = ""
)

// initEnvValue 变量初始化配置 命令行参数 > 环境变量 > 配置文件
func initEnvValue() {
	if secretID == "" {
		if os.Getenv("SECRETID") != "" {
			secretID = os.Getenv("SECRETID")
		} else {
			secretID = viper.GetString("cloud.secretID")
		}
	}

	if secretKey == "" {
		if os.Getenv("SECRETKEY") != "" {
			secretKey = os.Getenv("SECRETKEY")
		} else {
			secretKey = viper.GetString("cloud.secretKey")
		}
	}

	if endPoint == "" {
		if os.Getenv("ENDPOINT") != "" {
			endPoint = os.Getenv("ENDPOINT")
		} else {
			endPoint = viper.GetString("cloud.endpoint")
		}
	}

	if bucketName == "" {
		if os.Getenv("BUCKET_NAME") != "" {
			bucketName = os.Getenv("BUCKET_NAME")
		} else {
			bucketName = viper.GetString("cloud.bucketName")
		}
	}
	if bucketPrefix == "" {
		if os.Getenv("BUCKET_PREFIX") != "" {
			bucketPrefix = os.Getenv("BUCKET_PREFIX")
		} else {
			bucketPrefix = viper.GetString("cloud.bucketPrefix")
		}
	}

	if localFileName == "" {
		if os.Getenv("LOCAL_FILENAME") != "" {
			localFileName = os.Getenv("LOCAL_FILENAME")
		} else {
			localFileName = viper.GetString("cloud.localFilePath")
		}
	}

	if token == "" {
		if os.Getenv("TOKEN") != "" {
			token = os.Getenv("TOKEN")
		} else {
			token = viper.GetString("dingding.token")
		}
	}

	if cloudtype == "" {
		if os.Getenv("CLOUD_TYPE") != "" {
			cloudtype = os.Getenv("CLOUD_TYPE")
		} else {
			cloudtype = viper.GetString("cloudtype")
		}
	}

	if title == "" {
		if os.Getenv("TITLE") != "" {
			title = os.Getenv("TITLE")
		} else {
			title = viper.GetString("dingding.title")
		}
	}
}

func loadConfig() {
	configPath := filepath.Dir(config)
	filename := filepath.Base(config)
	strings.Split(filename, ".")
	viper.SetConfigName(strings.Split(filename, ".")[0])
	viper.SetConfigType(strings.Split(filename, ".")[1])
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
}
