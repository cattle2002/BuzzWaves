package pkkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/mergermarket/go-pkcs7"
)

func AesDecrypt(ciphertext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	encrypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(encrypted) < aes.BlockSize {
		return "", fmt.Errorf("密文长度不足")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	if len(encrypted)%aes.BlockSize != 0 {
		return "", fmt.Errorf("密文长度不是块大小的倍数")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encrypted, encrypted)

	unpaddedPlaintext, err := pkcs7.Unpad(encrypted, aes.BlockSize)
	if err != nil {
		return "", err
	}

	return string(unpaddedPlaintext), nil
}
