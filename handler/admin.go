package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/models"
	"github.com/zhangxuesong/josephblog/pkg/jwt"
	"github.com/zhangxuesong/josephblog/pkg/orm"
	"github.com/zhangxuesong/josephblog/pkg/redis"
	"log"
	"strconv"
)

type AdminController struct {
	// 管理员控制器
}

type login struct {
	Username string `json:"username" binding:"required" ` //账号
	Password string `json:"password" binding:"required" ` // 密码
}

// @Title 登录接口
// @Summary 用户登录
// @Tags admin   管理员操作
// @Accept  json
// @Produce  json
// @Param   body    body    handler.login   false   "body"
// @Success 200 {object}  handler.ResponseModel 	"{code:200,msg:ok,data:{token:}}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Router /login [post]
func (AdminController) Login(c *gin.Context) {
	reqData := login{}
	err := c.ShouldBind(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	admin := models.Admin{}
	notFound, err := orm.First(models.Admin{Username: reqData.Username}, &admin)
	if err != nil {
		if notFound {
			log.Println("没有此条记录 %v", reqData.Username)
			resBadRequest(c, "没有此条记录")
			return
		}
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	if reqData.Password != admin.Password {
		resBusinessP(c, "密码错误")
		return
	}
	adminId := strconv.FormatInt(int64(admin.ID), 10)
	token, err := jwt.CreateToken(adminId)
	if err != nil {
		resBusinessP(c, err.Error())
		return
	}
	var fields = make(map[string]interface{}, 10)
	fields["id"] = admin.ID
	_, err = redis.BatchHashSet(redis.Redis, "access_token_"+token, fields)
	if err != nil {
		resBusinessP(c, err.Error())
		return
	}
	resSuccess(c, gin.H{"token": token})
}
