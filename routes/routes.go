package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/track-the-trails/internal/auth"
	"github.com/yourusername/track-the-trails/internal/projects"
	"github.com/yourusername/track-the-trails/internal/tasks"
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
		projectGroup.PUT("/update/:id", projects.UpdateProjectHandler)
		projectGroup.GET("/GetAllProjects", projects.GetMyProjectsHandler)
		projectGroup.GET("/GetProject/:id", projects.GetProjectHandler)
		projectGroup.POST("/:id/members", middleware.JWTAuthMiddleware(), projects.AddMemberHandler)
		projectGroup.GET("/:id/GetAllmembers", middleware.JWTAuthMiddleware(), projects.GetProjectMembersHandler)
	}

	userGroup := api.Group("/users")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.GET("/GetAll", projects.GetAllUsersHandler)
	}

	TaskGroup := api.Group("/tasks")
	TaskGroup.Use(middleware.JWTAuthMiddleware())
	{
		TaskGroup.POST("/:id/assign", tasks.AssignTaskHandler)
		TaskGroup.GET("/:id/Gettasksbyuser", middleware.JWTAuthMiddleware(), tasks.GetProjectTasksHandler)
		TaskGroup.GET("/:id/members/:userId/tasks", middleware.JWTAuthMiddleware(), tasks.GetMemberTasksHandler)
	}
}
