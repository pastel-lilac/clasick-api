package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/tozastation/clasick-core/domain/repository"
	"github.com/tozastation/clasick-core/infrastructure/persistence/model/db"
)
type userRepository struct {
	*gorm.DB
}

func NewUserRepository(orm *gorm.DB) repository.IUserRepository {
	return &userRepository{
		orm,
	}
}

func (repo *userRepository) InsertUser(ctx context.Context, user *db.User) error {
	err := repo.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) SelectExistUser(ctx context.Context, userName string) (*db.User, error) {
	result := new(db.User)
	err := repo.DB.Where("name = ?", userName).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *userRepository) SelectUserFromID(ctx context.Context, userID int32) (*db.User, error) {
	result := new(db.User)
	err := repo.DB.Where("id = ?", userID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}