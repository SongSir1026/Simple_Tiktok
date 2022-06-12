package utils

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

type OssServer struct {
}

var once sync.Once
var ossServerInstance *OssServer

func NewOssServer() *OssServer {
	once.Do(func() {
		ossServerInstance = &OssServer{}
	})
	return ossServerInstance
}
func (that *OssServer) Init(accessKeyId, accessKeySecret, bucketName, endPoint string) (client *oss.Client, bucket *oss.Bucket, err error) {
	client, err = oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return
	}
	if bucketName != "" {
		bucket, err = client.Bucket(bucketName)
		if err != nil {
			return
		}
	}
	return
}
func UploadFile(fileName string, fileByte []byte, userName string) (url string, err error) {
	accessKeyId := "LTAI5tLU1XR7rZHLHwaWVc2g"
	accessKeySecret := "pBGt9O3lnDZpW0Dzs6qUirpOS6BHDw"
	bucketName := "simplesiktok-songsir"
	endPoint := "https://oss-cn-beijing.aliyuncs.com"
	path := "https://simplesiktok-songsir.oss-cn-beijing.aliyuncs.com"
	_, bucket, err := NewOssServer().Init(accessKeyId, accessKeySecret, bucketName, endPoint)
	if err != nil {
		fmt.Println(err)
	}

	//设置路径
	folderName := userName
	yunFileTmpPath := "uploads" + folderName + "/" + fileName
	err = bucket.PutObject(yunFileTmpPath, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		return url, err
	}

	return path + "/" + yunFileTmpPath, nil
	//uploads\13393413460/share_3c9ace42965fb27e3debb96de01a5499.mp4
	//
	//https://simplesiktok-songsir.oss-cn-beijing.aliyuncs.com/uploads%5C13393413460/share_3c9ace42965fb27e3debb96de01a5499.mp4
	//https://oss-cn-beijing.aliyuncs.com/uploads/13393413460/share_3c9ace42965fb27e3debb96de01a5499.mp4
}
