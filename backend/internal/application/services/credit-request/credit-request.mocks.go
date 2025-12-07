package creditRequest

import (
	"errors"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

/* Mock de CreditRequestRepository */

type MockCreditRequestRepository struct {
	Requests map[uint]*models.CreditRequest
	NextID   uint

	ErrFindAll    error
	ErrFindByID   error
	ErrCreate     error
	ErrUpdate     error
	ErrDelete     error
	ErrUpdateRisk error
	ErrFindData   error

	UpdateRiskCalled bool
	LastRiskID       uint
	LastScore        float64
	LastCategory     string
	LastExplanation  string
}

var _ ports.CreditRequestRepository = (*MockCreditRequestRepository)(nil)

func NewMockCreditRequestRepository(initial []*models.CreditRequest) *MockCreditRequestRepository {
	m := &MockCreditRequestRepository{
		Requests: make(map[uint]*models.CreditRequest),
		NextID:   1,
	}
	for _, cr := range initial {
		if cr.ID == 0 {
			cr.ID = m.NextID
			m.NextID++
		}
		m.Requests[cr.ID] = cr
		if cr.ID >= m.NextID {
			m.NextID = cr.ID + 1
		}
	}
	return m
}

func (m *MockCreditRequestRepository) FindAll(customerID *uint) ([]models.CreditRequest, error) {
	if m.ErrFindAll != nil {
		return nil, m.ErrFindAll
	}

	var res []models.CreditRequest
	for _, cr := range m.Requests {
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

	cr, ok := m.Requests[id]
	if !ok {
		return nil, errors.New("credit request no encontrada")
	}

	copy := *cr
	return &copy, nil
}

func (m *MockCreditRequestRepository) HasRequestsByCustomerID(customerID uint) (bool, error) {
	for _, cr := range m.Requests {
		if cr.CustomerID == customerID {
			return true, nil
		}
	}
	return false, nil
}

func (m *MockCreditRequestRepository) Create(creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	if m.ErrCreate != nil {
		return nil, m.ErrCreate
	}
	if creditRequest.ID == 0 {
		creditRequest.ID = m.NextID
		m.NextID++
	}
	m.Requests[creditRequest.ID] = creditRequest

	copy := *creditRequest
	return &copy, nil
}

func (m *MockCreditRequestRepository) Update(id uint, creditRequest *models.CreditRequest) (*models.CreditRequest, error) {
	if m.ErrUpdate != nil {
		return nil, m.ErrUpdate
	}
	existing, ok := m.Requests[id]
	if !ok {
		return nil, errors.New("credit request no encontrada")
	}

	creditRequest.ID = id

	*existing = *creditRequest
	m.Requests[id] = existing

	copy := *existing
	return &copy, nil
}

func (m *MockCreditRequestRepository) Delete(id uint) error {
	if m.ErrDelete != nil {
		return m.ErrDelete
	}
	delete(m.Requests, id)
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

	cr, ok := m.Requests[id]
	if !ok {
		return nil, errors.New("credit request no encontrada")
	}

	copy := *cr
	return &copy, nil
}

func (m *MockCreditRequestRepository) FindDataToEvaluateRisk(id uint) (models.Customer, *models.CreditRequest, []models.CreditRequest, []models.CustomerAsset, error) {
	if m.ErrFindData != nil {
		return models.Customer{}, nil, nil, nil, m.ErrFindData
	}

	cr, ok := m.Requests[id]
	if !ok {
		return models.Customer{}, nil, nil, nil, errors.New("credit request no encontrada")
	}

	customer := models.Customer{
		ID:            cr.CustomerID,
		Name:          "Test Customer",
		MonthlyIncome: 5_000_000,
	}

	otherCredits := []models.CreditRequest{}
	assets := []models.CustomerAsset{}

	return customer, cr, otherCredits, assets, nil
}

/* Mock de CustomerRepository */
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

/* Mock de CreditStatusRepository */

type MockCreditStatusRepository struct {
	Statuses map[uint]*models.CreditStatus

	ErrFindByID error
}

var _ ports.CreditStatusRepository = (*MockCreditStatusRepository)(nil)

func NewMockCreditStatusRepository(statuses []*models.CreditStatus) *MockCreditStatusRepository {
	m := &MockCreditStatusRepository{
		Statuses: make(map[uint]*models.CreditStatus),
	}
	for _, s := range statuses {
		m.Statuses[s.ID] = s
	}
	return m
}

func (m *MockCreditStatusRepository) FindAll() ([]models.CreditStatus, error) {
	var res []models.CreditStatus
	for _, s := range m.Statuses {
		res = append(res, *s)
	}
	return res, nil
}

func (m *MockCreditStatusRepository) FindByID(id uint) (*models.CreditStatus, error) {
	if m.ErrFindByID != nil {
		return nil, m.ErrFindByID
	}
	if s, ok := m.Statuses[id]; ok {
		return s, nil
	}
	return nil, nil
}

/* Mock de CustomerAssetRepository */

type MockCustomerAssetRepository struct {
	Assets map[uint]*models.CustomerAsset

	ErrCountByCRID error
}

var _ ports.CustomerAssetRepository = (*MockCustomerAssetRepository)(nil)

func NewMockCustomerAssetRepository(assets []*models.CustomerAsset) *MockCustomerAssetRepository {
	m := &MockCustomerAssetRepository{
		Assets: make(map[uint]*models.CustomerAsset),
	}
	for _, a := range assets {
		m.Assets[a.ID] = a
	}
	return m
}

func (m *MockCustomerAssetRepository) FindAll(creditRequestID *uint) ([]models.CustomerAsset, error) {
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
	if a, ok := m.Assets[id]; ok {
		return a, nil
	}
	return nil, nil
}

func (m *MockCustomerAssetRepository) CountByCreditRequestID(creditRequestID uint) (int64, error) {
	if m.ErrCountByCRID != nil {
		return 0, m.ErrCountByCRID
	}
	var count int64
	for _, a := range m.Assets {
		if a.CreditRequestID == creditRequestID {
			count++
		}
	}
	return count, nil
}

func (m *MockCustomerAssetRepository) Create(ca *models.CustomerAsset) error {
	m.Assets[ca.ID] = ca
	return nil
}

func (m *MockCustomerAssetRepository) Update(id uint, data *models.CustomerAsset) (*models.CustomerAsset, error) {
	m.Assets[id] = data
	return data, nil
}

func (m *MockCustomerAssetRepository) Delete(id uint) error {
	delete(m.Assets, id)
	return nil
}

/* Mock de RiskEvaluator */
type MockRiskEvaluator struct {
	Score       float64
	Category    string
	Explanation string
	Err         error

	Called       bool
	CalledWithCR uint
}

var _ ports.RiskEvaluator = (*MockRiskEvaluator)(nil)

func (m *MockRiskEvaluator) Evaluate(customer models.Customer, currentCreditRequest models.CreditRequest,
	otherCredits []models.CreditRequest, assets []models.CustomerAsset) (float64, string, string, error) {

	m.Called = true
	m.CalledWithCR = currentCreditRequest.ID

	if m.Err != nil {
		return 0, "", "", m.Err
	}
	return m.Score, m.Category, m.Explanation, nil
}
