package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {	
	c.HTML(http.StatusOK, "page/index", gin.H{
		"title": "Go GIN MVC Starter ⚡",
	})
}

func CreditPage(c *gin.Context) {	
	c.HTML(http.StatusOK, "page/credits", gin.H{
		"title": "Credits",
	})
}

func StructurePage(c *gin.Context) {	
	c.HTML(http.StatusOK, "page/structure", gin.H{
		"title": "Structure",
	})
}

func ApiPage(c *gin.Context) {	
	c.HTML(http.StatusOK, "page/api", gin.H{
		"title": "API",
	})
}
