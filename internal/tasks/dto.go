package tasks

type AssignTaskRequest struct {
	AssignedTo  string `json:"assigned_to" binding:"required,uuid"`
	Title       string `json:"title" binding:"required,min=3,max=255"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date" binding:"required"`
}
