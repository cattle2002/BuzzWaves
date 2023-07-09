package biz

import (
	"BuzzWaves/internal/model/buzzMinio/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将文件保存到指定的目录
	// 生成文件名
	extension := filepath.Ext(file.Filename)
	fileName := "uploaded_file" + extension
	// 指定保存的目录
	uploadPath := "D:\\gocode\\BuzzWaves\\internal\\model\\buzzMinio\\upload\\" + fileName
	fmt.Println(uploadPath)

	err = c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//upload.UploadFile("D:\\gocode\\BuzzWaves\\internal\\model\\buzzMinio\\upload\\uploaded_file.pak", "hello", "text/plain", fileName)
	upload.UploadFile(uploadPath, "hello", "text/plain", fileName)
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})

}
