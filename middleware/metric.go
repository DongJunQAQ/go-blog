package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strings"
	"time"
)

const SERVICE = "blog"

var (
	requestCounter = promauto.NewCounterVec(prometheus.CounterOpts{Name: "request_counter"}, []string{"service", "interface"}) //统计接口被调用的次数
	//promauto.NewCounterVec(): 创建计数器向量，用于计算事件的发生次数
	//prometheus.CounterOpts{Name: "request_counter"}: 指定计数器的名字
	//[]string{"service", "interface"}: 计数器的标签，用来区分不同的服务和接口
	requestTimer = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "request_timer"}, []string{"service", "interface"}) //统计接口的响应时间
	//promauto.NewGaugeVec: 创建仪表向量，用于测量某些量值的当前状态，比如请求的响应时间
	restfulMapping = map[string]string{"uid": ":uid", "bid": ":bid"}
)

func ConvertUrl(ctx *gin.Context) string { //将实际的URL路径转换为符合RESTful路径参数格式的URL，如将/blog/3转换为/blog/:bid
	url := ctx.Request.URL.Path    //提取原始的请求路径，如取出/blog/3
	for _, p := range ctx.Params { //p.Key：表示路径参数的名称（例如 bid）；p.Value：表示实际的路径参数值（例如 3）
		value, exists := restfulMapping[p.Key] //如果请求路径为/blog/3，那么p.Key为bid，然后去restfulMapping里面查询并将:bid赋值给value
		if exists {                            //如果restfulMapping中有与p.Key匹配的条目则执行下一步
			url = strings.Replace(url, p.Value, value, 1) //替换，p.Value是当前参数的实际值（如3），value是映射中对应的路径参数格式（如 :bid），1表示仅替换第一次出现的匹配值
		}
	}
	return url
}
func Metric() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		begin := time.Now()             //记录该请求的开始时间
		ctx.Next()                      //执行后续handler处理程序
		convertedURL := ConvertUrl(ctx) //转换请求的路径，convertedURL存储转换后的URL路径
		requestCounter.WithLabelValues(SERVICE, convertedURL).Inc()
		//WithLabelValues(SERVICE, convertedURL)：根据服务名称和转换后的URL路径设置计数器的标签值，如request_counter{interface="/blog/:bid",service="blog"}
		//Inc()：将计数器的值增加1
		requestTimer.WithLabelValues(SERVICE, convertedURL).Set(float64(time.Since(begin).Milliseconds()))
		//time.Since(begin).Milliseconds()：计算从begin时间到当前时间的间隔并将结果转换为毫秒
		//Set()：将仪表的值设置为计算得到的响应时间
	}
}
