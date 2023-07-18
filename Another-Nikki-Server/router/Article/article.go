package Article

import (
	"Another-Nikki/dal"
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var article dal.Article

func PostArticle(c *gin.Context) {
	var newArticle model.Article
	err := c.BindJSON(&newArticle)
	if util.HandleError(c, err, "参数填写错误") {
		return
	}

	if newArticle.Content == "" {
		util.SendResp(c, 404, nil, "不可以上传空的留言哦~")
		return
	}
	if strings.TrimSpace(newArticle.Content) == "" {
		util.SendResp(c, 404, nil, "全都是空格不行的呢～")
		return
	}
	err = article.Create(newArticle.Title, newArticle.Content, newArticle.Description)
	if util.HandleError(c, err, "数据库创建记录失败") {
		return
	}
	util.SendResp(c, 200, nil, "success")
}

func GetArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("ID"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := article.GetInfoById(id)
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func GetArticleByPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if util.HandleError(c, err, "参数错误") {
		return
	}
	ret, err := article.GetList(page)
	if util.HandleError(c, err, "数据库查询记录失败") {
		return
	}
	util.SendResp(c, 200, ret, "success")
}

func UpdateArticle(c *gin.Context) {
	var Article model.ArticleUpdate
	err := c.BindJSON(&Article)
	if util.HandleError(c, err, "参数填写错误") {
		return
	}
	if Article.Content == "" {
		util.SendResp(c, 404, nil, "不可以上传空的留言哦~")
		return
	}
	if strings.TrimSpace(Article.Content) == "" {
		util.SendResp(c, 404, nil, "全都是空格不行的呢～")
		return
	}
	err = article.UpdateArticle(Article.ID, Article.Title, Article.Description, Article.Content)
	if util.HandleError(c, err, "数据库更新记录失败") {
		return
	}
	util.SendResp(c, 200, nil, "更新成功")
}
