package services

import (
	"errors"
	"fmt"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	risk "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/ai"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database"
	"gorm.io/gorm"
)

type CustomerAssetService struct {
	DB *gorm.DB
}

func (s *CustomerAssetService) GetAllCustomerAssets(creditRequestId *uint) ([]models.CustomerAsset, error) {
	query := database.DB.Model(&models.CustomerAsset{})

	if creditRequestId != nil {

		var creditRequest models.CreditRequest
		err := database.DB.First(&creditRequest, *creditRequestId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("no existe solicitud de credito %d", *creditRequestId)
			}
			return nil, err
		}

		query = query.Where("credit_request_id = ?", *creditRequestId)
	}

	var customerAssets []models.CustomerAsset
	err := query.Order("created_at desc").Find(&customerAssets).Error
	if err != nil {
		return nil, err
	}

	return customerAssets, nil
}

func (s *CustomerAssetService) GetCustomerAssetByID(id uint) (*models.CustomerAsset, error) {
	var customerAsset models.CustomerAsset
	err := database.DB.First(&customerAsset, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe bien del cliente de id %d", id)
		}
		return nil, err
	}
	return &customerAsset, nil
}

func (s *CustomerAssetService) CreateCustomerAsset(customerAsset *models.CustomerAsset) (*models.CustomerAsset, error) {

	var customerId = customerAsset.CustomerID

	fmt.Println(customerAsset.CustomerID)
	var customer models.Customer
	err := database.DB.First(&customer, customerId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe cliente con id %d", customerId)
		}
		return nil, err
	}

	var assetId = customerAsset.AssetID
	var asset models.Asset
	err = database.DB.First(&asset, assetId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe bien con id %d", assetId)
		}
	}

	err = database.DB.Create(customerAsset).Error

	if err != nil {
		return nil, err
	}

	var creditRequest models.CreditRequest
	err = database.DB.First(&creditRequest, customerAsset.CreditRequestID).Error
	if err != nil {
		return nil, fmt.Errorf("no se pudo obtener la solicitud de crédito: %w", err)
	}

	_, err = risk.EvaluateCreditRisk(database.DB, &creditRequest)
	if err != nil {
		return nil, err
	}

	return customerAsset, nil
}

func (s *CustomerAssetService) UpdateCustomerAsset(id uint, customerAssetData *models.CustomerAsset) (*models.CustomerAsset, error) {

	customerAsset, err := s.GetCustomerAssetByID(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(customerAssetData)

	var customerId = customerAssetData.CustomerID

	var customer models.Customer
	err = database.DB.Select("id", "name", "email", "phone_number", "document_number", "document_type_id", "monthly_income", "status").First(&customer, customerId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe cliente con id %d", customerId)
		}
		return nil, err
	}

	var assetId = customerAssetData.AssetID
	var asset models.Asset
	err = database.DB.First(&asset, assetId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no existe bien con id %d", assetId)
		}
	}

	err = database.DB.Model(customerAsset).Updates(customerAssetData).Error

	if err != nil {
		return nil, err
	}

	return customerAsset, nil
}

func (s *CustomerAssetService) DeleteCustomerAsset(id uint) error {

	customerAsset, err := s.GetCustomerAssetByID(id)
	if err != nil {
		return err
	}

	var creditRequest models.CreditRequest
	if err := database.DB.
		Where("id = ?", customerAsset.CreditRequestID).
		First(&creditRequest).Error; err != nil {
		return fmt.Errorf("no se pudo obtener la solicitud de crédito asociada")
	}

	if err := database.DB.Delete(&customerAsset).Error; err != nil {
		return err
	}

	_, err = risk.EvaluateCreditRisk(database.DB, &creditRequest)
	if err != nil {
		return fmt.Errorf("activo eliminado, pero ocurrió un error al recalcular el riesgo")
	}

	return nil
}
