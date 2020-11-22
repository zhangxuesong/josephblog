package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/handler"
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

	apiv1 := router.Group("/api/v1")
	tagC := handler.TagController{}
	{
		apiv1.GET("/tags", tagC.List)
		apiv1.POST("/tag", tagC.Create)
		apiv1.PUT("/tag/:id", tagC.Update)
		apiv1.DELETE("/tag/:id", tagC.Delete)
	}
	return router
}
