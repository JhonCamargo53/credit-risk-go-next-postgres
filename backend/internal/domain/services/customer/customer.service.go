package customer

import (
	"errors"
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type CustomerService struct {
	customerRepo      ports.CustomerRepository
	documentTypeRepo  ports.DocumentTypeRepository
	creditRequestRepo ports.CreditRequestRepository
}

func NewCustomerService(customerRepo ports.CustomerRepository, documentTypeRepo ports.DocumentTypeRepository, creditRequestRepo ports.CreditRequestRepository,
) *CustomerService {
	return &CustomerService{
		customerRepo:      customerRepo,
		documentTypeRepo:  documentTypeRepo,
		creditRequestRepo: creditRequestRepo,
	}
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.customerRepo.FindAllOrderedByCreatedDesc()
}

func (s *CustomerService) GetCustomerByID(id uint) (*models.Customer, error) {
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("no existe cliente con id %d", id)
	}
	return customer, nil
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) (*models.Customer, error) {

	// Validar email único
	if customer.Email != "" {
		existing, err := s.customerRepo.FindByEmail(customer.Email)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, fmt.Errorf("ya existe un cliente con el email %s", customer.Email)
		}
	}

	// Validar documento único
	if customer.DocumentNumber != "" && customer.DocumentTypeId != 0 {
		existing, err := s.customerRepo.FindByDocument(customer.DocumentNumber, customer.DocumentTypeId, nil)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, fmt.Errorf("ya existe un cliente con este tipo y número de documento")
		}
	}

	// Crear cliente
	if err := s.customerRepo.Create(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) UpdateCustomer(id uint, customerData *models.Customer) (*models.Customer, error) {

	// Obtener el cliente actual
	customer, err := s.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	// Validar tipo de documento si viene
	if customerData.DocumentTypeId != 0 {
		exists, err := s.documentTypeRepo.ExistsByID(customerData.DocumentTypeId)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, fmt.Errorf("no existe documento con id %d", customerData.DocumentTypeId)
		}
	}

	// Validar email único si cambia
	if customerData.Email != "" && customer.Email != customerData.Email {
		existingCustomer, err := s.customerRepo.FindByEmail(customerData.Email)
		if err != nil && !errors.Is(err, nil) {
			return nil, err
		}
		if existingCustomer != nil && existingCustomer.ID != id {
			return nil, fmt.Errorf("ya existe un cliente con el email %s", customerData.Email)
		}
	}

	// Validar documento único si viene info completa
	if customerData.DocumentNumber != "" && customerData.DocumentTypeId != 0 {
		existingDoc, err := s.customerRepo.FindByDocument(customerData.DocumentNumber, customerData.DocumentTypeId, &id)
		if err != nil {
			return nil, err
		}
		if existingDoc != nil {
			return nil, fmt.Errorf("ya existe un cliente con este tipo y número de documento")
		}
	}

	// Actualizar
	updated, err := s.customerRepo.Update(id, customerData)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *CustomerService) DeleteCustomer(id uint) error {

	// Verificar existencia
	_, err := s.GetCustomerByID(id)
	if err != nil {
		return err
	}

	// Verificar si tiene solicitudes de crédito asociadas
	hasCreditRequests, err := s.creditRequestRepo.HasRequestsByCustomerID(id)
	if err != nil {
		return err
	}
	if hasCreditRequests {
		return fmt.Errorf("el cliente tiene solicitudes de crédito asociadas")
	}

	// Eliminar
	if err := s.customerRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
