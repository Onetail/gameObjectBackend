package model

import (
	"time"
)

type userGenderEnum struct {
	Male  string `enum:"MALE"`
    Female string `enum:"FEMALE"`
	Other string `enum:"OTHER"`
}


type User struct {
	BaseModel
	Nickname string `gorm:"type:varchar(128);not null"`
	Birthday time.Time `gorm:"index:idx_time"`
	PhoneCountryCode    string `gorm:"column:phoneCountryCode;type:varchar(16)"`
	PhoneNumber    string   `gorm:"column:phoneNumber;type:varchar(32)"`
	Gender    userGenderEnum `gorm:"column:gender;type:varchar(16)"`
	Region    string `gorm:"type:varchar(32)"`
}