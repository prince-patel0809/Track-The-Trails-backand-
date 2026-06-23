package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/track-the-trails/internal/auth"
	"github.com/yourusername/track-the-trails/internal/projects"
	"github.com/yourusername/track-the-trails/middleware"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", auth.RegisterHandler)
			authGroup.POST("/login", auth.LoginHandler)

			authGroup.GET(
				"/profile",
				middleware.JWTAuthMiddleware(),
				auth.GetProfileHandler,
			)
		}
	}

	projectGroup := api.Group("/projects")
	projectGroup.Use(middleware.JWTAuthMiddleware())
	{
		projectGroup.POST("/create", projects.CreateProjectHandler)
		projectGroup.PUT(
			"/update/:id",
			projects.UpdateProjectHandler,
		)
	}
}
