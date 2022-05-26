package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



var (
	DB *gorm.DB
	SQLLogMode logger.LogLevel = logger.Silent
)

type Todo struct {
	ID int `json:"id"`
	Content string `json:"content"`
	Status bool `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:1qaz9ol.@(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(SQLLogMode),
	})

	return err
}

func main() {
	if err := initMySQL(); err != nil {
		log.Fatalf("database connect failed: %v", err)
	}

	DB.AutoMigrate(&Todo{})

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "test",
			})
		})

		v1.POST("/add", func(ctx *gin.Context) {
			var todo Todo
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

			if err := DB.Create(&todo).Error; err != nil {
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
		})

		v1.GET("/list", func(ctx *gin.Context) {
			var todos []Todo

			if err := DB.Find(&todos).Error; err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"code": 10003,
					"msg": err.Error(),
					"data": nil,
				})
				ctx.Abort()
				return
			}

			// ctx.JSON(http.StatusOK, gin.H{
			// 	"code": 00002,
			// 	"msg": "successful",
			// 	"data": todos,
			// })
			ctx.JSON(http.StatusOK, todos)
		})

		v1.POST("delete", func(ctx *gin.Context) {
			var d map[string]interface{}

			if err := ctx.ShouldBindJSON(&d); err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"code": 10004,
					"msg": err.Error(),
					"data": nil,
				})
				ctx.Abort()
				return
			}

			id, ok := d["id"]
			if !ok {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 10005,
					"msg": "no id to delete",
					"data": nil,
				})
				ctx.Abort()
				return
			}

			if err := DB.Model(&Todo{}).Delete("id = ?", id).Error; err != nil {
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
				"data": id,
			})
		})

		v1.POST("/update", func(ctx *gin.Context) {
			var new Todo
			var old Todo

			if err := ctx.ShouldBindJSON(&new); err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"code": 10007,
					"msg": err.Error(),
					"data": nil,
				})
				ctx.Abort()
				return
			}

			if new.ID == 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 10008,
					"msg": "id must be supportted when updating",
					"data": nil,
				})
				ctx.Abort()
				return
			}

			if err := DB.Find(&old, new.ID).Error; err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"code": 10009,
					"msg": err.Error(),
					"data": nil,
				})
				ctx.Abort()
				return
			}

			if old.Status == new.Status && old.Content == new.Content {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 00004,
					"msg": "successful",
					"data": "no new change",
				})
				return
			}

			new.Content = old.Content
			if err := DB.Save(&new).Error; err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"code": 10010,
					"msg": err.Error(),
					"data": nil,
				})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"code": 00004,
				"msg": "successful",
				"data": new,
			})
		})
	}

	r.Run(":9091")
}