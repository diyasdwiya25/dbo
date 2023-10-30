package main

import (
	"dbo/config"
	"dbo/handler"
	"dbo/helper"
	"dbo/repository"
	"dbo/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	db := config.InitDB()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productsRepository := repository.NewProductsRepository(db)
	productsService := service.NewProductsService(productsRepository)
	productsHandler := handler.NewProductsHandler(productsService)

	customersRepository := repository.NewCustomersRepository(db)
	customersService := service.NewCustomersService(customersRepository)
	customersHandler := handler.NewCustomersHandler(customersService)

	usersRepository := repository.NewUsersRepository(db)
	usersService := service.NewUsersService(usersRepository)
	authService := service.NewAuthService()
	usersHandler := handler.NewUsersHandler(usersService, authService)

	orderDetailRepository := repository.NewOrderDetailRepository(db)
	orderDetailService := service.NewOrderDetailService(orderDetailRepository, productsRepository)

	ordersRepository := repository.NewOrdersRepository(db)
	ordersService := service.NewOrdersService(ordersRepository, productsRepository)
	ordersHandler := handler.NewOrdersHandler(ordersService, orderDetailService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/login", usersHandler.Login)

	api.Use(authMiddleware(authService))
	api.GET("/category", categoryHandler.Show)
	api.POST("/category/create", categoryHandler.Create)
	api.POST("/category/edit/:id", categoryHandler.Edit)
	api.GET("/category/find/:id", categoryHandler.FindById)
	api.GET("/category/delete/:id", categoryHandler.Delete)

	api.GET("/products", productsHandler.Show)
	api.POST("/products/create", productsHandler.Create)
	api.POST("/products/edit/:id", productsHandler.Edit)
	api.GET("/products/find/:id", productsHandler.FindById)
	api.GET("/products/delete/:id", productsHandler.Delete)

	api.GET("/customers", customersHandler.Show)
	api.POST("/customers/create", customersHandler.Create)
	api.POST("/customers/edit/:id", customersHandler.Edit)
	api.GET("/customers/find/:id", customersHandler.FindById)
	api.GET("/customers/delete/:id", customersHandler.Delete)

	api.GET("/users", usersHandler.Show)
	api.POST("/users/create", usersHandler.Create)
	api.POST("/users/edit/:id", usersHandler.Edit)
	api.GET("/users/find/:id", usersHandler.FindById)
	api.GET("/users/delete/:id", usersHandler.Delete)

	api.GET("/orders", ordersHandler.Show)
	api.POST("/orders/create", ordersHandler.Create)
	api.POST("/orders/edit/:id", ordersHandler.Edit)
	api.GET("/orders/find/:id", ordersHandler.FindById)
	api.GET("/orders/delete/:id", ordersHandler.Delete)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func authMiddleware(authService service.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		uid := int(claim["user_id"].(float64))
		name := claim["name"]
		email := claim["email"]

		c.Set("uid", uid)
		c.Set("name", name)
		c.Set("email", email)
	}
}
