package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"

	"gorm.io/gorm"
)

type CustomerGormRepository struct {
	db *gorm.DB
}

func NewCustomerGormRepository(db *gorm.DB) ports.CustomerRepository {
	return &CustomerGormRepository{
		db: db,
	}
}

func (r *CustomerGormRepository) FindAllOrderedByCreatedDesc() ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.Order("created_at desc").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerGormRepository) FindByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.Select(
		"id",
		"name",
		"email",
		"phone_number",
		"document_number",
		"document_type_id",
		"monthly_income",
		"status",
	).First(&customer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerGormRepository) FindByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.Where("email = ?", email).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerGormRepository) FindByDocument(documentNumber string, documentTypeID uint, excludeID *uint) (*models.Customer, error) {
	var customer models.Customer
	query := r.db.Where("document_number = ? AND document_type_id = ?", documentNumber, documentTypeID)

	if excludeID != nil {
		query = query.Where("id <> ?", *excludeID)
	}

	if err := query.First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerGormRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerGormRepository) Update(id uint, customerData *models.Customer) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&customer).Updates(customerData).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Customer{}, id).Error
}
