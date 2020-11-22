package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangxuesong/josephblog/models"
	"github.com/zhangxuesong/josephblog/pkg/orm"
	"log"
)

type ArticleController struct {
	// 文章控制器
}

// 获取文章列表
// @Summary 获取文章列表
// @Tags   article  文章管理
// @Accept  json
// @Produce  json
// @Param   page         query    int       false     "页码,默认为1"
// @Param   num          query    int       false     "返回条数,默认为10"
// @Param   sort         query    string    false     "排序字段,默认为created_at"
// @Param   key          query    string    false     "搜索关键字"
// @Param   orderType    query    string    false     "排序规则,默认为DESC"
// @Param   beginAt      query    string    false     "开始时间"
// @Param   endAt        query    string    false     "结束时间"
// @Success 200 {array}   models.Article 	"文章列表"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /articles [get]
func (ArticleController) List(c *gin.Context) {
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
		arr = append(arr, v)
		arr = append(arr, v)
		whereOrder = append(whereOrder, orm.PageWhere{Where: " title like ? or describe like ? or content like ?", Value: arr})
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
	list := make([]models.Article, 0)
	var indexPage orm.IndexPage
	indexPage.Page = reqData.Page
	indexPage.Num = reqData.Num
	err = orm.GetPageProload(&models.Article{}, &models.Article{}, &list, &indexPage, "Tag", reqData.Sort, whereOrder...)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccessPage(c, indexPage, list)
}

// 获取文章详情
// @Summary 获取文章详情
// @Tags   article  文章管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.IdReq    true     "文章ID"
// @Success 200 {object}   models.Article 	    "{code:200,msg:ok}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /article/:id [get]
func (ArticleController) Detail(c *gin.Context) {
	reqData := IdReq{}
	err := c.ShouldBindUri(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	article := models.Article{}
	_, err = orm.FirstByIDRelated(&article, reqData.Id, &article.Tag)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccess(c, article)
}

type articleCreateReq struct {
	TagID    int    `form:"tag_id" json:"tag_id" binding:"required"`             // 标签ID
	Title    string `form:"title" json:"title" binding:"required,min=2,max=100"` // 标题
	Describe string `form:"describe" json:"describe"`                            // 描述
	Content  string `form:"content" json:"content"`                              // 内容
}

// 创建文章
// @Summary 创建文章
// @Tags   article  文章管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.articleCreateReq    true     "文章信息"
// @Success 200 {object}   handler.ResponseModel 	"{code:200,msg:ok,data:{id:"id"}}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /article [post]
func (ArticleController) Create(c *gin.Context) {
	reqData := articleCreateReq{}
	err := c.ShouldBind(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	tag := models.Tag{}
	notFound, err := orm.FirstByID(&tag, uint(reqData.TagID))
	if err != nil {
		if notFound {
			resBadRequest(c, "tag_id: "+err.Error())
			return
		}
		resBusinessP(c, err.Error())
		return
	}
	article := models.Article{}
	article.TagID = reqData.TagID
	article.Title = reqData.Title
	article.Describe = reqData.Describe
	article.Content = reqData.Content
	err = orm.Create(&article)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccess(c, gin.H{"id": article.ID})
}

type articleUpdateReq struct {
	TagID    int    `form:"tag_id" json:"tag_id"`                       // 标签ID
	Title    string `form:"title" json:"title"`                         // 标题
	Describe string `form:"describe" json:"describe"`                   // 描述
	Content  string `form:"content" json:"content"`                     // 内容
}

// 修改文章
// @Summary 修改文章
// @Tags   article  文章管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.IdReq    true     "文章ID"
// @Param   body    body    handler.articleUpdateReq    true     "文章信息"
// @Success 200 {object}   handler.ResponseModel 	"{code:200,msg:ok}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /article/:id [put]
func (ArticleController) Update(c *gin.Context) {
	reqUriData := IdReq{}
	reqData := articleUpdateReq{}
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
	tag := models.Tag{}
	notFound, err := orm.FirstByID(&tag, uint(reqData.TagID))
	if err != nil {
		if notFound {
			resBadRequest(c, "tag_id: "+err.Error())
			return
		}
		resBusinessP(c, err.Error())
		return
	}
	where := models.Article{}
	where.ID = reqUriData.Id
	err = orm.Updates(&where, &reqData)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccessMsg(c)
}

// 删除文章
// @Summary 删除文章
// @Tags   article  文章管理
// @Accept  json
// @Produce  json
// @Param   body    body    handler.IdReq    true     "文章ID"
// @Success 200 {object}   handler.ResponseModel 	"{code:200,msg:ok}"
// @Failure 400  {object} handler.ResponseModel "{code:400,msg:无效的请求参数}"
// @Failure 500 {object} handler.ResponseModel  "{code:500,msg:服务器故障}"
// @Security MustToken
// @Router /article/:id [delete]
func (ArticleController) Delete(c *gin.Context) {
	reqData := IdReq{}
	err := c.ShouldBindUri(&reqData)
	if err != nil {
		resBadRequest(c, err.Error())
		return
	}
	_, err = orm.DeleteByID(models.Article{}, reqData.Id)
	if err != nil {
		log.Println(err)
		resBusinessP(c, err.Error())
		return
	}
	resSuccessMsg(c)
}
