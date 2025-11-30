package seed

import "gorm.io/gorm"

func SeedUsers(db *gorm.DB) error {
	query := `
    INSERT INTO users (name, role_id, email, password, status, created_at, updated_at)
    VALUES
        ('JHON CAMARGO', 1, 'jcamargocelin@gmail.com', '$2a$14$Vqz0iyVJhpE4A4cdeJn.Juv1Iz4jVfNP8w4d7W7h4NgtsmT16ZkGK', true, NOW(), NOW()),
        ('ROBERTO MORALES', 1, 'robertomorales@gmail.com', '$2a$14$Ee4i0WcHJ9pDyR3sQnO9r..C6xHXmYU75Qag8HNDE.hVnC1aiRmKG', true, NOW(), NOW())
    ON CONFLICT (email) DO NOTHING;
    `
	return db.Exec(query).Error
}
