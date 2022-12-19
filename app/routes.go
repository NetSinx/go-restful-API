package app

import (
	prodController "github.com/NetSinx/go-restful-API/controller/product"
	userController "github.com/NetSinx/go-restful-API/controller/user"
	"github.com/NetSinx/go-restful-API/middleware"

	// "github.com/NetSinx/go-restful-API/middleware"
	prodRepo "github.com/NetSinx/go-restful-API/repository/product"
	userRepo "github.com/NetSinx/go-restful-API/repository/user"
	prodService "github.com/NetSinx/go-restful-API/service/product"
	userService "github.com/NetSinx/go-restful-API/service/user"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	createProductRepository := prodRepo.NewProductRepository(DB)
	createProductService := prodService.NewProductService(createProductRepository)
	createProductController := prodController.NewProductController(createProductService)

	createUserRepository := userRepo.NewUserRepository(DB)
	createUserService := userService.NewUserService(createUserRepository)
	createUserController := userController.NewUserController(createUserService)
	
	router := gin.New()
	
	authGroup := router.Group("/auth")
	authGroup.POST("/login", createUserController.Login)
	authGroup.POST("/register", createUserController.Register)

	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware())
	apiGroup.POST("/products", createProductController.Create)
	apiGroup.PUT("/product/:productId", createProductController.Update)
	apiGroup.DELETE("/product/:productId", createProductController.Delete)
	apiGroup.GET("/product/:productId", createProductController.FindById)
	apiGroup.GET("/products", createProductController.FindAll)
	apiGroup.DELETE("/user/:nameUser", createUserController.Delete)

	return router
}