package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/bytedance/sonic"
	"strings"
)

type JwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type JwtPayload struct {
	Sub string `json:"sub"` //主题
	Uid uint   `json:"uid"`
	Iat int    `json:"iat"` //发布时间戳
}

var (
	DefaultHeader = JwtHeader{Alg: "HS256", Typ: "JWT"}
	Secret        = "7900802"
)

func GenerateJWT(header JwtHeader, payload JwtPayload, secret string) (string, error) {
	var part1, part2, signature string
	//将header转换为JSON，然后进行Base64编码
	if bs1, err := sonic.Marshal(header); err != nil { //序列化，将结构体转换为JSON
		return "", err
	} else {
		part1 = base64.RawURLEncoding.EncodeToString(bs1) //将JSON编码为Base64格式，RawURLEncoding是一种Base64编码形式，它移除了填充字符=并将字符+和/替换为URL安全的-和_
	}
	//将payload转换为JSON，然后进行Base64编码
	if bs2, err := sonic.Marshal(payload); err != nil {
		return "", err
	} else {
		part2 = base64.RawURLEncoding.EncodeToString(bs2)
	}
	//将Base64编码后的header.payload哈希加密，然后再将密文Base64编码，最后得到签名
	h := hmac.New(sha256.New, []byte(secret))                    //创建HMAC(基于哈希的消息认证码)对象，算法为SHA-256，密钥为secret
	h.Write([]byte(part1 + "." + part2))                         //将Base64编码后的header.payload写入HMAC对象
	signature = base64.RawURLEncoding.EncodeToString(h.Sum(nil)) //进行加密并将加密后的密文Base64编码
	return part1 + "." + part2 + "." + signature, nil
}
func VerifyJWT(token, secret string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("无效token，只有%d部分", len(parts))
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(parts[0] + "." + parts[1]))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if signature != parts[2] {
		return nil, nil, fmt.Errorf("验证失败")
	}
	var part1, part2 []byte
	var err error
	if part1, err = base64.RawURLEncoding.DecodeString(parts[0]); err != nil {
		return nil, nil, fmt.Errorf("header Base64反解失败")
	}
	if part2, err = base64.RawURLEncoding.DecodeString(parts[1]); err != nil {
		return nil, nil, fmt.Errorf("payload Base64反解失败")
	}
	var header JwtHeader
	var payload JwtPayload
	if err = sonic.Unmarshal(part1, &header); err != nil {
		return nil, nil, fmt.Errorf("header JSON反序列化失败")
	}
	if err = sonic.Unmarshal(part2, &payload); err != nil {
		return nil, nil, fmt.Errorf("payload JSON反序列化失败")
	}
	return &header, &payload, nil
}
