package routers

import (
	"time"

	"github.com/easonnong/front-end-demo/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()
	//实现跨域访问
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站
		AllowAllOrigins: true,
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享，确定共享
		AllowCredentials: true,
		//容许跨域的原点网站，可以直接return true
		/* AllowOriginFunc: func(origin string) bool {
			return true
		}, */
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})

	//gin里面使用中间件
	engine.Use(mwCORS)

	engine.GET("/users", api.ListUser)
	engine.POST("/add", api.AddUser)
	engine.Run()
}
