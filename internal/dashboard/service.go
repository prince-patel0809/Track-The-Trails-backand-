package dashboard

func GetDashboard(userID string) (*DashboardResponse, error) {

	// Project Statistics
	totalProjects, _ := GetTotalProjects(userID)
	activeProjects, _ := GetActiveProjects(userID)
	archivedProjects, _ := GetArchivedProjects(userID)
	ownedProjects, _ := GetOwnedProjects(userID)

	// Task Statistics
	assignedTasks, _ := GetAssignedTasksCount(userID)
	pendingTasks, _ := GetPendingTasksCount(userID)
	inProgressTasks, _ := GetInProgressTasksCount(userID)
	completedTasks, _ := GetCompletedTasksCount(userID)

	// Owner Statistics
	totalMembers, _ := GetTotalTeamMembers(userID)
	totalProjectTasks, _ := GetTotalProjectTasks(userID)

	// Activity Statistics
	todayCompleted, _ := GetTasksCompletedToday(userID)

	return &DashboardResponse{
		Projects: ProjectStats{
			Total:    totalProjects,
			Active:   activeProjects,
			Archived: archivedProjects,
			Owned:    ownedProjects,
		},
		Tasks: TaskStats{
			AssignedToMe: assignedTasks,
			Pending:      pendingTasks,
			InProgress:   inProgressTasks,
			Completed:    completedTasks,
		},
		Owner: OwnerStats{
			OwnedProjects:     ownedProjects,
			TotalTeamMembers:  totalMembers,
			TotalProjectTasks: totalProjectTasks,
		},
		Activity: ActivityStats{
			TasksCompletedToday: todayCompleted,
		},
	}, nil
}
