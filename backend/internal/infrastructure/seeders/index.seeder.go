package seed

import (
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {

	if err := SeedRoles(db); err != nil {
		return err
	}

	if err := SeedDocumentTypes(db); err != nil {
		return err
	}

	if err := SeedAssets(db); err != nil {
		return err
	}

	if err := SeedCreditStatuses(db); err != nil {
		return err
	}

	if err := SeedUsers(db); err != nil {
		return err
	}

	return nil
}
