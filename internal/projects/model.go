package projects

import (
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/track-the-trails/config"
)

type Project struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	OwnerID     uuid.UUID `gorm:"column:owner_id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	Status      string    `gorm:"column:status"`
	IsArchived  bool      `gorm:"column:is_archived"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Project) TableName() string {
	return "projects"
}

type ProjectMember struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`
	ProjectID uuid.UUID `gorm:"column:project_id"`
	UserID    uuid.UUID `gorm:"column:user_id"`
	Role      string    `gorm:"column:role"`
	JoinedAt  time.Time `gorm:"column:joined_at"`
}

func (ProjectMember) TableName() string {
	return "project_members"
}

func AddMember(member *ProjectMember) error {

	return config.DB.Create(member).Error
}
