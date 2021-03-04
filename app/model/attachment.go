package model



type Attachment struct {
	BaseModel
	Uri 				string `gorm:"column:uri;type:varchar(128);not null"`
	Name  				string `gorm:"column:name;type:varchar(128);not null"`
	FileExtension    	string `gorm:"column:fileExtension;type:varchar(256);not null"`
	FileType    		string `gorm:"column:fileType;type:varchar(128)"`
}
