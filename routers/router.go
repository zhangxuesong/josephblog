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
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	apiPrefix := "/api"
	admin := router.Group(apiPrefix + "/v1")
	admin.Use(middleware.JWTAuth())
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
