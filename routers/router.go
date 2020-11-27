package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/handler"
	"github.com/zhangxuesong/josephblog/middleware"
	"github.com/zhangxuesong/josephblog/pkg/config"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Config.Runmode)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin", "*")                   //跨域
		//c.Header("Access-Control-Allow-Headers", "token,Content-Type") //必须的请求头
		//c.Header("Access-Control-Allow-Methods", "OPTIONS,POST,GET,PUT,DELETE")   //接收的请求方法
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "token, origin, content-type, accept")
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(200)
		}
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	apiPrefix := "/api"
	adminC := handler.AdminController{}
	router.POST(apiPrefix+"/login", adminC.Login)
	admin := router.Group(apiPrefix + "/admin")
	admin.Use(middleware.JWTAuth())
	admin.GET("/info", adminC.Info)
	tagC := handler.TagController{}
	{
		admin.GET("/tags", tagC.List)
		admin.POST("/tag", tagC.Create)
		admin.PUT("/tag/:id", tagC.Update)
		admin.DELETE("/tag/:id", tagC.Delete)
	}
	articleC := handler.ArticleController{}
	{
		admin.GET("/articles", articleC.List)
		admin.GET("/article/:id", articleC.Detail)
		admin.POST("/article", articleC.Create)
		admin.PUT("/article/:id", articleC.Update)
		admin.DELETE("/article/:id", articleC.Delete)
	}
	return router
}
