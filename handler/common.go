package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/pkg/orm"
	"net/http"
	"time"
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponsePageData struct {
	IndexPage orm.IndexPage `json:"index"`
	Res       interface{}   `json:"res"`
}

type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}

type ListReq struct {
	Page      uint64    `json:"page" form:"page" `                                          //页数
	Num       uint64    `json:"num"  form:"num" `                                           //数量
	Key       string    `json:"key" form:"key" `                                            //搜索关键字
	Sort      string    `json:"sort" form:"sort"`                                           //排序字段
	OrderType string    `json:"order_type" form:"order_type"`                               //排序规则
	BeginAt   time.Time `json:"begin_at" form:"begin_at" time_format:"2006-01-02 03:04:05"` //开始时间
	EndAt     time.Time `json:"end_at" form:"end_at" time_format:"2006-01-02 03:04:05"`     //结束时间
}

type IdReq struct {
	Id uint `form:"id" json:"id" uri:"id" binding:"required"`
}

func (l *ListReq) getListQuery(c *gin.Context) (err error) {
	err = c.ShouldBind(l)
	if err != nil {
		return err
	}
	//给默认参数
	if l.Page == 0 {
		l.Page = 1
	}
	if l.Num == 0 {
		l.Num = 10
	}
	if l.Sort == "" {
		l.Sort = "created_at"
	}
	if l.OrderType == "" {
		l.OrderType = "DESC"
	}
	if l.OrderType != "DESC" && l.OrderType != "ASC" {
		return errors.New("orderType 不是期望的值")
	}
	l.Sort = l.Sort + "  " + l.OrderType
	return
}

// 响应JSON数据
func resJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
}

// 响应成功
func resSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: http.StatusOK, Message: "ok", Data: v}
	resJSON(c, http.StatusOK, &ret)
}

// 响应成功
func resSuccessMsg(c *gin.Context) {
	ret := ResponseModel{Code: http.StatusOK, Message: "ok"}
	resJSON(c, http.StatusOK, &ret)
}

//参数错误
func resBadRequest(c *gin.Context, msg string) {
	ret := ResponseModel{Code: http.StatusBadRequest, Message: "参数绑定错误: " + msg}
	resJSON(c, http.StatusOK, &ret)
}

//业务错误
func resBusinessP(c *gin.Context, msg string) {
	ret := ResponseModel{Code: http.StatusBadGateway, Message: msg}
	resJSON(c, http.StatusOK, &ret)
}

//权限错误
func resUnauthorized(c *gin.Context) {
	ret := ResponseModel{Code: http.StatusUnauthorized, Message: "未授权"}
	resJSON(c, http.StatusOK, &ret)
}

// 响应错误-服务端故障
func resErrSrv(c *gin.Context) {
	ret := ResponseModel{Code: http.StatusInternalServerError, Message: "服务端故障"}
	resJSON(c, http.StatusOK, &ret)
}

// 响应成功-分页数据
func resSuccessPage(c *gin.Context, indexPage orm.IndexPage, list interface{}) {
	ret := ResponsePage{Code: http.StatusOK, Message: "ok", Data: ResponsePageData{IndexPage: indexPage, Res: list}}
	resJSON(c, http.StatusOK, &ret)
}
