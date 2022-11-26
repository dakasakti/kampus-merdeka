package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (u *UserRepository) UserAvail(cred model.User) error {
	err := u.db.Find(&cred).RowsAffected
	if err != 1 {
		return errors.New("record not found")
	}

	return nil
	// TODO: replace this
}

func (u *UserRepository) CheckPassLength(pass string) bool {
	return len(pass) <= 5
}

func (u *UserRepository) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
