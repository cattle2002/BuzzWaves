package download

import (
	"BuzzWaves/internal/model/buzzMinio"
	"context"
	"github.com/minio/minio-go/v7"
	//"fmt"
	"log"
)

// DownloadFile  下载文件
func DownloadFile(BucketName string, RemoteFileName string, NewLocalFile string) {

	// 设置下载文件的路径和目标桶名
	//objectName := "/gonevideo"
	//bucketName := "hello"
	//// 创建一个新的文件
	//filePath := "./hello2.mp4"
	// 使用FGetObjectEC函数下载文件并解码
	err := buzzMinio.MinioClient.FGetObject(context.Background(), BucketName, RemoteFileName, NewLocalFile, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully downloaded and decoded", RemoteFileName, "to", NewLocalFile)
}
