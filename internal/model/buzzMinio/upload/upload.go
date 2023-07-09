package upload

import (
	"BuzzWaves/internal/model/buzzMinio"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
)

//UploadFile  上传文件minio
func UploadFile(uploadLocalPosition string, BucketName string, contentType string, objectName string) {
	// 设置上传文件的路径和目标桶名
	//filePath := "C:\\Users\\gongzhaowei\\Documents\\TencentMeeting\\2023-07-08 15.55.08 gonegone的快速会议 706578818\\meeting_01.mp4"
	//bucketName := "video"

	// 打开本地文件
	file, err := os.Open(uploadLocalPosition)
	if err != nil {
		fmt.Println("error 1", err)
		log.Fatalln(err)
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("error 2", err)
		log.Fatalln(err)
	}

	// 创建一个新的Minio对象
	//objectName := "gonevideo"
	//contentType := "video/mp4"

	// 使用PutObject上传文件

	n, err := buzzMinio.MinioClient.PutObject(context.Background(), BucketName, objectName, file, fileInfo.Size(), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		fmt.Println("上传文件失败", err)
		log.Fatalln(err)
	}

	log.Println("Successfully uploaded", objectName, "of size", n, "bytes.")
}
