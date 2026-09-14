package driver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	infrastructures "github.com/wilismadiun/tembusptn/services/user-service/src/Infrastructures"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Interfaces/http"
	"gorm.io/gorm"
)

func Router(router *gin.Engine, db *gorm.DB) {
	fmt.Println("router 1")
	handler := infrastructures.Container(db)

	fmt.Println("router 2")
	// router
	http.UserRouter(router, handler)
}
