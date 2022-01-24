package repositories

import (
	"github.com/Levor/birthday/internal/db/models"
	"github.com/jinzhu/gorm"
)

type WorkersRepository struct {
	db *gorm.DB
}

func NewWorkersRepository(db *gorm.DB) *WorkersRepository {
	return &WorkersRepository{
		db: db,
	}
}

func (r *WorkersRepository) FindAll() ([]*models.Worker, error) {
	var Workers []*models.Worker
	if err := r.db.Model(&models.Worker{}).Find(&Workers).Error; err != nil {
		return Workers, err
	}
	return Workers, nil
}

func (r *WorkersRepository) FindByUserId(id int) ([]*models.Worker, error) {
	var m []*models.Worker
	if err := r.db.Model(&m).Where("id = ?", id).First(&m).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (r *WorkersRepository) Create(m *models.Worker) (*models.Worker, error) {
	err := r.db.Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *WorkersRepository) Update(m *models.Worker, nm *models.Worker) error {
	return r.db.Model(&m).Update(nm).Error
}

func (r *WorkersRepository) Delete(id int) error {
	return r.db.Delete(&models.Worker{}, "id = ?", id).Error
}
