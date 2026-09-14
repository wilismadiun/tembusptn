package driver

import (
	"github.com/gin-gonic/gin"
	infrastructures "github.com/wilismadiun/tembusptn/services/user-service/src/Infrastructures"
	"github.com/wilismadiun/tembusptn/services/user-service/src/Interfaces/http"
	"gorm.io/gorm"
)

func Router(router *gin.Engine, db *gorm.DB) {
	handler := infrastructures.Container(db)

	// router
	http.UserRouter(router, handler)
}
