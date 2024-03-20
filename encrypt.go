package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

// TODO: 添加到生成的密文当中，且每次随机生成
var iv = []byte{0xa3, 0x38, 0x9e, 0x48, 0xb5, 0xe6, 0xba, 0x4, 0xca, 0xd4, 0xcb, 0x47, 0x88, 0xa5, 0xc5, 0x41}

// 因为是命令行程序，所以遇到错误直接退出，不进行其他行为
func ErrorWithEixt(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 使用 aes256 进行加密，key 长度必须是 32
func Encrypt(msg string, key []byte) []byte {
	if len(key) < 32 {
		ErrorWithEixt(errors.New("key is too short"))
	}

	block, err := aes.NewCipher(key)
	ErrorWithEixt(err)
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))

	return ciphertext
}

// 对密文进行 aes256 解密
func Decrypt(msgg string, key []byte) string {
	msg, err := hex.DecodeString(msgg)
	ErrorWithEixt(err)
	if len(key) < 32 {
		ErrorWithEixt(errors.New("key is too short"))
	}

	block, err := aes.NewCipher(key)
	ErrorWithEixt(err)
	plaintext := make([]byte, len(msg)-aes.BlockSize)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, msg[aes.BlockSize:])

	return string(plaintext)
}

// 计算文件 sha256sum 作为密钥
func FileHash256(file string) []byte {
	f, err := os.Open(file)
	ErrorWithEixt(err)

	hash := sha256.New()
	_, err = f.WriteTo(hash)
	ErrorWithEixt(err)

	hashByte := hash.Sum(nil)
	return hashByte
}
