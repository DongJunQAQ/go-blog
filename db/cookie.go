package db

import (
	"GoBlog/utils"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const KeyPrefix = "auth_cookie_"

var ctx = context.Background()

func SetCookieToRedis(cookieValue string, uid uint) { //把Cookie随机字符串作为键，uid作为值写入Redis
	redisClient := ConnectRedis()
	if err := redisClient.Set(ctx, KeyPrefix+cookieValue, uid, time.Hour*24).Err(); err != nil {
		utils.LogRus.Errorf("用户%d写入Cookie至Redis失败", uid)
	} else {
		utils.LogRus.Infof("用户%d写入Cookie至Redis成功:%s", uid, cookieValue)
	}
}
func GetCookieFromRedis(cookieValue string) (uid string) { //根据Cookie获取uid
	redisClient := ConnectRedis()
	uid, err := redisClient.Get(ctx, KeyPrefix+cookieValue).Result()
	if err != nil { //判断查询是否出错，如果有错再判断错误类型
		if !errors.Is(redis.Nil, err) { //如果不是"数据库查询无数据"的错误则打印日志
			utils.LogRus.Errorf("获取用户%s的验证信息失败:%s", uid, err)
		}
	} else {
		utils.LogRus.Infof("验证信息%s对应的用户为%s", KeyPrefix+cookieValue, uid)
	}
	return uid
}
