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

func GetProjectByID(id string) (*Project, error) {

	var project Project

	err := config.DB.
		Where("id = ?", id).
		First(&project).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &project, nil
}
