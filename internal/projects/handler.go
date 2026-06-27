package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProjectHandler(c *gin.Context) {

	var req CreateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	userID := c.MustGet("userID").(string)

	project, err := CreateProjectService(userID, req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Project created successfully",
		"project": gin.H{
			"id":          project.ID,
			"title":       project.Title,
			"description": project.Description,
			"status":      project.Status,
			"owner_id":    project.OwnerID,
			"created_at":  project.CreatedAt,
		},
	})
}

func UpdateProjectHandler(c *gin.Context) {

	projectID := c.Param("id")

	userID := c.MustGet("userID").(string)

	var req UpdateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})

		return
	}

	err := UpdateProjectService(
		projectID,
		userID,
		req,
	)

	if err != nil {

		switch err.Error() {

		case "project not found":

			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Project not found",
			})

		case "unauthorized":

			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Only owner can update project",
			})

		default:

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Project updated successfully",
	})
}

func GetMyProjectsHandler(c *gin.Context) {

	userID := c.MustGet("userID").(string)

	projects, err := GetMyProjects(userID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"count":    len(projects),
		"projects": projects,
	})
}

func GetProjectHandler(c *gin.Context) {

	projectID := c.Param("id")

	userID := c.MustGet("userID").(string)

	project, err := GetProjectService(projectID, userID)

	if err != nil {

		switch err.Error() {

		case "project not found":
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Project not found",
			})

		case "unauthorized":
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Access denied",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"project": project,
	})
}

func AddMemberHandler(c *gin.Context) {

	projectID := c.Param("id")

	ownerID := c.MustGet("userID").(string)

	var req AddMemberRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request",
		})
		return
	}

	err := AddMemberService(projectID, ownerID, req)

	if err != nil {

		switch err.Error() {

		case "project not found":
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Project not found",
			})

		case "only owner can add members":
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Only the project owner can add members",
			})

		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})

		case "user already exists in project":
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "User is already a member of this project",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Member added successfully",
	})
}

func GetAllUsersHandler(c *gin.Context) {

	users, err := GetAllUsersService()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   len(users),
		"users":   users,
	})
}

func GetProjectMembersHandler(c *gin.Context) {

	projectID := c.Param("id")
	userID := c.MustGet("userID").(string)

	members, err := GetProjectMembersService(
		projectID,
		userID,
	)

	if err != nil {

		switch err.Error() {

		case "project not found":

			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Project not found",
			})

		case "unauthorized":

			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Only the project owner can view members",
			})

		default:

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   len(members),
		"members": members,
	})
}
