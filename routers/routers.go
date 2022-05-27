package routers

import (
	"todo-list/controller"

	"github.com/gin-gonic/gin"
)



func Register(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/add", controller.AddOption)

		v1.GET("/list", controller.ListOptions)

		v1.POST("delete", controller.DeleteOption)

		v1.POST("/update", controller.UpdateOptionState)
	}
}