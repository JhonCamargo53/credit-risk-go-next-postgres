package seed

import "gorm.io/gorm"

func SeedDocumentTypes(db *gorm.DB) error {
	query := `
    INSERT INTO document_types (code, description, created_at, updated_at)
    VALUES
        ('CC', 'Cédula de Ciudadanía', NOW(), NOW()),
        ('TI', 'Tarjeta de Identidad', NOW(), NOW()),
        ('CE', 'Cédula de Extranjería', NOW(), NOW()),
        ('PP', 'Pasaporte', NOW(), NOW()),
        ('NIT', 'Número de Identificación Tributaria', NOW(), NOW())
    ON CONFLICT (code) DO NOTHING;
    `
	return db.Exec(query).Error
}
