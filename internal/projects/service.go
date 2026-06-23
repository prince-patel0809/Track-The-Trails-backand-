package projects

import (
	"errors"

	"github.com/google/uuid"
)

func CreateProjectService(
	userID string,
	req CreateProjectRequest,
) (*Project, error) {

	ownerID, err := uuid.Parse(userID)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	project := Project{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Title:       req.Title,
		Description: req.Description,
		Status:      "active",
	}

	err = CreateProject(&project)

	if err != nil {
		return nil, errors.New("failed to create project")
	}

	return &project, nil
}

func UpdateProjectService(
	projectID string,
	userID string,
	req UpdateProjectRequest,
) error {

	project, err := GetProjectByID(projectID)

	if err != nil {
		return errors.New("failed to fetch project")
	}

	if project == nil {
		return errors.New("project not found")
	}

	if project.OwnerID.String() != userID {
		return errors.New("unauthorized")
	}

	if req.Title != "" {
		project.Title = req.Title
	}

	if req.Description != "" {
		project.Description = req.Description
	}

	if req.Status != "" {
		project.Status = req.Status
	}

	err = UpdateProject(project)

	if err != nil {
		return errors.New("failed to update project")
	}

	return nil
}
