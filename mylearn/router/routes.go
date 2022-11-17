package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine{
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}