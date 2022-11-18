package controllers

import  (
	"github.com/gin-gonic/gin"
	"mylearn/logic"
	"mylearn/models"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验

	var p models.ParamSignUP
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	// 2.业务处理
	logic.SignUp(&p)
	// 3.返回响应

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
