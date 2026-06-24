package projects

import (
	"errors"

	"github.com/yourusername/track-the-trails/config"
	"gorm.io/gorm"
)

func CreateProject(project *Project) error {

	return config.DB.Create(project).Error
}

func UpdateProject(project *Project) error {
	return config.DB.Save(project).Error
}

func GetProjectsByOwner(ownerID string) ([]Project, error) {

	var projects []Project

	err := config.DB.
		Where("owner_id = ?", ownerID).
		Order("created_at DESC").
		Find(&projects).Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func GetProjectByID(projectID string) (*Project, error) {

	var project Project

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
