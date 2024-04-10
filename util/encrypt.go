package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// TODO: 添加到生成的密文当中，且每次随机生成
var _iv = []byte{0xa3, 0x38, 0x9e, 0x48, 0xb5, 0xe6, 0xba, 0x4, 0xca, 0xd4, 0xcb, 0x47, 0x88, 0xa5, 0xc5, 0x41}

// 计算文件 sha256sum 作为密钥
func FileHash256(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	hash := sha256.New()
	_, err = f.WriteTo(hash)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

func EncrypeString(msg string, keyFile string) (string, error) {
	keyByte, err := FileHash256(keyFile)
	if err != nil {
		return "", err
	}
	msgByte := []byte(msg)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(msgByte))
	stream := cipher.NewCFBEncrypter(block, _iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], msgByte)
	return hex.EncodeToString(ciphertext), nil
}

func DecryptString(msg string, keyFile string) (string, error) {
	passwordByte, err := hex.DecodeString(msg)
	if err != nil {
		return "", err
	}
	keyByte, err := FileHash256(keyFile)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	plaintext := make([]byte, len(passwordByte)-aes.BlockSize)
	stream := cipher.NewCFBDecrypter(block, _iv)
	stream.XORKeyStream(plaintext, passwordByte[aes.BlockSize:])
	return string(plaintext), nil
}
