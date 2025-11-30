package seed

import "gorm.io/gorm"

func SeedRoles(db *gorm.DB) error {
	query := `
    INSERT INTO roles (name, access, status, created_at, updated_at)
    VALUES
        ('ADMIN', 1000, true, NOW(), NOW()),
        ('EMPLOYEE', 100, true, NOW(), NOW())
    ON CONFLICT (name) DO NOTHING;
    `
	return db.Exec(query).Error
}
