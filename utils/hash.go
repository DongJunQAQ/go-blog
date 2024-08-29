package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string { //不管输入的字符串长度为多少最终输出的都是32位的字符串
	hash := md5.New()                 //初始化MD5哈希计算器
	hash.Write([]byte(text))          //将字符串转换为字节切片，并将其写入hash对象中
	digest := hash.Sum(nil)           //生成哈希值的字节切片，md5哈希的结果为128bit
	return hex.EncodeToString(digest) //将哈希字节切片转换为十六进制字符串，并返回该字符串，十六进制等于4bit，所以返回值为128/4=32bit
}
