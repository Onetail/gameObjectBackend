package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (model *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}

type BaseModel struct {
	ID        string    `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time `gorm:"column:createdAt;index:idx_time;" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt;index:idx_time;" sql:"DEFAULT:current_timestamp ON update current_timestamp"`
}
