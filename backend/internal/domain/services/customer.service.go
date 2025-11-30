package services

import (
	"errors"
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"gorm.io/gorm"
)

type CustomerService struct {
	DB *gorm.DB
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := database.DB.Order("created_at desc").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *CustomerService) GetCustomerByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := database.DB.Select("id", "name", "email", "phone_number", "document_number", "document_type_id", "monthly_income", "status").First(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe cliente con id %d", id)
		}
		return nil, err
	}
	return &customer, nil
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	var existing models.Customer

	err := database.DB.Where("email = ?", customer.Email).First(&existing).Error
	if err == nil {
		return nil, fmt.Errorf("ya existe un cliente con el email %s", customer.Email)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = database.DB.Where("document_number = ? AND document_type_id = ?", customer.DocumentNumber, customer.DocumentTypeId).First(&existing).Error
	if err == nil {
		return nil, fmt.Errorf("ya existe un cliente con este tipo y número de documento")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = database.DB.Create(customer).Error

	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *CustomerService) UpdateCustomer(id uint, customerData *models.Customer) (*models.Customer, error) {
	customer, err := s.GetCustomerByID(id)

	if err != nil {
		return nil, err
	}

	if customerData.DocumentTypeId != 0 {
		var docType models.DocumentType
		if err := database.DB.First(&docType, customerData.DocumentTypeId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("no existe documento con id %d", customerData.DocumentTypeId)
			}
			return nil, err
		}
	}

	if customerData.Email != "" && customer.Email != customerData.Email {
		var existingCustomer models.Customer
		err := database.DB.Where("email = ?", customerData.Email).First(&existingCustomer).Error
		if err == nil && existingCustomer.ID != id {
			return nil, fmt.Errorf("ya existe un cliente con el email %s", customerData.Email)
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if customerData.DocumentNumber != "" && customerData.DocumentTypeId != 0 {
		var existingDoc models.Customer
		err := database.DB.Where(
			"document_number = ? AND document_type_id = ? AND id <> ?",
			customerData.DocumentNumber,
			customerData.DocumentTypeId,
			id,
		).First(&existingDoc).Error
		if err == nil {
			return nil, fmt.Errorf("ya existe un cliente con este tipo y número de documento")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if err := database.DB.Model(&customer).Updates(customerData).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) DeleteCustomer(id uint) error {

	customer, err := s.GetCustomerByID(id)

	if err != nil {
		return err
	}

	var count int64
	if err := database.DB.Model(&models.CreditRequest{}).
		Where("customer_id = ?", id).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("el cliente tiene solicitudes de crédito asociadas")
	}

	if err := database.DB.Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}
