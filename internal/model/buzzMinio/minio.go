package buzzMinio

import (
	"BuzzWaves/pkkg"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client
var err error
var DownloadFilePosition = " D:\\gocode\\BuzzWaves\\internal\\model\\buzzMinio\\download\\"

func GetMinioClient() {
	MinioClient, err = minio.New(pkkg.GetMinioEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(pkkg.GetMinioAccessKeyID(), pkkg.GetMinioSecretAccessKey(), ""),
		Secure: pkkg.GetMinioSecretUseSSL(),
	})
	if err != nil {
		fmt.Println("创建客户端失败")
		log.Fatalln(err)
	}
}
