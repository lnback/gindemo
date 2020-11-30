package main

import (
	"gindemo/controller"
	"gindemo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r * gin.Engine) *gin.Engine  {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)

	categoryRoutes := r.Group("/categories")

	categoryController := controller.NewCategoryController()

	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.GET("/:id",categoryController.Show)
	categoryRoutes.DELETE("/:id",categoryController.Delete)
	categoryRoutes.PUT("/:id",categoryController.Update)


	return r
}