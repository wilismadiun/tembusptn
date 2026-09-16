package main

import (
	"subscription-service/commons/database"
	"subscription-service/src/Infrastructures/driver"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	driver.Router(router, database.DB)

	router.Run(":3001")
}
