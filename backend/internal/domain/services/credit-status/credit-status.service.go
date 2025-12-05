package creditStatus

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
)

type CreditStatusService struct {
	repo ports.CreditStatusRepository
}

func NewCreditStatusService(repo ports.CreditStatusRepository) *CreditStatusService {
	return &CreditStatusService{
		repo: repo,
	}
}

func (s *CreditStatusService) GetAllCreditStatuses() ([]models.CreditStatus, error) {
	return s.repo.FindAll()
}
