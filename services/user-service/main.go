package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilismadiun/tembusptn/services/user-service/commons/database"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Infrastructures/driver"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	driver.Router(router, database.DB)

	router.Run(":3000")
}
