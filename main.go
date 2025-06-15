package main

import (
	"log"
	"os"

	"github.com/hsmfawaz/hsm-backup/handlers/cli"
	"github.com/hsmfawaz/hsm-backup/models"
	cliV2 "github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := configureDB()
	app := &cliV2.App{
		Name:  "Hsm Server Backup CLI",
		Usage: "Manage server project backups",
		Commands: []*cliV2.Command{
			{
				Name:  "init",
				Usage: "Initialize the backup system",
				Action: func(c *cliV2.Context) error {
					db.AutoMigrate(&models.App{}, &models.BackupPolicy{}, &models.Backup{})
					log.Println("Backup system initialized successfully.")
					return nil
				},
			},
			{
				Name: "sync",
				Flags: []cliV2.Flag{
					&cliV2.StringFlag{
						Name:        "file",
						Aliases:     []string{"f"},
						DefaultText: "config.json",
						Value:       "config.json",
						Usage:       "Path to the configuration file",
					},
				},
				Usage: "Synchronize the backup system",
				Action: func(c *cliV2.Context) error {
					return cli.SyncFileHandler(c, db.WithContext(c.Context))
				},
			},
			{
				Name: "perform",
				Flags: []cliV2.Flag{
					&cliV2.StringFlag{
						Name:    "only",
						Aliases: []string{"o"},
						Usage:   "Perform backup for a specific apps (Comma-separated list)",
					},
				},
				Usage: "Perform backups for all apps or specified apps",
				Action: func(c *cliV2.Context) error {
					return cli.PerformBackupHandler(c, db.WithContext(c.Context).Debug())
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func configureDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}
