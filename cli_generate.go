package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urfave/cli/v2"
)

// generate password
// 生成长度固定为 13 同时包含小写字母，大写字母，数字，特殊字符的随机字符串
func generate(ctx *cli.Context) error {
	const passwordLen = 13
	sed := rand.New(rand.NewSource(time.Now().UnixNano()))
	charSet := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"!@#$%^&*()_+",
	}

	// 生成 4 个非 0 随机数，和为 13
	charSetLen := len(charSet)
	charCont := make([]int, charSetLen)
	var sum = passwordLen
	for i := 0; i < charSetLen; i++ {
		// charSetLen-1-i 待生成的随机数个数 +1 保证生成的数不为 0
		charCont[i] = sed.Intn(sum-(charSetLen-1-i)) + 1
		sum -= charCont[i]
	}

	// 按照生成的随机数依次使用小写，大写，数字，特殊字符进行填充
	passwords := make([]byte, passwordLen)
	v := 0
	for i := 0; i < charSetLen; i++ {
		length := len(charSet[i])
		for j := 0; j < charCont[i]; j++ {
			passwords[v] = charSet[i][sed.Intn(length)]
			v++
		}
	}

	// 打乱密码顺序，进一步提高随机性，非必要
	for i := 0; i < 6; i++ {
		ri := sed.Intn(13)
		passwords[i], passwords[ri] = passwords[ri], passwords[i]
	}

	fmt.Println("gen: ", string(passwords))
	return nil
}
