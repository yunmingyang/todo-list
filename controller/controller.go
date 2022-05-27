package controller

import (
	"net/http"
	"todo-list/dao"
	"todo-list/models"

	"github.com/gin-gonic/gin"
)



func ListOptions(ctx *gin.Context) {
	todos, err := dao.ListOptions()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10003,
			"msg": err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 00002,
		"msg": "successful",
		"data": todos,
	})
}

func AddOption(ctx *gin.Context) {
	var todo models.Todo
	// If content-type is not correct, ShouldBind will not bind data to struct correctly
	// Also BindJSON will write 400 in the response.
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10001,
			"msg": err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	if err := dao.AddOption(&todo); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10002,
			"msg": err.Error(),
			"data": todo,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 00001,
		"msg": "successful",
		"data": todo,
	})
}

func DeleteOption(ctx *gin.Context) {
	var params map[string]int

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10004,
			"msg": err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	id, ok := params["id"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 10005,
			"msg": "no ID supported",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	if err := dao.DeleteOption(id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10006,
			"msg": err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 00003,
		"msg": "successful",
		"data": params,
	})
}

func UpdateOptionState(ctx *gin.Context) {
	var params map[string]interface{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10007,
			"msg": err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	status, ok := params["status"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 10008,
			"msg": "status must be supportted when updating",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	id, ok := params["id"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 10009,
			"msg": "id must be supportted when updating",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	if id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 10010,
			"msg": "id must be not 0",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	old, err := dao.GetOption(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10011,
			"msg": "get option by id failed: " + err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	if old.Status == status {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10012,
			"msg": "no change as status is the same",
			"data": old,
		})
		return
	}

	if err := dao.UpdateOptionState(id, status); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 10013,
			"msg": err.Error(),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 00004,
		"msg": "successful",
		"data": params,
	})
}