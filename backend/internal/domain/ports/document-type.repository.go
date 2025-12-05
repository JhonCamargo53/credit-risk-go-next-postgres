package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type DocumentTypeRepository interface {
	FindAll() ([]models.DocumentType, error)
	ExistsByID(id uint) (bool, error)
}
