package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type RoleGormRepository struct {
	db *gorm.DB
}

func NewRoleGormRepository(db *gorm.DB) ports.RoleRepository {
	return &RoleGormRepository{
		db: db,
	}
}

func (r *RoleGormRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleGormRepository) FindByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}
