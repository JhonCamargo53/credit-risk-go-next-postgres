package ports

import "github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"

type CustomerRepository interface {
	FindAllOrderedByCreatedDesc() ([]models.Customer, error)
	FindByID(id uint) (*models.Customer, error)
	FindByEmail(email string) (*models.Customer, error)
	FindByDocument(documentNumber string, documentTypeID uint, excludeID *uint) (*models.Customer, error)
	Create(customer *models.Customer) error
	Update(id uint, customerData *models.Customer) (*models.Customer, error)
	Delete(id uint) error
}
