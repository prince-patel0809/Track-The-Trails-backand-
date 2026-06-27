package projects

import (
	"errors"

	"github.com/yourusername/track-the-trails/config"
	"github.com/yourusername/track-the-trails/internal/auth"
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

func GetProjectById(projectID string) (*Project, error) {

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

func UserExists(userID string) (bool, error) {

	var count int64

	err := config.DB.
		Table("users").
		Where("id = ?", userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func IsAlreadyMember(projectID, userID string) (bool, error) {

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

func GetAllUsers() ([]auth.User, error) {

	var users []auth.User

	err := config.DB.
		Order("created_at DESC").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetProjectMembers(projectID string) ([]MemberResponse, error) {

	var members []MemberResponse

	err := config.DB.
		Table("project_members pm").
		Select(`
			u.id,
			u.name,
			u.email,
			u.bio,
			u.theme,
			pm.role,
			pm.joined_at
		`).
		Joins("JOIN users u ON u.id = pm.user_id").
		Where("pm.project_id = ?", projectID).
		Order("pm.joined_at ASC").
		Scan(&members).Error

	if err != nil {
		return nil, err
	}

	return members, nil
}
