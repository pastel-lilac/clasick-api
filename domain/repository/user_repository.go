package repository

import (
	"context"
	"github.com/tozastation/clasick-core/infrastructure/persistence/model/db"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *db.User) error
	SelectExistUser(ctx context.Context, userName string) (*db.User, error)
	SelectUserFromID(ctx context.Context, userID int32) (*db.User, error)
}
