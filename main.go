package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
