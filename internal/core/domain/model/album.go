package model

import (
	"github.com/pastel-lilac/clasick-api/internal/core/common"
)

type Album struct {
	common.BaseModel
	Name        string `gorm:"type:varchar(100);"`
	JacketPhoto string `gorm:"type:varchar(100);"`
	// --- Relation --- //
	//Musics []*Music `gorm:"many2many:albums_musics;"`
}
