package model

import (
	"github.com/pastel-lilac/clasick-api/internal/core/common"
)

type Music struct {
	common.BaseModel
	Name string `gorm:"type:varchar(100);"`
	// --- Relation --- //
	//Albums []*Album `gorm:"many2many:albums_musics;"`
}
