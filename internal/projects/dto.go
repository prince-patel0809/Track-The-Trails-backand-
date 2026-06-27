package projects

import "time"

type CreateProjectRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=255"`
	Description string `json:"description"`
	IsArchived  bool   `json:"is_archived"`
}

type UpdateProjectRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AddMemberRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
}

type MemberResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Bio      string    `json:"bio"`
	Theme    string    `json:"theme"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}
