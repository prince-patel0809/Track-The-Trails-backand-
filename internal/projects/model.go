package projects

import (
	"time"

	"github.com/google/uuid"
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
