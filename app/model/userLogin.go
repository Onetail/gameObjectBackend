package model

type userLoginType struct {
	Email    string `enum:"EMAIL"`
	Facebook string `enum:"FACEBOOK"`
	Google   string `enum:"GOOGLE"`
}

type UserLogin struct {
	BaseModel
	UserId          string        `gorm:"type:varchar(128);column:userId;not null"`
	Email           string        `gorm:"type:varchar(128);column:email;not null"`
	Password        string        `gorm:"type:varchar(256);column:password;not null"`
	ThirdpartyId    string        `gorm:"type:varchar(128);column:thirdpartyId"`
	ThirdpartyToken string        `gorm:"type:varchar(256);column:thirdpartyToken"`
	LoginType       userLoginType `gorm:"type:varchar(32);column:loginType"`
	IsSuspension    uint          `gorm:"type:tinyint(1);column:isSuspension"`
}

func (UserLogin) TableName() string {
	return "userLogins"
}
