package tasks

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/yourusername/track-the-trails/internal/projects"
)

func AssignTaskService(
	projectID string,
	ownerID string,
	req AssignTaskRequest,
) (*Task, error) {

	project, err := GetProjectByID(projectID)
	if err != nil {
		return nil, errors.New("failed to fetch project")
	}

	if project == nil {
		return nil, errors.New("project not found")
	}

	// Only owner can assign task
	if project.OwnerID.String() != ownerID {
		return nil, errors.New("only owner can assign task")
	}

	// Check member exists in project
	isMember, err := IsProjectMember(projectID, req.AssignedTo)

	if err != nil {
		return nil, errors.New("failed to verify member")
	}

	if !isMember {
		return nil, errors.New("user is not a project member")
	}

	projectUUID, _ := uuid.Parse(projectID)
	ownerUUID, _ := uuid.Parse(ownerID)
	memberUUID, _ := uuid.Parse(req.AssignedTo)

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)

	if err != nil {
		return nil, errors.New("invalid due date format")
	}

	task := Task{
		ID:          uuid.New(),
		ProjectID:   projectUUID,
		AssignedBy:  ownerUUID,
		AssignedTo:  memberUUID,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      "pending",
		DueDate:     dueDate,
	}

	err = CreateTask(&task)

	if err != nil {
		return nil, errors.New("failed to assign task")
	}

	return &task, nil
}

func GetMyTasksService(
	projectID string,
	userID string,
) ([]TaskResponse, error) {

	project, err := projects.GetProjectByID(projectID)

	if err != nil {
		return nil, errors.New("failed to fetch project")
	}

	if project == nil {
		return nil, errors.New("project not found")
	}

	isMember, err := IsProjectMember(projectID, userID)

	if err != nil {
		return nil, errors.New("failed to verify member")
	}

	if !isMember && project.OwnerID.String() != userID {
		return nil, errors.New("unauthorized")
	}

	return GetMyTasks(projectID, userID)
}

func GetProjectTasksService(
	projectID string,
	userID string,
) ([]TaskResponse, error) {

	project, err := projects.GetProjectByID(projectID)

	if err != nil {
		return nil, errors.New("failed to fetch project")
	}

	if project == nil {
		return nil, errors.New("project not found")
	}

	if project.OwnerID.String() == userID {
		return GetAllProjectTasks(projectID)
	}

	// Check member
	isMember, err := IsProjectMember(projectID, userID)

	if err != nil {
		return nil, errors.New("failed to verify member")
	}

	if !isMember {
		return nil, errors.New("unauthorized")
	}

	// Member sees only their own tasks
	return GetMyTasks(projectID, userID)
}

func GetMemberTasksService(
	projectID string,
	memberID string,
	ownerID string,
) ([]TaskResponse, error) {

	project, err := projects.GetProjectByID(projectID)

	if err != nil {
		return nil, errors.New("failed to fetch project")
	}

	if project == nil {
		return nil, errors.New("project not found")
	}

	// Only owner can view another member's tasks
	if project.OwnerID.String() != ownerID {
		return nil, errors.New("only owner can access this")
	}

	isMember, err := IsProjectMember(projectID, memberID)

	if err != nil {
		return nil, errors.New("failed to verify member")
	}

	if !isMember {
		return nil, errors.New("user is not a member")
	}

	return GetMemberTasks(projectID, memberID)
}
