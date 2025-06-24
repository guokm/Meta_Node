package main

import (
	"fmt"
	"github/metanode/go_blog/config"
	"github/metanode/go_blog/middleware"
	"github/metanode/go_blog/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	config.InitLog()
	config.InitMysql()

	r := gin.Default()
	r.Use(middleware.GlobalException(), config.ReqLog())

	routers.InitRouter(r)

	// Start the server on port 8080
	if err := r.Run(":" + config.GlobalConfig.Port); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
