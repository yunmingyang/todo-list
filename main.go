package main

import (
	"log"
	"todo-list/routers"
	"todo-list/dao"

	"github.com/gin-gonic/gin"
)



func main() {
	if err := dao.InitMySQL(); err != nil {
		log.Fatalf("database connect failed: %v", err)
	}
	if err := dao.AutoMigrate(); err != nil {
		log.Fatalf("gorm automigrate for gorm failed: %v", err)
	}

	r := gin.Default()
	routers.Register(r)

	r.Run(":9091")
}