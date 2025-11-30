package services

import (
	"errors"
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	risk "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/ai"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"gorm.io/gorm"
)

type CreditRequestService struct {
	DB *gorm.DB
}

func (s *CreditRequestService) GetAllCreditRequests(customerId *uint) ([]models.CreditRequest, error) {
	var creditRequests []models.CreditRequest
	query := database.DB

	if customerId != nil {
		var customer models.Customer
		err := database.DB.First(&customer, *customerId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("no existe cliente con id %d", *customerId)
			}
			return nil, err
		}
		query = query.Where("customer_id = ?", *customerId)
	}

	err := query.Order("created_at desc").Find(&creditRequests).Error
	if err != nil {
		return nil, err
	}

	return creditRequests, nil
}

func (s *CreditRequestService) GetCreditRequestByID(id uint) (*models.CreditRequest, error) {

	var creditRequest models.CreditRequest
	err := database.DB.First(&creditRequest, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe solicitud de crédito con id %d", id)
		}
		return nil, err
	}

	return &creditRequest, nil

}

func (s *CreditRequestService) CreateCreditRequest(creditRequest *models.CreditRequest) (*models.CreditRequest, error) {

	var customerId = creditRequest.CustomerID

	var customer models.Customer
	err := database.DB.Select("id", "name", "email", "phone_number", "document_number", "document_type_id", "monthly_income", "status").First(&customer, customerId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe cliente con id %d", customerId)
		}
		return nil, err
	}

	var creditStatusId = creditRequest.CreditStatusID
	var creditStatus models.CreditStatus
	err = database.DB.First(&creditStatus, creditStatusId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe el estado de solicitud con id %d", creditStatusId)
		}
	}

	err = database.DB.Create(creditRequest).Error

	if err != nil {
		return nil, err
	}

	_, err = risk.EvaluateCreditRisk(database.DB, creditRequest)
	if err != nil {
		return nil, err
	}

	return creditRequest, nil

}

func (s *CreditRequestService) UpdateCreditRequest(id uint, creditRequestData *models.CreditRequest) (*models.CreditRequest, error) {

	creditRequest, err := s.GetCreditRequestByID(id)

	if err != nil {
		return nil, err
	}

	var customerId = creditRequestData.CustomerID

	var customer models.Customer
	err = database.DB.Select("id", "name", "email", "phone_number", "document_number", "document_type_id", "monthly_income", "status").First(&customer, customerId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe cliente con id %d", customerId)
		}
		return nil, err
	}

	var creditStatusId = creditRequestData.CreditStatusID
	var creditStatus models.CreditStatus
	err = database.DB.First(&creditStatus, creditStatusId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe el estado de solicitud con id %d", creditStatusId)
		}
	}

	err = database.DB.Model(creditRequest).Updates(creditRequestData).Error

	if err != nil {
		return nil, err
	}

	_, err = risk.EvaluateCreditRisk(database.DB, creditRequest)
	if err != nil {
		return nil, err
	}

	return creditRequest, nil

}

func (s *CreditRequestService) DeleteCreditRequest(id uint) error {

	creditRequest, err := s.GetCreditRequestByID(id)
	if err != nil {
		return err
	}

	var count int64
	if err := database.DB.Model(&models.CustomerAsset{}).
		Where("credit_request_id = ?", id).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("no se puede eliminar la solicitud de crédito porque tiene activos asociados")
	}

	if err := database.DB.Delete(&creditRequest).Error; err != nil {
		return err
	}

	return nil
}
