package tools

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 手机号登陆自动生成用户密码
func GenPass() string {
	passwordLength := 8
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var password []byte
	for i := 0; i < passwordLength; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		password = append(password, charset[index.Int64()])
	}
	fmt.Println(string(password))
	return string(password)
}
