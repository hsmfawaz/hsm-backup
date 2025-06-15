package cli

import (
	"log"
	"strings"
	"time"

	"github.com/hsmfawaz/hsm-backup/models"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func PerformBackupHandler(c *cli.Context, db *gorm.DB) error {
	//todo check if the temp directory is writable
	var apps []models.App
	specificApps := strings.Split(c.String("only"), ",")

	if (len(specificApps) == 1 && specificApps[0] != "") || len(specificApps) > 1 {
		if err := db.Where("name IN ?", specificApps).Find(&apps).Error; err != nil {
			return err
		}
	} else {
		if err := db.Find(&apps).Error; err != nil {
			return err
		}
	}

	if len(apps) == 0 {
		log.Println("No apps found for backup.")
		return nil
	}

	log.Printf("Found %d apps for backup.", len(apps))
	appIDs := make([]uint, len(apps))
	for i, app := range apps {
		appIDs[i] = app.ID
	}
	var policies []models.BackupPolicy
	if err := db.
		Preload("App").
		Where("app_id IN ?", appIDs).
		Where("enabled = ?", true).
		Where("(next_backup_date <= ? or next_backup_date IS NULL)", time.Now()).
		Find(&policies).Error; err != nil {
		return err
	}

	if len(policies) == 0 {
		log.Println("No enabled backup policies found for the selected apps.")
		return nil
	}

	log.Printf("Found %d enabled backup policies for the selected apps.", len(policies))

	for _, policy := range policies {
		log.Printf("Performing backup for app: %s (%s)", policy.App.Name, policy.Name)
		if err := performBackup(policy.App); err != nil {
			log.Printf("Error performing backup for app %s (%s): %v", policy.App.Name, policy.Name, err)
			continue
		}

		db.Model(&policy).Update("next_backup_date", time.Now().Add(time.Duration(policy.Interval)*time.Hour))
		db.Model(&policy.App).Where("id = ?", policy.App.ID).Update("last_backup", time.Now().Format(time.RFC3339))
		db.Create(&models.Backup{
			AppID:          policy.App.ID,
			BackupPolicyID: policy.ID,
		})
	}

	log.Println("All backups completed successfully.")
	return nil
}

func performBackup(app models.App) error {
	log.Printf("Backup completed for app: %s", app.Name)
	//todo if the type is filesystem check if the path exists and readable
	//todo check the directory or file size and update the app.LastDiskUsage
	//todo check the current available disk space
	//todo
	return nil
}
