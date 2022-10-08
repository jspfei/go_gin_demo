package main

import (
	"CloudRestaurant/tool"

	"CloudRestaurant/controller"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	app := gin.Default()
	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router((router))
}
