package seed

import "gorm.io/gorm"

func SeedCreditStatuses(db *gorm.DB) error {
	query := `
		INSERT INTO credit_statuses (id, name, status, created_at, updated_at)
		VALUES
			(1, 'PENDIENTE', true, NOW(), NOW()),
			(2, 'APROBADO', true, NOW(), NOW()),
			(3, 'RECHAZADO', true, NOW(), NOW()),
			(4, 'EN ESTUDIO', true, NOW(), NOW())
		ON CONFLICT (id) DO NOTHING;
	`
	return db.Exec(query).Error
}
