package seed

import "gorm.io/gorm"

func SeedAssets(db *gorm.DB) error {
	query := `
    INSERT INTO assets (name, description, created_at, updated_at)
    VALUES
        ('INMUEBLE', 'Bien raíz como casas y apartamentos', NOW(), NOW()),
        ('VEHICULO', 'Carros, motos y otros medios de transporte', NOW(), NOW()),
        ('ELECTRODOMESTICO', 'Equipos como neveras, televisores y similares', NOW(), NOW()),
        ('OTRO', 'Cualquier otro activo', NOW(), NOW())
    ON CONFLICT (name) DO NOTHING;
    `
	return db.Exec(query).Error
}
