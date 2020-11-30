package main

import (
	"gindemo/controller"
	"gindemo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r * gin.Engine) *gin.Engine  {
	r.Use(middleware.CORSMiddleware(),middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)

	categoryRoutes := r.Group("/categories")

	categoryController := controller.NewCategoryController()

	categoryRoutes.POST("",categoryController.Create)
	categoryRoutes.GET("/:id",categoryController.Show)
	categoryRoutes.DELETE("/:id",categoryController.Delete)
	categoryRoutes.PUT("/:id",categoryController.Update)

	postRoutes := r.Group("/posts")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()

	postRoutes.POST("",postController.Create)
	postRoutes.GET("/:id",postController.Show)
	postRoutes.DELETE("/:id",postController.Delete)
	postRoutes.PUT("/:id",postController.Update)
	postRoutes.POST("/page/list",postController.PageList)
	return r
}