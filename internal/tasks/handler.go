package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignTaskHandler(c *gin.Context) {

	projectID := c.Param("id")

	ownerID := c.MustGet("userID").(string)

	var req AssignTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})

		return
	}

	task, err := AssignTaskService(
		projectID,
		ownerID,
		req,
	)

	if err != nil {

		switch err.Error() {

		case "project not found":

			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Project not found",
			})

		case "only owner can assign task":

			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Only project owner can assign task",
			})

		case "user is not a project member":

			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Selected user is not a member of this project",
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
		"message": "Task assigned successfully",
		"task":    task,
	})
}

func GetMyTasksHandler(c *gin.Context) {

	projectID := c.Param("id")

	userID := c.MustGet("userID").(string)

	tasks, err := GetMyTasksService(
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
				"message": "You are not a member of this project",
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
		"count":   len(tasks),
		"tasks":   tasks,
	})
}
func GetProjectTasksHandler(c *gin.Context) {

	projectID := c.Param("id")
	userID := c.MustGet("userID").(string)

	tasks, err := GetProjectTasksService(projectID, userID)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   len(tasks),
		"tasks":   tasks,
	})
}

func GetMemberTasksHandler(c *gin.Context) {

	projectID := c.Param("id")
	memberID := c.Param("userId")
	ownerID := c.MustGet("userID").(string)

	tasks, err := GetMemberTasksService(
		projectID,
		memberID,
		ownerID,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   len(tasks),
		"tasks":   tasks,
	})
}
