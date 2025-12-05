package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type RoleRepository interface {
	FindAll() ([]models.Role, error)
	FindByID(id uint) (*models.Role, error)
}
