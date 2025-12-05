package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type UserRepository interface {
	FindAllOrderedByCreatedDesc() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Save(user *models.User) error
	Delete(id uint) error
}
