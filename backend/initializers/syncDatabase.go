package initializers

import "admin-dashboard/models"

func SyncDatabase() {
	// Migrate the schema using the global DB instance
	if err := DB.AutoMigrate(&models.User{},&models.Announcement{},&models.Transaction{}); err != nil {
		// Handle error (log, panic, return an error, etc.)
		panic("Failed to perform auto migration: " + err.Error())
	}
}
