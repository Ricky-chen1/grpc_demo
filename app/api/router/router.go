package router

import (
	"grpc_demo/app/api/handler"
	"grpc_demo/app/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/user/register", handler.UserRegister)
	r.POST("/user/login", handler.UserLogin)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		v1.POST("/task", handler.TaskCreate)
		v1.PUT("/task", handler.TaskUpdate)
		v1.DELETE("/task", handler.TaskDelete)
		v1.GET("/task/search", handler.TaskSearch)
		v1.GET("/task/status", handler.TaskListGet)
	}

	return r
}
