package router

import (
	"github.com/gin-gonic/gin"
	"github.com/peterouob/devops_golang/services"
)

func HandleRouter(r *gin.Engine) {
	r.GET("/", services.GetAll)
	r.GET("/:id", services.GetByID)
	r.POST("/create", services.Create)
	r.POST("/update/:id", services.Update)
	r.POST("/delete/:id", services.Delete)
}
