package controllers

import  (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mylearn/logic"
	"mylearn/models"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验

	var p models.ParamSignUP
	if err := c.ShouldBindJSON(&p); err != nil {

		// 判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
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
