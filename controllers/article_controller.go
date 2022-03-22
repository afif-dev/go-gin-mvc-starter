package controllers

import (
	"gin-mvc/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleForm struct {
	ID 			uint 		`form:"id"`
	Title		string 		`form:"title" binding:"required"`
	Short_Desc 	*string 	`form:"short_desc"`
	Content 	*string 	`form:"content"`
	Tags 		*string 	`form:"tags"`
}

func  GetArticles(c *gin.Context) map[string]interface{} {	
	pagination := &models.Pagination{}
	pagination.Limit = 4
	page_q := c.Query("page")
	page_curr, _ := strconv.Atoi(page_q)
	pagination.Page = page_curr
	
	message := "Articles Success"

	var articles []models.Articles
	var qy_articles []models.Articles_Select
    models.DB.Model(&articles).Count(&pagination.TotalRow)

	if(int(pagination.TotalRow) > pagination.GetOffset()) {
		models.DB.Model(&articles).Order("id desc").Limit(pagination.Limit).Offset(pagination.GetOffset()).Find(&qy_articles)
	} else {
		message = "Articles is not found"
	}

	var results = map[string]interface{} {
		"message": message,
		"page": pagination.GetPage(),
		"page_prev": pagination.GetPreviosPage(),
		"page_next": pagination.GetNextPage(),
		"page_total": pagination.GetTotalPage(),
		"page_limit": pagination.GetLimit(),
		"page_offset": pagination.GetOffset(),
		"page_total_rows": pagination.TotalRow,
		"data": qy_articles,
	}
	return results
}

// get articles api
func  GetArticlesAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"results": GetArticles(c),
	})
}

// get articles page
func  GetArticlesPage(c *gin.Context) {
	results := GetArticles(c);
	
	page := []int{}
	page_curr := results["page"].(int)
	page_total := results["page_total"].(int)	
	if page_total > 0 {
		for i:=1; i <= page_total; i++ {
			page = append(page, i)
		}
	}
	if page_curr <= page_total {
		c.HTML(http.StatusOK, "articles/index", gin.H{
			"title": "Articles Page",
			"pg_page": page,
			"results": results,
			// "safeHTML": func(v string) template.HTML {
			// 	return template.HTML(v)
			// },
		})
	} else {
		c.Redirect(http.StatusFound,"/articles")
	}
}

// get article page
func  GetArticlePage(c *gin.Context) {
	id := c.Param("id")
	var article models.Articles
	result := models.DB.First(&article, id)
	if result.Error == nil {
		c.HTML(http.StatusOK, "articles/article", gin.H{
			"title": article.Title,
			"results": article,
		})
	} else {
		c.Redirect(http.StatusFound,"/")
	}
}

// get article api
func  GetArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Articles
	result := models.DB.First(&article, id)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "article success",
			"data": article,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article is not found",
			"data": nil,
		})
	}
}

// add new article
func AddArticle(c *gin.Context) {
	var form ArticleForm

	if c.ShouldBind(&form) == nil {
		if form.Title != "" {
			article := models.Articles{
				Title: form.Title,
				Short_Desc: form.Short_Desc,
				Content: form.Content,
				Tags: form.Tags,
			}
			models.DB.Create(&article)

			c.JSON(http.StatusCreated, gin.H{
				"id": article.ID,
				"message": "add article is success",
			})
		}
	}
}

// update selected article
func UpdateArticle(c *gin.Context) {
	var form ArticleForm

	if c.ShouldBind(&form) == nil {
		if form.ID > 0 {
			var article models.Articles
			result := models.DB.First(&article, form.ID)
			if(result.Error == nil) {
				if form.Title != "" {
					article.Title = form.Title
				}
				if form.Short_Desc != nil {
					article.Short_Desc = form.Short_Desc
				}
				if form.Content != nil {
					article.Content = form.Content
				}
				if form.Tags != nil {
					article.Tags = form.Tags
				}
				models.DB.Save(&article)

				c.JSON(http.StatusOK, gin.H{
					"message": "update article is success",
					"data": article,
				})
			} else {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "article is not found",
					"data": nil,
				})
			}
		}
	}
}

// delete selected article
func DeleteArticle(c *gin.Context) {
	id := c.PostForm("id")
	if id != "" {
		var article []models.Articles
		models.DB.Delete(&article, id)
		
	}
	c.JSON(http.StatusOK, gin.H {
		"id": id,
		"status_code": http.StatusOK,
		"message": "delete article is success",
	})	
}
