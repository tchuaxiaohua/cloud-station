package cmd

import (
	"fmt"
	"log"

	"github.com/tchuaxiaohua/cloud-station/store"
	"github.com/tchuaxiaohua/cloud-station/store/aliyun"
	"github.com/tchuaxiaohua/cloud-station/store/tx"
	"github.com/tchuaxiaohua/cloud-station/utils"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	//Use:   "start",
	Short: "Usage: ",
	Long:  "Usage: cloud-station COMMAND",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 配置文件加载
		if config != "" {
			loadConfig()
		}
		// 变量加载
		initEnvValue()

		// 文件判断
		ok := utils.PathExists(localFileName)
		// 初始化云商参数
		cloudOptions := store.NewOptions(endPoint, secretID, secretKey)
		// 初始化 桶参数配置
		cloudStore := store.NewStorageReq(bucketName, bucketPrefix, utils.FileName(localFileName), localFileName)
		fmt.Println(bucketName)
		if ok {
			var fileUrl string
			switch cloudtype {
			case "aliyun":
				// 初始化云商参数
				aliClient, err := aliyun.NewAliOssStore(cloudOptions)
				if err != nil {
					log.Fatal("初始化错误", err)
				}
				// 初始化云存储
				fileUrl, err = aliClient.Upload(cloudStore)
				if err != nil {
					log.Fatal("文件上传失败", err)
				}
				log.Printf("文件上传成功,上传地址:%s", fileUrl)
			case "txyun":
				txClient, err := tx.NewTxCosStore(cloudOptions, bucketPrefix)
				if err != nil {
					log.Fatal("初始化错误", err)
				}
				// 初始化云存储
				fileUrl, err = txClient.Upload(cloudStore)

				if err != nil {
					log.Fatal("文件上传失败", err)
				}
				log.Printf("文件上传成功,上传地址:%s", fileUrl)
			default:
				fmt.Println("未知云商,支持阿里/腾讯【aliyun/txyun】")
			}

			// 钉钉通知
			msg := utils.SendDingDing(token, title, fileUrl)
			log.Println("钉钉发送状态:", msg)
		}
		return nil
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", "", "Bucker/Cos 对应的地域")
	StartCmd.PersistentFlags().StringVarP(&secretID, "access-key-id", "i", "", "AccessKey ID")
	StartCmd.PersistentFlags().StringVarP(&secretKey, "access-key-secret", "k", "", "AccessKey Secret")
	StartCmd.PersistentFlags().StringVarP(&cloudtype, "cloud", "C", "", "指定云商【aliyun/txyun】")
	StartCmd.PersistentFlags().StringVarP(&bucketName, "bucket", "b", "", "指定bucket名称")
	StartCmd.PersistentFlags().StringVarP(&bucketPrefix, "prefix", "p", "", "指定上传后文件路径前缀")
	StartCmd.PersistentFlags().StringVarP(&localFileName, "localpath", "f", "", "指定本地上传文件")
	StartCmd.PersistentFlags().StringVarP(&token, "dingtoken", "t", "", "钉钉token")
	StartCmd.PersistentFlags().StringVarP(&title, "title", "T", title, "钉钉通知主题")
	StartCmd.PersistentFlags().StringVarP(&title, "config", "c", "", "指定配置文件")
}
