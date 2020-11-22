package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/models"
	"github.com/zhangxuesong/josephblog/pkg/orm"
	"log"
)

type TagController struct {
	// 文章标签控制器
}

// 获取标签列表
// @Summary 获取标签列表
// @Tags   tag  标签管理
// @Accept  json
// @Produce  json
// @Param   page         query    int       false     "页码,默认为1"
// @Param   num          query    int       false     "返回条数,默认为10"
// @Param   sort         query    string    false     "排序字段,默认为created_at"
// @Param   key          query    string    false     "搜索关键字"
// @Param   orderType    query    string    false     "排序规则,默认为DESC"
// @Param   beginAt      query    string    false     "开始时间"
// @Param   endAt        query    string    false     "结束时间"
// @Success 200 {array}   models.Tag 	"标签列表"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /tags [get]
func (TagController) List(c *gin.Context) {
	reqData := ListReq{}
	err := reqData.getListQuery(c)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	var whereOrder []orm.PageWhere
	if reqData.Key != "" {
		v := "%" + reqData.Key + "%"
		var arr []interface{}
		arr = append(arr, v)
		whereOrder = append(whereOrder, orm.PageWhere{Where: " name like ? ", Value: arr})
	}
	if !reqData.BeginAt.IsZero() {
		var arr []interface{}
		arr = append(arr, reqData.BeginAt)
		whereOrder = append(whereOrder, orm.PageWhere{Where: " created_at > ? ", Value: arr})
	}
	if !reqData.EndAt.IsZero() {
		var arr []interface{}
		arr = append(arr, reqData.EndAt)
		whereOrder = append(whereOrder, orm.PageWhere{Where: " created_at < ? ", Value: arr})
	}
	list := make([]models.Tag, 0)
	var indexPage orm.IndexPage
	indexPage.Page = reqData.Page
	indexPage.Num = reqData.Num
	err = orm.GetPage(&models.Tag{}, &models.Tag{}, &list, &indexPage, reqData.Sort, whereOrder...)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccessPage(c, indexPage, list)
}

type tagCreateReq struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=10"` //标签名称
}

// 创建标签
// @Summary 创建标签
// @Tags   tag  标签管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.tagCreateReq    true     "标签信息"
// @Success 200 {object}   handler.ResponseModel 	"{code:200,msg:ok,data:{id:"id"}}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /tag [post]
func (TagController) Create(c *gin.Context) {
	reqData := tagCreateReq{}
	err := c.ShouldBind(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	tag := models.Tag{}
	tag.Name = reqData.Name
	err = orm.Create(&tag)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccess(c, gin.H{"id": tag.ID})
}

type tagUpdateReq struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=10"` //标签名称
}

// 修改标签
// @Summary 修改标签
// @Tags   tag  标签管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.IdReq    true     "标签ID"
// @Param   body    body    handler.tagUpdateReq    true     "标签信息"
// @Success 200 {object}   handler.ResponseModel 	"{code:200,msg:ok}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /tag/:id [put]
func (TagController) Update(c *gin.Context) {
	reqUriData := IdReq{}
	reqData := tagUpdateReq{}
	err := c.ShouldBindUri(&reqUriData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	err = c.ShouldBind(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	where := models.Tag{}
	where.ID = reqUriData.Id
	err = orm.Updates(&where, &reqData)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccessMsg(c)
}

// 删除标签
// @Summary 删除标签
// @Tags   tag  标签管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.IdReq    true     "标签ID"
// @Success 200 {object}   handler.ResponseModel 	"{code:200,msg:ok}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /tag/:id [delete]
func (TagController) Delete(c *gin.Context) {
	reqData := IdReq{}
	err := c.ShouldBindUri(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	_, err = orm.DeleteByID(models.Tag{}, reqData.Id)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccessMsg(c)
}
