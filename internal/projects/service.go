package projects

import (
	"errors"

	"github.com/google/uuid"
	"github.com/yourusername/track-the-trails/internal/auth"
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
		IsArchived:  req.IsArchived,
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

func GetMyProjects(userID string) ([]Project, error) {

	projects, err := GetProjectsByOwner(userID)

	if err != nil {
		return nil, errors.New("failed to fetch projects")
	}

	return projects, nil
}

func GetProjectService(projectID string, userID string) (*Project, error) {

	project, err := GetProjectByID(projectID)

	if err != nil {
		return nil, errors.New("failed to fetch project")
	}

	if project == nil {
		return nil, errors.New("project not found")
	}

	if project.OwnerID.String() != userID {
		return nil, errors.New("unauthorized")
	}

	return project, nil
}

func AddMemberService(projectID string, ownerID string, req AddMemberRequest) error {

	project, err := GetProjectById(projectID)

	if err != nil {
		return errors.New("failed to fetch project")
	}

	if project == nil {
		return errors.New("project not found")
	}

	// Only Owner
	if project.OwnerID.String() != ownerID {
		return errors.New("only owner can add members")
	}

	exists, err := UserExists(req.UserID)

	if err != nil {
		return errors.New("failed to verify user")
	}

	if !exists {
		return errors.New("user not found")
	}

	memberExists, err := IsAlreadyMember(projectID, req.UserID)

	if err != nil {
		return errors.New("failed to verify member")
	}

	if memberExists {
		return errors.New("user already exists in project")
	}

	projectUUID, _ := uuid.Parse(projectID)
	userUUID, _ := uuid.Parse(req.UserID)

	member := ProjectMember{
		ID:        uuid.New(),
		ProjectID: projectUUID,
		UserID:    userUUID,
		Role:      "member",
	}

	return AddMember(&member)
}

func GetAllUsersService() ([]auth.User, error) {

	users, err := GetAllUsers()

	if err != nil {
		return nil, errors.New("failed to fetch users")
	}

	return users, nil
}
