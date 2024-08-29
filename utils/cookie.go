package utils

import (
	"math/rand"
)

var charSlice = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") //该变量包含了所有要用于生成随机字符串的字符，将该字符串转换为rune切片，这样每个字符都可以被独立处理。rune类型表示Unicode字符，适用于处理字符集。

func GenerateCookies(length int) string { //生成一段随机Cookie
	cookie := make([]rune, length)    //创建一个指定长度的rune切片，该切片用来存储生成的随机字符
	charSliceLength := len(charSlice) //字符集合的总数，获取有多少个字符
	for i := range cookie {
		cookie[i] = charSlice[rand.Intn(charSliceLength)] // 从charSlice切片中随机选择一个字符，并将其赋值给cookie[i]，rand.Intn(charSliceLength)生成一个范围在0到charSliceLength-1之间的随机整数，作为charSlice切片的索引
	}
	return string(cookie)
}
