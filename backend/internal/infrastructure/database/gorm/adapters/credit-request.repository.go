package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type CreditRequestGormRepository struct {
	db *gorm.DB
}

func NewCreditRequestGormRepository(db *gorm.DB) ports.CreditRequestRepository {
	return &CreditRequestGormRepository{
		db: db,
	}
}

func (r *CreditRequestGormRepository) FindAll(customerID *uint) ([]models.CreditRequest, error) {
	var creditRequests []models.CreditRequest

	query := r.db
	if customerID != nil {
		query = query.Where("customer_id = ?", *customerID)
	}

	if err := query.Order("created_at desc").Find(&creditRequests).Error; err != nil {
		return nil, err
	}

	return creditRequests, nil
}

func (r *CreditRequestGormRepository) FindByID(id uint) (*models.CreditRequest, error) {
	var creditRequest models.CreditRequest
	if err := r.db.First(&creditRequest, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &creditRequest, nil
}

func (r *CreditRequestGormRepository) HasRequestsByCustomerID(customerID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&models.CreditRequest{}).Where("customer_id = ?", customerID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CreditRequestGormRepository) Create(cr *models.CreditRequest) (*models.CreditRequest, error) {
	if err := r.db.Create(cr).Error; err != nil {
		return nil, err
	}
	return cr, nil
}

func (r *CreditRequestGormRepository) Update(id uint, crData *models.CreditRequest) (*models.CreditRequest, error) {
	var cr models.CreditRequest
	if err := r.db.First(&cr, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&cr).Updates(crData).Error; err != nil {
		return nil, err
	}

	return &cr, nil
}

func (r *CreditRequestGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.CreditRequest{}, id).Error
}

func (r *CreditRequestGormRepository) UpdateCreditRiskEvaluation(id uint, score float64, category string, explanation string) (*models.CreditRequest, error) {

	err := r.db.Model(&models.CreditRequest{}).Where("id = ?", id).Updates(map[string]interface{}{
		"risk_score":       score,
		"risk_category":    category,
		"risk_explanation": explanation,
	}).Error

	if err != nil {
		return nil, err
	}

	creditRequest, err := r.FindByID(id)

	if err != nil {
		return nil, err
	}

	return creditRequest, nil

}

func (r *CreditRequestGormRepository) FindDataToEvaluateRisk(id uint) (models.Customer, *models.CreditRequest, []models.CreditRequest, []models.CustomerAsset, error) {

	var customer models.Customer
	var creditRequest models.CreditRequest
	var previousRequests []models.CreditRequest
	var customerAssets []models.CustomerAsset

	if err := r.db.Preload("Customer").First(&creditRequest, id).Error; err != nil {
		return customer, nil, nil, nil, err
	}

	customer = creditRequest.Customer

	if err := r.db.Where("customer_id = ? AND id <> ?", customer.ID, creditRequest.ID).Find(&previousRequests).Error; err != nil {
		return customer, nil, nil, nil, err
	}

	if err := r.db.Where("credit_request_id = ?", creditRequest.ID).Find(&customerAssets).Error; err != nil {
		return customer, nil, nil, nil, err
	}

	return customer, &creditRequest, previousRequests, customerAssets, nil
}
