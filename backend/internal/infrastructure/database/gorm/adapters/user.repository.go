package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) ports.UserRepository {
	return &UserGormRepository{
		db: db,
	}
}

func (r *UserGormRepository) FindAllOrderedByCreatedDesc() ([]models.User, error) {
	var users []models.User
	if err := r.db.Order("created_at desc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserGormRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserGormRepository) Save(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
