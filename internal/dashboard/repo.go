package dashboard

import "github.com/yourusername/track-the-trails/config"

func GetTotalProjects(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("project_members").
		Where("user_id = ?", userID).
		Count(&count).Error

	return count, err
}

func GetActiveProjects(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("projects p").
		Joins("JOIN project_members pm ON pm.project_id = p.id").
		Where("pm.user_id = ? AND p.is_archived = false", userID).
		Count(&count).Error

	return count, err
}

func GetArchivedProjects(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("projects p").
		Joins("JOIN project_members pm ON pm.project_id = p.id").
		Where("pm.user_id = ? AND p.is_archived = true", userID).
		Count(&count).Error

	return count, err
}

func GetOwnedProjects(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("projects").
		Where("owner_id = ?", userID).
		Count(&count).Error

	return count, err
}

func GetAssignedTasksCount(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks").
		Where("assigned_to = ?", userID).
		Count(&count).Error

	return count, err
}

func GetPendingTasksCount(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks").
		Where("assigned_to = ? AND status = ?", userID, "pending").
		Count(&count).Error

	return count, err
}

func GetInProgressTasksCount(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks").
		Where("assigned_to = ? AND status = ?", userID, "in_progress").
		Count(&count).Error

	return count, err
}

func GetCompletedTasksCount(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks").
		Where("assigned_to = ? AND status = ?", userID, "completed").
		Count(&count).Error

	return count, err
}

func GetTotalTeamMembers(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("project_members pm").
		Joins("JOIN projects p ON p.id = pm.project_id").
		Where("p.owner_id = ?", userID).
		Count(&count).Error

	return count, err
}

func GetTotalProjectTasks(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks t").
		Joins("JOIN projects p ON p.id = t.project_id").
		Where("p.owner_id = ?", userID).
		Count(&count).Error

	return count, err
}

func GetTasksCompletedToday(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks").
		Where(`
			assigned_to = ?
			AND status = ?
			AND DATE(completed_at) = CURRENT_DATE
		`, userID, "completed").
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetTasksCreatedToday(userID string) (int64, error) {

	var count int64

	err := config.DB.
		Table("tasks").
		Where(`
			assigned_to = ?
			AND DATE(created_at) = CURRENT_DATE
		`, userID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
