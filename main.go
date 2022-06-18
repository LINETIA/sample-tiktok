package main

import (
	"Gin/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	db := common.InitDB()
	if db == nil {
		fmt.Println("db = nil")
	}

	router := gin.Default()
	// 添加 Get 请求路由
	router = CollectRoute(router)
	router.Run()
}




