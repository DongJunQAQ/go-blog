package db

import (
	"GoBlog/models"
	"GoBlog/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	myDbLog       ormlog.Interface //ORM日志接口
	blogMySQL     *gorm.DB         //为了使MySQL连接池单例模式，我们需要将其设置为全局变量，然后每次创建连接池时判断该变量中是否有值
	blogMySQLOnce sync.Once        //但是在高并发的情况下，我们需要开通协程来创建连接池，这就导致了每个协程都会创建一个连接池，也就破坏了单列模式，因此我们需要使用sync.Once
	blogRedis     *redis.Client
	blogRedisOnce sync.Once
	dbConfig      = utils.ReadConfig("db") //读取数据库配置文件
)

func init() { //优先执行该函数
	utils.InitLog("log")  //初始化自定义的LogRus日志
	myDbLog = ormlog.New( //自定义ORM日志
		log.New(os.Stdout, "\r\n", log.LstdFlags), //将日志打印至标准输出(终端)，分隔符使用回车换行符，日志记录日期和时间
		ormlog.Config{ //配置ORM日志
			SlowThreshold: 100 * time.Millisecond, //慢查询的阈值，查询时间超过100毫秒就被视为慢查询，慢查询指的是在数据库中执行时间较长的查询操作，慢查询在日志中会有特殊的标记
			LogLevel:      ormlog.Info,            //ORM日志级别
			Colorful:      true,                   //日志显示彩色
		},
	)
}
func CreateMysqlPool(host, user, password, dbname string, port int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname) //拼接后返回字符串
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: myDbLog, PrepareStmt: true})                                //使用我们自定义的ORM日志并启用SQL预编译以提高查询效率，将SQL语句与实际数据分开处理，在预编译阶段SQL语句被解析并生成执行计划而数据输入（参数）是在执行阶段提供的
	if err != nil {
		utils.LogRus.Panicf("%s连接MySQL失败:%s", dsn, err)
	}
	sqlDB, _ := db.DB() //设置数据库连接池，提高并发性能
	//连接池:在高并发情况下，有很多用户来访问网站，我们就需要开辟许多线程来同时做某件事情，所以我们需要在连接池中提前创建好连接
	sqlDB.SetMaxOpenConns(100) //连接池中最多有100个连接
	sqlDB.SetMaxIdleConns(20)  //连接池中最多保留20个空闲连接，避免有很多空闲连接占用资源
	utils.LogRus.Infof("已成功连接到MySQL的%s库", dbname)
	err = db.AutoMigrate(&models.User{}, &models.Blog{}) //迁移，根据结构体创建数据表
	if err != nil {
		utils.LogRus.Panicf("数据迁移失败:%s", err)
	}
	return db
}
func ConnectMySQL() *gorm.DB {
	blogMySQLOnce.Do(func() { //保证在高并发多协程情况下if代码块只被执行一次，确保连接池的单例模式
		if blogMySQL == nil { //如果该值为空则创建连接池
			host := dbConfig.GetString("mysql.host")
			user := dbConfig.GetString("mysql.user")
			password := dbConfig.GetString("mysql.password")
			dbname := dbConfig.GetString("mysql.dbname")
			port := dbConfig.GetInt("mysql.port")
			blogMySQL = CreateMysqlPool(host, user, password, dbname, port) //创建MySQL连接池，我们操作数据库时都需要调用该方法，而连接池只需创建一个，我们希望该池子是个单例模式即整个运行期间只创建一个连接池
		}
	})
	return blogMySQL
}
func CreateRedisClient(address, password string, db int) *redis.Client {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		utils.LogRus.Panicf("连接Redis失败:%s", err)
	} else {
		utils.LogRus.Infof("已成功连接到Redis的%d库", db)
	}
	return rdb
}
func ConnectRedis() *redis.Client {
	blogRedisOnce.Do(func() {
		if blogRedis == nil {
			host := dbConfig.GetString("redis.host")
			port := dbConfig.GetInt("redis.port")
			db := dbConfig.GetInt("redis.db")
			password := dbConfig.GetString("redis.password")
			redisAddress := fmt.Sprintf("%s:%d", host, port)
			blogRedis = CreateRedisClient(redisAddress, password, db)
		}
	})
	return blogRedis
}
