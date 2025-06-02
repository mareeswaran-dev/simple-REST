package routes

import (
	"user-crud-api/app"
	"user-crud-api/handlers"
	"user-crud-api/repository"
	"user-crud-api/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures all the routes for the application
func SetupRouter(app *app.App) *gin.Engine {
	r := gin.Default()

	// Initialize repository and service
	userRepo := repository.NewUserRepository(app.DB)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(app, userService)

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}
