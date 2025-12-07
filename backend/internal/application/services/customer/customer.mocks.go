package customer

import (
	"errors"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

/*  Mock de CustomerRepository */

type MockCustomerRepository struct {
	Customers map[uint]*models.Customer
	NextID    uint

	ErrFindAll        error
	ErrFindByID       error
	ErrFindByEmail    error
	ErrFindByDocument error
	ErrCreate         error
	ErrUpdate         error
	ErrDelete         error
}

var _ ports.CustomerRepository = (*MockCustomerRepository)(nil)

func NewMockCustomerRepository(initial []*models.Customer) *MockCustomerRepository {
	m := &MockCustomerRepository{
		Customers: make(map[uint]*models.Customer),
		NextID:    1,
	}

	for _, customer := range initial {
		if customer.ID == 0 {
			customer.ID = m.NextID
			m.NextID++
		}
		m.Customers[customer.ID] = customer
		if customer.ID >= m.NextID {
			m.NextID = customer.ID + 1
		}
	}

	return m
}

func (m *MockCustomerRepository) FindAllOrderedByCreatedDesc() ([]models.Customer, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}

	res := make([]models.Customer, 0, len(m.Customers))
	for _, customer := range m.Customers {
		res = append(res, *customer)
	}
	return res, nil
}

func (m *MockCustomerRepository) FindByID(id uint) (*models.Customer, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}

	if customer, ok := m.Customers[id]; ok {
		return customer, nil
	}
	return nil, nil
}

func (m *MockCustomerRepository) FindByEmail(email string) (*models.Customer, error) {
	if m.ErrFindByEmail != nil {
		return nil, m.ErrFindByEmail
	}

	for _, customer := range m.Customers {
		if customer.Email == email {
			return customer, nil
		}
	}
	return nil, nil
}

func (m *MockCustomerRepository) FindByDocument(documentNumber string, documentTypeID uint, excludeID *uint) (*models.Customer, error) {
	if m.ErrFindByDocument != nil {
		return nil, m.ErrFindByDocument
	}

	for _, customer := range m.Customers {
		if excludeID != nil && customer.ID == *excludeID {
			continue
		}
		if customer.DocumentNumber == documentNumber && customer.DocumentTypeId == documentTypeID {
			return customer, nil
		}
	}
	return nil, nil
}

func (m *MockCustomerRepository) Create(customer *models.Customer) error {
	if m.ErrCreate != nil {
		return m.ErrCreate
	}

	if customer.ID == 0 {
		customer.ID = m.NextID
		m.NextID++
	}
	m.Customers[customer.ID] = customer
	return nil
}

func (m *MockCustomerRepository) Update(id uint, customerData *models.Customer) (*models.Customer, error) {
	if m.ErrUpdate != nil {
		return nil, m.ErrUpdate
	}

	customer, ok := m.Customers[id]
	if !ok {
		return nil, errors.New("no existe cliente")
	}

	if customerData.Name != "" {
		customer.Name = customerData.Name
	}
	if customerData.Email != "" {
		customer.Email = customerData.Email
	}
	if customerData.DocumentNumber != "" {
		customer.DocumentNumber = customerData.DocumentNumber
	}
	if customerData.DocumentTypeId != 0 {
		customer.DocumentTypeId = customerData.DocumentTypeId
	}

	return customer, nil
}

func (m *MockCustomerRepository) Delete(id uint) error {
	if m.ErrDelete != nil {
		return m.ErrDelete
	}

	delete(m.Customers, id)
	return nil
}

/* Mock de DocumentTypeRepository */

type MockDocumentTypeRepository struct {
	ExistingIDs map[uint]bool

	ErrExistsByID error
	ErrFindAll    error
}

var _ ports.DocumentTypeRepository = (*MockDocumentTypeRepository)(nil)

func (m *MockDocumentTypeRepository) FindAll() ([]models.DocumentType, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}

	res := make([]models.DocumentType, 0, len(m.ExistingIDs))
	for id := range m.ExistingIDs {
		res = append(res, models.DocumentType{ID: id})
	}
	return res, nil
}

func (m *MockDocumentTypeRepository) ExistsByID(id uint) (bool, error) {
	if m.ErrExistsByID != nil {
		return false, m.ErrExistsByID
	}
	return m.ExistingIDs[id], nil
}

/*  Mock de CreditRequestRepository */
type MockCreditRequestRepository struct {
	HasRequests                map[uint]bool
	ErrHasRequestsByCustomerID error
}

var _ ports.CreditRequestRepository = (*MockCreditRequestRepository)(nil)

func (m *MockCreditRequestRepository) FindAll(customerID *uint) ([]models.CreditRequest, error) {
	return nil, nil
}

func (m *MockCreditRequestRepository) FindByID(id uint) (*models.CreditRequest, error) {
	return nil, nil
}

func (m *MockCreditRequestRepository) HasRequestsByCustomerID(customerID uint) (bool, error) {
	if m.ErrHasRequestsByCustomerID != nil {
		return false, m.ErrHasRequestsByCustomerID
	}
	return m.HasRequests[customerID], nil
}

func (m *MockCreditRequestRepository) Create(creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	return creditRequest, nil
}

func (m *MockCreditRequestRepository) Update(id uint, creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	return nil, nil
}

func (m *MockCreditRequestRepository) Delete(id uint) error {
	return nil
}

func (m *MockCreditRequestRepository) UpdateCreditRiskEvaluation(id uint, score float64, category string, explanation string) (*models.CreditRequest, error) {
	return nil, nil
}

func (m *MockCreditRequestRepository) FindDataToEvaluateRisk(id uint) (models.Customer, *models.CreditRequest, []models.CreditRequest, []models.CustomerAsset, error) {
	return models.Customer{}, nil, nil, nil, nil
}
