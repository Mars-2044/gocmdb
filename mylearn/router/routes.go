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

	r.POST("/signup", controller.SignUpHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}