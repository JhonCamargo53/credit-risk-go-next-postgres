package customerAsset

import (
	"errors"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

/* Mock de CustomerAssetRepository */

type MockCustomerAssetRepository struct {
	Assets map[uint]*models.CustomerAsset
	NextID uint

	ErrFindAll  error
	ErrFindByID error
	ErrCreate   error
	ErrUpdate   error
	ErrDelete   error
}

var _ ports.CustomerAssetRepository = (*MockCustomerAssetRepository)(nil)

func NewMockCustomerAssetRepository(initial []*models.CustomerAsset) *MockCustomerAssetRepository {
	m := &MockCustomerAssetRepository{
		Assets: make(map[uint]*models.CustomerAsset),
		NextID: 1,
	}

	for _, a := range initial {
		if a.ID == 0 {
			a.ID = m.NextID
			m.NextID++
		}
		m.Assets[a.ID] = a
		if a.ID >= m.NextID {
			m.NextID = a.ID + 1
		}
	}

	return m
}

func (m *MockCustomerAssetRepository) FindAll(creditRequestID *uint) ([]models.CustomerAsset, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}

	var res []models.CustomerAsset
	for _, a := range m.Assets {
		if creditRequestID != nil {
			if a.CreditRequestID == *creditRequestID {
				res = append(res, *a)
			}
		} else {
			res = append(res, *a)
		}
	}
	return res, nil
}

func (m *MockCustomerAssetRepository) FindByID(id uint) (*models.CustomerAsset, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}

	if a, ok := m.Assets[id]; ok {
		return a, nil
	}
	return nil, nil
}

func (m *MockCustomerAssetRepository) CountByCreditRequestID(creditRequestID uint) (int64, error) {
	var count int64
	for _, a := range m.Assets {
		if a.CreditRequestID == creditRequestID {
			count++
		}
	}
	return count, nil
}

func (m *MockCustomerAssetRepository) Create(ca *models.CustomerAsset) error {
	if m.ErrCreate != nil {
		return m.ErrCreate
	}

	if ca.ID == 0 {
		ca.ID = m.NextID
		m.NextID++
	}
	m.Assets[ca.ID] = ca
	return nil
}

func (m *MockCustomerAssetRepository) Update(id uint, data *models.CustomerAsset) (*models.CustomerAsset, error) {
	if m.ErrUpdate != nil {
		return nil, m.ErrUpdate
	}

	existing, ok := m.Assets[id]
	if !ok {
		return nil, errors.New("asset no encontrado")
	}

	if data.CreditRequestID != 0 {
		existing.CreditRequestID = data.CreditRequestID
	}
	if data.AssetID != 0 {
		existing.AssetID = data.AssetID
	}
	if data.CustomerID != 0 {
		existing.CustomerID = data.CustomerID
	}
	if data.MarketValue != 0 {
		existing.MarketValue = data.MarketValue
	}
	if data.Description != "" {
		existing.Description = data.Description
	}

	return existing, nil
}

func (m *MockCustomerAssetRepository) Delete(id uint) error {
	if m.ErrDelete != nil {
		return m.ErrDelete
	}
	delete(m.Assets, id)
	return nil
}

/*  Mock de CustomerRepository */

type MockCustomerRepository struct {
	Customers map[uint]*models.Customer

	ErrFindByID error
}

var _ ports.CustomerRepository = (*MockCustomerRepository)(nil)

func NewMockCustomerRepository(customers []*models.Customer) *MockCustomerRepository {
	m := &MockCustomerRepository{
		Customers: make(map[uint]*models.Customer),
	}
	for _, c := range customers {
		m.Customers[c.ID] = c
	}
	return m
}

func (m *MockCustomerRepository) FindAllOrderedByCreatedDesc() ([]models.Customer, error) {
	var res []models.Customer
	for _, c := range m.Customers {
		res = append(res, *c)
	}
	return res, nil
}

func (m *MockCustomerRepository) FindByID(id uint) (*models.Customer, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}
	if c, ok := m.Customers[id]; ok {
		return c, nil
	}
	return nil, nil
}

func (m *MockCustomerRepository) FindByEmail(email string) (*models.Customer, error) {
	for _, c := range m.Customers {
		if c.Email == email {
			return c, nil
		}
	}
	return nil, nil
}

func (m *MockCustomerRepository) FindByDocument(documentNumber string, documentTypeID uint, excludeID *uint) (*models.Customer, error) {
	return nil, nil
}

func (m *MockCustomerRepository) Create(customer *models.Customer) error {
	m.Customers[customer.ID] = customer
	return nil
}

func (m *MockCustomerRepository) Update(id uint, customerData *models.Customer) (*models.Customer, error) {
	return nil, nil
}

func (m *MockCustomerRepository) Delete(id uint) error {
	delete(m.Customers, id)
	return nil
}

