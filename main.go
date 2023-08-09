package main

import (
	"zal-api/src/config"
	route "zal-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.DB()


	route.Api(r, db)
	
	r.Run()
}