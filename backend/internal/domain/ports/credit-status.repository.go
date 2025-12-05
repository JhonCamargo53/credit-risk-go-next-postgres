package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type CreditStatusRepository interface {
	FindAll() ([]models.CreditStatus, error)
	FindByID(id uint) (*models.CreditStatus, error)
}
