package routes

import (
	"product/config"
	"product/controller"
	"product/middleware"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {

	db := config.GetDB()

	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authController := controller.AuthController{DB: db}

	authMiddleware := middleware.Authentication().MiddlewareFunc()

	{
		authGroup.GET("/profile", authMiddleware, authController.GetProfile)
		authGroup.POST("/sign-up", authController.SignUp)
		authGroup.POST("/sign-in", middleware.Authentication().LoginHandler)
	}

	userGroup := v1.Group("/users")
	userController := controller.UserController{DB: db}
	authorizedMiddleware := middleware.Authorization

	userGroup.Use(authMiddleware, authorizedMiddleware)
	{
		userGroup.GET("", userController.FindAll)
		userGroup.GET("/:id", userController.FindOne)
		userGroup.POST("", userController.Create)
		userGroup.DELETE("/:id", userController.Delete)
	}

	productGroup := v1.Group("/products")
	productController := controller.ProductController{DB: db}

	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)

	productGroup.Use(authMiddleware, authorizedMiddleware)
	{
		productGroup.POST("", productController.Create)
		productGroup.PATCH("/:id", productController.Update)
		productGroup.DELETE("/:id", productController.Delete)
	}

	categoryGroup := v1.Group("/categorys")
	categoryController := controller.CategoryController{DB: db}

	categoryGroup.GET("", categoryController.FindAll)
	categoryGroup.GET("/:id", categoryController.FindOne)

	categoryGroup.Use(authMiddleware, authorizedMiddleware)
	{
		categoryGroup.POST("", categoryController.Create)
		categoryGroup.PATCH("/:id", categoryController.Update)
		categoryGroup.DELETE("/:id", categoryController.Delete)
	}
}
