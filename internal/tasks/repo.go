package tasks

import (
	"errors"

	"github.com/yourusername/track-the-trails/config"
	"github.com/yourusername/track-the-trails/internal/projects"
	"gorm.io/gorm"
)

func CreateTask(task *Task) error {

	return config.DB.Create(task).Error

}

func GetProjectByID(projectID string) (*projects.Project, error) {

	var project projects.Project

	err := config.DB.
		Where("id = ?", projectID).
		First(&project).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &project, nil
}

func IsProjectMember(projectID string, userID string) (bool, error) {

	var count int64

	err := config.DB.
		Table("project_members").
		Where("project_id = ? AND user_id = ?", projectID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetMyTasks(projectID string, userID string) ([]Task, error) {

	var tasks []Task

	err := config.DB.
		Where("project_id = ? AND assigned_to = ?", projectID, userID).
		Order("created_at DESC").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
func GetAllProjectTasks(projectID string) ([]Task, error) {

	var tasks []Task

	err := config.DB.
		Where("project_id = ?", projectID).
		Order("created_at DESC").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
