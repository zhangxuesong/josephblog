package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/pkg/jwt"
	"github.com/zhangxuesong/josephblog/pkg/redis"
	"net/http"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		//中心化的管理端需要从缓存中取这个token是否过期
		_, err = redis.Redis.Exists(token).Result()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 4,
				"msg":  "token 失效",
			})
			c.Abort()
			return
		}
		if claims.UID == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效 token",
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("UID", claims.UID)
	}
}
