package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `gorm:"column:id;type:uuid;primaryKey"`
	ProjectID   uuid.UUID  `gorm:"column:project_id"`
	AssignedBy  uuid.UUID  `gorm:"column:assigned_by"`
	AssignedTo  uuid.UUID  `gorm:"column:assigned_to"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Priority    string     `gorm:"column:priority"`
	Status      string     `gorm:"column:status"`
	DueDate     time.Time  `gorm:"column:due_date"`
	CompletedAt *time.Time `gorm:"column:completed_at"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}