/* Mock de AssetRepository */

type MockAssetRepository struct {
	Assets map[uint]*models.Asset

	ErrFindByID error
}

var _ ports.AssetRepository = (*MockAssetRepository)(nil)

func NewMockAssetRepository(assets []*models.Asset) *MockAssetRepository {
	m := &MockAssetRepository{
		Assets: make(map[uint]*models.Asset),
	}
	for _, a := range assets {
		m.Assets[a.ID] = a
	}
	return m
}

func (m *MockAssetRepository) FindByID(id uint) (*models.Asset, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}
	if a, ok := m.Assets[id]; ok {
		return a, nil
	}
	return nil, nil
}

func (m *MockAssetRepository) FindAll() ([]models.Asset, error) {
	var res []models.Asset
	for _, a := range m.Assets {
		res = append(res, *a)
	}
	return res, nil
}

/*  Mock de CreditRequestRepository */

type MockCreditRequestRepository struct {
	CreditRequests map[uint]*models.CreditRequest

	UpdateRiskCalled bool
	LastRiskID       uint
	LastScore        float64
	LastCategory     string
	LastExplanation  string

	ErrFindByID   error
	ErrUpdateRisk error
}

var _ ports.CreditRequestRepository = (*MockCreditRequestRepository)(nil)

func NewMockCreditRequestRepository(requests []*models.CreditRequest) *MockCreditRequestRepository {
	m := &MockCreditRequestRepository{
		CreditRequests: make(map[uint]*models.CreditRequest),
	}
	for _, cr := range requests {
		m.CreditRequests[cr.ID] = cr
	}
	return m
}

func (m *MockCreditRequestRepository) FindAll(customerID *uint) ([]models.CreditRequest, error) {
	var res []models.CreditRequest
	for _, cr := range m.CreditRequests {
		if customerID != nil {
			if cr.CustomerID == *customerID {
				res = append(res, *cr)
			}
		} else {
			res = append(res, *cr)
		}
	}
	return res, nil
}

func (m *MockCreditRequestRepository) FindByID(id uint) (*models.CreditRequest, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}
	if cr, ok := m.CreditRequests[id]; ok {
		return cr, nil
	}
	return nil, nil
}

func (m *MockCreditRequestRepository) HasRequestsByCustomerID(customerID uint) (bool, error) {
	for _, cr := range m.CreditRequests {
		if cr.CustomerID == customerID {
			return true, nil
		}
	}
	return false, nil
}

func (m *MockCreditRequestRepository) Create(creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	m.CreditRequests[creditRequest.ID] = creditRequest
	return creditRequest, nil
}

func (m *MockCreditRequestRepository) Update(id uint, creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	m.CreditRequests[id] = creditRequest
	return creditRequest, nil
}

func (m *MockCreditRequestRepository) Delete(id uint) error {
	delete(m.CreditRequests, id)
	return nil
}

func (m *MockCreditRequestRepository) UpdateCreditRiskEvaluation(id uint, score float64, category string, explanation string) (*models.CreditRequest, error) {
	if m.ErrUpdateRisk != nil {
		return nil, m.ErrUpdateRisk
	}
	m.UpdateRiskCalled = true
	m.LastRiskID = id
	m.LastScore = score
	m.LastCategory = category
	m.LastExplanation = explanation
	return nil, nil
}

func (m *MockCreditRequestRepository) FindDataToEvaluateRisk(id uint) (models.Customer, *models.CreditRequest, []models.CreditRequest, []models.CustomerAsset, error) {
	cr, ok := m.CreditRequests[id]
	if !ok {
		return models.Customer{}, nil, nil, nil, errors.New("credit request no encontrada")
	}

	// Para tests: devolvemos datos mínimos
	customer := models.Customer{
		ID:            cr.CustomerID,
		Name:          "Test Customer",
		MonthlyIncome: 5_000_000,
	}

	otherCredits := []models.CreditRequest{}
	assets := []models.CustomerAsset{}

	return customer, cr, otherCredits, assets, nil
}

/*  Mock de RiskEvaluator */

type MockRiskEvaluator struct {
	// Valores que vamos a devolver
	Score       float64
	Category    string
	Explanation string
	Err         error

	// Para verificar que se llamó
	Called        bool
	CalledWithCR  uint
	CalledWithCus uint
}

var _ ports.RiskEvaluator = (*MockRiskEvaluator)(nil)

func (m *MockRiskEvaluator) Evaluate(customer models.Customer, currentCreditRequest models.CreditRequest,
	otherCredits []models.CreditRequest, assets []models.CustomerAsset) (float64, string, string, error) {

	m.Called = true
	m.CalledWithCR = currentCreditRequest.ID
	m.CalledWithCus = customer.ID

	if m.Err != nil {
		return 0, "", "", m.Err
	}
	return m.Score, m.Category, m.Explanation, nil
}
