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

type RowsAffectedModel struct {
	Data int `json:"data" example:"1" format:"int64"`
}

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不進行解析
	if len(data) == 2 {
		*t = DateTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	*t = DateTime(now)
	return
}
