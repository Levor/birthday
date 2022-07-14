package repositories

import (
	"github.com/Levor/birthday/internal/db/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	var user []*models.User
	if err := r.db.Model(&models.Worker{}).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindByLogin(login string) (*models.User, error) {
	var m []*models.User
	if err := r.db.Where("login = ?", login).First(&m).Error; err != nil {
		return nil, err
	}
	return m[0], nil
}

func (r *UserRepository) Create(m *models.User) (*models.User, error) {
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *UserRepository) ChangePassword(login, password string) error {
	return r.db.Model(&models.User{}).Where("login = ?", login).Update("password", password).Error
}
func (r *UserRepository) ChangeStatus(login string, status bool) error {
	return r.db.Model(&models.User{}).Where("login = ?", login).Update("is_logged_in", status).Error
}

func (r *UserRepository) Delete(id int) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
