package model

import (
	"github.com/pastel-lilac/clasick-api/internal/core/common"
)

type Artist struct {
	common.BaseModel
	Name string
	// --- Relation --- //
	//Musics []*Music
	//Albums []*Album
}