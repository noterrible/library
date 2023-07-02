package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"libraryManagementSystem/appV2/logic"
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/tools"
	_ "libraryManagementSystem/docs"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func New() *gin.Engine {
	r := gin.Default()
	go CronTask()          //定时任务刷新消息
	go Cacheheating(3, 10) //缓存预热
	r.Static("/static", "./appV2/resource")
	r.Use(limitedFlow(5, 10*time.Second))
	userRouter(r)
	adminRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//账号密码登录
	r.GET("/getCode", logic.GetCode)
	r.POST("/userLogin", logic.UserLogin)
	//手机验证码登录
	r.GET("/getPhoneCode", logic.GetPhoneCode)
	r.POST("/userLoginPhoneCode", logic.UserLoginPhoneCode)

	r.POST("/users", logic.AddUser)
	r.POST("/adminLogin", logic.LibrarianLogin)
	//游客可以浏览书籍和分类
	book := r.Group("/books")
	{
		book.GET("", logic.SearchBook)
		book.GET("/page", logic.PageBook)
		book.GET("/:id", logic.GetBook)
	}
	r.GET("/categories", logic.PageCategory)
	return r
}
func CronTask() {
	// 创建定时器，每隔半个小时执行一次
	t := time.NewTicker(time.Minute * 30)
	defer t.Stop()
	// 循环执行
	for {
		select {
		// 定时器触发时执行的操作
		case <-t.C:
			model.ListeningTask()
		}
	}

}

// 限制用户在一个端对单独的一个接口的访问
func limitedFlow(maxCount int, t time.Duration) gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取用户ip
		ip := context.ClientIP()
		//获取用户UA
		ua := context.GetHeader("User-Agent")
		//获取用户访问资源路径和请求参数，
		url := strings.Split(context.Request.URL.Path+"?"+context.Request.URL.RawQuery, "/")
		lastUrl := url[len(url)-1:]
		pathStr := fmt.Sprintf("%v_%v_%v", ip, ua, lastUrl)
		//fmt.Println("用户访问:", pathStr)
		requestQuery := model.StopRestartRequestConn
		//访问的路径次数加1
		requestQuery.Incr(context, pathStr)
		//获取访问次数
		reqCountString, _ := requestQuery.Get(context, pathStr).Result()
		reqCount, _ := strconv.Atoi(reqCountString)
		//超过最大次数，限制访问
		if reqCount > maxCount {
			context.JSON(http.StatusOK, tools.Response{
				Code:    tools.OK,
				Message: "请求太快，请休息一下重试~",
				Data:    nil,
			})
			context.Abort()
			return
		} else if reqCount == 1 { //第一次访问，设置过期时间为t
			// 设置键 "counter" 的过期时间为 120 秒
			if _, err := requestQuery.Expire(context, pathStr, t).Result(); err != nil {
				context.JSON(http.StatusOK, tools.Response{
					Code:    tools.OK,
					Message: "未知错误",
				})
				context.Abort()
				return
			}
		}
		//每次请求时，将请求url存到redis
		context.Next()
	}
}

// Cacheheating  缓存预热前n页数据，页大小为size
func Cacheheating(n, size int) {
	// 创建一个定时器，每 3秒触发一次
	ticker1 := time.NewTicker(3 * time.Second)
	defer ticker1.Stop()
	// 在goroutine中运行定时器
	for {
		select {
		case <-ticker1.C:
			for n, size = 1, 10; n <= 3; n++ { //预热前3页数据
				// 调用 Loader 函数加载图书信息
				turnPageInfo := model.TurnPageInfo{
					Id:               int64((n - 1) * size), //游标
					Page:             n,
					Limit:            size,
					PaginationMethod: "next",
					Sort:             "",
				}
				//游标查询分页的数据
				pageData, _ := model.TurnPages(model.BookInfo{}, turnPageInfo)
				//redis键
				key := fmt.Sprintf("books_%v_%v_%v", turnPageInfo.Page, turnPageInfo.Limit, turnPageInfo.Sort)
				/*查的数据存入redis,失效时间设置为随机4~8*/
				//设置随机数种子
				rand.Seed(time.Now().UnixNano())
				//失效时间 将查出的数据存到redis
				randInt := rand.Intn(4) + 4
				loseTime := time.Duration(randInt) * time.Second
				//压缩
				var pageGzipData string
				var errTG error
				if pageGzipData, errTG = tools.ToGzip(pageData); errTG != nil {
					fmt.Println("压缩错误", errTG.Error())
					return
				}
				model.InfoCacheRedisConn.Set(context.Background(), key, pageGzipData, loseTime)
			}
		}
	}
}
