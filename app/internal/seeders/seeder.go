package seeders

import (
	"gorm.io/gorm"
)

// RunSeeders executes all seeders in the correct order
func RunSeeders(db *gorm.DB) error {
	// Add all seeders here in the correct order
	seeders := []func(*gorm.DB) error{
		SeedProducts,
	}

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
