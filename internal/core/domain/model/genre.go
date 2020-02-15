package model

import "github.com/pastel-lilac/clasick-api/internal/core/common"

type Genre struct {
	common.BaseModel
	Name string `gorm:"type:varchar(100);"`
}
