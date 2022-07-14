package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/db/models"
	"github.com/Levor/birthday/internal/db/repositories"
)

type UserService struct {
	ur  *repositories.UserRepository
	cfg *config.Config
	t   *Token
}

func NewUserService(ur *repositories.UserRepository, cfg *config.Config, t *Token) *UserService {
	return &UserService{ur: ur, cfg: cfg, t: t}
}

func (us *UserService) Login(resp models.User) (string, error) {
	user, err := us.ur.FindByLogin(resp.Login)
	h := md5.Sum([]byte(resp.Password))
	password := hex.EncodeToString(h[:])
	user.Login = resp.Login
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("Wrong login")
	} else {
		if user.Password == password {
			err := us.ur.ChangeStatus(user.Login, true)
			if err != nil {
				return "", err
			}
			token, err := us.t.Issue(user)
			return token, nil
		} else {
			return "", errors.New("Wrong password")
		}
	}
}
func (us *UserService) Logout(resp models.User) error {
	err := us.ur.ChangeStatus(resp.Login, false)
	if err != nil {
		return err
	}
	return nil
}
func (us *UserService) SignUp(user *models.User) (*models.User, error) {
	h := md5.Sum([]byte(user.Password))
	password := hex.EncodeToString(h[:])
	user.Password = password
	newUser, err := us.ur.Create(user)
	if err != nil {
		return &models.User{}, err
	}
	newUser.Password = ""
	return newUser, nil
}
func (us *UserService) ChangePassword(login, password string) error {
	h := md5.Sum([]byte(password))
	newPass := hex.EncodeToString(h[:])
	err := us.ur.ChangePassword(login, newPass)
	if err != nil {
		return err
	}
	return nil
}
