package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	err := u.db.Delete(&model.Session{}, "token = ?", tokenTarget).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	err := u.db.Where("username = ?", session.Username).Updates(&session).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	var session model.Session

	err := u.db.Where("token = ?", token).Find(&session).Error
	if err != nil {
		return model.Session{}, err
	}

	res := u.TokenExpired(session)
	if res {
		return model.Session{}, errors.New("token is expired")
	}

	return session, nil
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	var session model.Session

	err := u.db.Where("username = ?", name).Find(&session).RowsAffected
	if err != 1 {
		return model.Session{}, errors.New("not available")
	}

	return session, nil
	// TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session

	err := u.db.Where("token = ?", token).Find(&session).RowsAffected
	if err != 1 {
		return model.Session{}, errors.New("not available")
	}

	return session, nil
	// TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
