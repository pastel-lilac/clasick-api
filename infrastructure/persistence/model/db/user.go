package db

type User struct {
	BaseModel
	Name         string `gorm:"not null"`
	Password     string `gorm:"not null"`
	Email        string
	PhoneNum     string
	AccessToken  string `gorm:"not null"`
	IconPath     string
	Introduction string
}
