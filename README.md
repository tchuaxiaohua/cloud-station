#### 一、项目背景

---

`cloud-station`云商中转站，主要用来在Docker或者pod环境时，我们想下载容器中文件时，可以直接使用该脚本，把文件上传至云存储，然后再进行下载，这样我们无需对容器做任何操作，把该脚本copy至容器即可，更多使用说明参考[cloud-station云商中转站](https://wiki.tbchip.com//pages/821ce2/)

#### 二、项目介绍

##### 2.1 使用说明

~~~sh
# 拉取代码
https://github.com/tchuaxiaohua/cloud-station.git

# 编译启动(依赖go环境)
cd cloud-station
go build  -ldflags "-s -w" -o cloud-station
chmod +x cloud-station
./cloud-station --help
~~~

##### 2.2 支持参数

> 我们把项目打包成二进制脚本后，可以通过帮助查看支持的参数

~~~sh
root@b682544d4097:/go/cloud-station# ./cloud-station --help
Usage: cloud-station COMMAND

Usage:
   [flags]

Flags:
  -i, --access-key-id string       AccessKey ID
  -k, --access-key-secret string   AccessKey Secret
  -b, --bucket string              指定bucket名称
  -C, --cloud string               指定云商【aliyun/txyun】
  -c, --config string              指定配置文件
  -t, --dingtoken string           钉钉token
  -e, --endpoint string            Bucker/Cos 对应的地域
  -h, --help                       help for this command
  -f, --localpath string           指定本地上传文件
  -p, --prefix string              指定上传后文件路径前缀
  -T, --title string               钉钉通知主题
~~~

**参数说明**

* `-i` 云商子账号ak 需要对存储服务有读写权限
  * 对应环境变量`SECRETID`
  * 对应配置文件`secretID`
* `-k` 云商子账号sk 需要对存储服务有读写权限
  * 对应环境变量`SECRETKEY`
  * 对应配置文件`secretKey`
* `-b` 云商云存储名称，这里只有阿里云的时候需要，腾讯云或自动获取
  * 对应环境变量`BUCKET_NAME`
  * 对应配置文件`bucketName`
* `-C` 指定云商，目前仅支持腾讯和阿里
  * 对应环境变量`CLOUD_TYPE`
  * 对应配置文件`cloudtype`
* `-c` 配置文件，该工具支持三种方式对参数赋值，后面会详细说明
* `-t` 钉钉通知时候机器人token
  * 对应环境变量`TOKEN`
  * 对应配置文件`token`
* `-e` 云存储地域
  * 对应环境变量`ENDPOINT`
  * 对应配置文件`endpoint`
* `-f` 需要上传的本地文件，不支持文件夹
  * 对应环境变量`LOCAL_FILENAME`
  * 对应配置文件`localFilePath`
* `-p` 上传至云存储之后的路径前缀，比如`${endpoint}/doc/pic.jpg` ,前缀为`doc`
  * 对应环境变量`BUCKET_PREFIX`
  * 对应配置文件`bucketPrefix`
* `-T` 钉钉通知时候的主题
  * 对应环境变量`TITLE`
  * 对应环境变量`title`

##### 2.3 变量三种注入方式

###### 2.3.1 命令行参数传递

> 如果只是临时使用，推荐使用这种方式

命令行参数传递的时候，需要注意，上面的参数除了`title`，其他都是必传，还需要注意云商类型

###### 2.3.2 环境变量方式

> 环境变量就是配置系统变量

系统变量方式，比较适合长期使用

###### 2.3.4 配置文件方式

配置文件，需要我们先创建一个`app.yaml`格式的文件，创建的时候确保该文件与脚本在同目录下，否则就需要使用`-c`参数指定，如果使用配置文件，则配置文件中的参数需要全部填写。

* 配置文件示例

```yaml
cloud:
  secretID: "ak"
  secretKey: "sk"
  endpoint: "oss-cn-hangzhou.aliyuncs.com"
  bucketName: "bucket名称"
  bucketPrefix: "test"
  localFilePath: "上传文件"

dingding:
  token: "钉钉机器人docker"
  title: "文件下载通知"

cloudtype: "aliyun"
```

