package cli

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hsmfawaz/hsm-backup/models"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func SyncFileHandler(c *cli.Context, db *gorm.DB) error {
	filePath := c.String("file")

	// Validate if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return cli.Exit("File does not exist: "+filePath, 1)
	}

	//read the file and parse it into json
	file, err := os.Open(filePath)
	if err != nil {
		return cli.Exit("Failed to open file: "+err.Error(), 1)
	}
	defer file.Close()

	var config []models.App
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return cli.Exit("Failed to parse JSON file: "+err.Error(), 1)
	}

	log.Printf("Synchronizing backup system with file: %s", filePath)

	err = syncApps(config, db)

	if err != nil {
		return cli.Exit("Failed to synchronize apps: "+err.Error(), 1)
	}

	return nil
}

func syncApps(config []models.App, db *gorm.DB) error {
	var lastError error

	for _, app := range config {
		var existingApp models.App
		if err := db.Where("name = ?", app.Name).First(&existingApp).Error; err != nil {
			lastError = err

			if err != gorm.ErrRecordNotFound {
				continue
			}

			if err := db.Omit("Policies").Create(&app).Error; err != nil {
				lastError = err
				log.Printf("Failed to create app: %s, %s", app.Name, err.Error())
				continue
			}

			log.Printf("Created new app: %s", app.Name)
			syncPolicies(db, &app)
			continue

		}

		// App exists, update it
		if err := db.Model(&existingApp).Omit("Policies").Updates(app).Error; err != nil {
			lastError = err
			log.Printf("Failed to update app: %s, %s", existingApp.Name, err.Error())
			continue
		}

		existingApp.Policies = app.Policies
		syncPolicies(db, &existingApp)
		log.Printf("Updated app: %s", existingApp.Name)
	}

	return lastError
}

func syncPolicies(db *gorm.DB, app *models.App) {
	for _, policy := range app.Policies {
		var existingPolicy models.BackupPolicy
		if err := db.Where("app_id = ? AND name = ?", app.ID, policy.Name).First(&existingPolicy).Error; err != nil {

			if err != gorm.ErrRecordNotFound {
				log.Printf("Failed to check existing policy: %s", err.Error())
				continue
			}

			policy.AppID = app.ID

			if err := db.Create(&policy).Error; err != nil {
				log.Printf("Failed to create policy: %s", err.Error())
			} else {
				log.Printf("Created new policy: %s, %s", app.Name, policy.Name)
			}
			continue
		}

		// Policy exists, update it
		if err := db.Model(&existingPolicy).Updates(policy).Error; err != nil {
			log.Printf("Failed to update policy: %s", err.Error())
		} else {
			log.Printf("Updated policy: %s, %s", app.Name, policy.Name)
		}
	}
}
