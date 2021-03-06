package model

import "encoding/json"

type User struct {
	BaseModel
	Nickname         string   `gorm:"type:varchar(128);not null"`
	Birthday         DateTime `gorm:"index:idx_time"`
	PhoneCountryCode string   `gorm:"column:phoneCountryCode;type:varchar(16)"`
	PhoneNumber      string   `gorm:"column:phoneNumber;type:varchar(32)"`
	Gender           string   `gorm:"column:gender;type:varchar(16)"`
	Region           string   `gorm:"type:varchar(32)"`
}

type UserResponseObject struct {
	Data User
}

type UserListResponseObject struct {
	Data []User
}

type UserSignInResponseObject struct {
	Data struct {
		token string
	}
}

type JwtUserBody struct {
	Email            string   `json:"email" binding:"required"`
	PhoneCountryCode string   `json:"phoneCountryCode" binding:"required"`
	PhoneNumber      string   `json:"phoneNumber" binding:"required"`
	Nickname         string   `json:"nickname" binding:"required"`
	Birthday         DateTime `json:"birthday" binding:"required"`
	Gender           string   `json:"gender" binding:"required"`
	Region           string   `json:"region,omitempty"`
}

type CreateUserBody struct {
	Email            string   `json:"email" binding:"required"`
	Password         string   `json:"password" binding:"required"`
	PhoneCountryCode string   `json:"phoneCountryCode" binding:"required"`
	PhoneNumber      string   `json:"phoneNumber" binding:"required"`
	Nickname         string   `json:"nickname" binding:"required"`
	Birthday         DateTime `json:"birthday" binding:"required"`
	Gender           string   `json:"gender" binding:"required"`
	Region           string   `json:"region,omitempty"`
}
type UpdateUserBody struct {
	Nickname string `json:"nickname,omitempty"`
	Birthday string `json:"birthday,omitempty"`
}

type PostSigninBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (t User) MarshalJSON() ([]byte, error) {
	type TmpJSON User
	return json.Marshal(&struct {
		TmpJSON
		Birthday DateTime `json:"birthday"`
	}{
		TmpJSON:  (TmpJSON)(t),
		Birthday: DateTime(t.Birthday),
	})
}
