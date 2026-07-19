package dashboard

type DashboardResponse struct {
	Projects ProjectStats  `json:"projects"`
	Tasks    TaskStats     `json:"tasks"`
	Owner    OwnerStats    `json:"owner"`
	Activity ActivityStats `json:"activity"`
}

type ProjectStats struct {
	Total    int64 `json:"total"`
	Active   int64 `json:"active"`
	Archived int64 `json:"archived"`
	Owned    int64 `json:"owned"`
}

type TaskStats struct {
	AssignedToMe   int64 `json:"assigned_to_me"`
	Pending        int64 `json:"pending"`
	InProgress     int64 `json:"in_progress"`
	Completed      int64 `json:"completed"`
	Cancelled      int64 `json:"cancelled"`
	Overdue        int64 `json:"overdue"`
	DueToday       int64 `json:"due_today"`
	DueThisWeek    int64 `json:"due_this_week"`
	CompletionRate int64 `json:"completion_rate"`
}

type OwnerStats struct {
	OwnedProjects         int64 `json:"owned_projects"`
	TotalTeamMembers      int64 `json:"total_team_members"`
	TotalProjectTasks     int64 `json:"total_project_tasks"`
	CompletedProjectTasks int64 `json:"completed_project_tasks"`
	PendingProjectTasks   int64 `json:"pending_project_tasks"`
}

type ActivityStats struct {
	TasksCompletedToday int64 `json:"tasks_completed_today"`
	TasksCreatedToday   int64 `json:"tasks_created_today"`
	UnreadNotifications int64 `json:"unread_notifications"`
	UnreadMessages      int64 `json:"unread_messages"`
}
