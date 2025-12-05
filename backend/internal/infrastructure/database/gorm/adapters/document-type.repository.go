package adapters

import (
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/ports"
	"gorm.io/gorm"
)

type DocumentTypeGormRepository struct {
	db *gorm.DB
}

func NewDocumentTypeGormRepository(db *gorm.DB) ports.DocumentTypeRepository {
	return &DocumentTypeGormRepository{
		db: db,
	}
}

func (r *DocumentTypeGormRepository) FindAll() ([]models.DocumentType, error) {
	var documentTypes []models.DocumentType
	if err := r.db.Find(&documentTypes).Error; err != nil {
		return nil, err
	}
	return documentTypes, nil
}
func (r *DocumentTypeGormRepository) ExistsByID(id uint) (bool, error) {
	var documentTypes []models.DocumentType
	return r.db.Model(documentTypes).Where("id = ?", id).Select("count(*) > 0").Find(&[]bool{}).Error == nil, nil
}
