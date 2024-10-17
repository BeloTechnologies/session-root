package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"session-root/routes"
	"session-root/utils"
)

func main() {
	utils.InitConfig()

	serverPort := viper.GetInt("server.port")

	r := gin.Default()

	routes.RootRoutes(r)

	e := r.Run(fmt.Sprintf(":%d", serverPort))
	if e != nil {
		return
	}
}
