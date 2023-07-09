package pkkg

import (
	"crypto/md5"
	"encoding/hex"
)

//func AesEncrypt(plaintext string, key []byte) (string, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//
//	paddedPlaintext, _ := pkcs7.Pad([]byte(plaintext), aes.BlockSize)
//
//	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
//	iv := ciphertext[:aes.BlockSize]
//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//		return "", err
//	}
//
//	mode := cipher.NewCBCEncrypter(block, iv)
//	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)
//
//	return hex.EncodeToString(ciphertext), nil
//}
func AesEncrypt(plaintext string) string {
	hasher := md5.New()

	// 将字符串转换为字节数组并进行加密
	hasher.Write([]byte(plaintext))

	// 获取加密后的字节数组
	encryptedBytes := hasher.Sum(nil)

	// 将字节数组转换为十六进制字符串表示
	encryptedString := hex.EncodeToString(encryptedBytes)

	return encryptedString
}
