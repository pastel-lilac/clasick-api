package db

type User struct {
	BaseModel
	Name         string `gorm:"not null"`
	Password     string `gorm:"not null; type:varchar(500)"`
	Email        string
	PhoneNum     string
	AccessToken  string `gorm:"not null; type:varchar(1000)"`
	IconPath     string
	Introduction string
}
