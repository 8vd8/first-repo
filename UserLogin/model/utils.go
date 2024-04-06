package model

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordHash(pwd string) string {
	//创建一个MD5哈希器
	hasher := md5.New()
	hasher.Write([]byte(pwd))                      //将密码转换成字节数组并写入哈希器
	pwdHash := hex.EncodeToString(hasher.Sum(nil)) //计算哈希值并将结果转成十六进制字符串
	return pwdHash                                 //返回哈希值
}
