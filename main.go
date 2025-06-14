package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Hsm Server Backup CLI",
		Usage: "Manage server project backups",
		Commands: []*cli.Command{
			{
				Name:  "app",
				Usage: "Manage apps (add, edit, delete, list)",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "Add a new app",
						Action: func(c *cli.Context) error {
							log.Println("Add app feature coming soon...")
							return nil
						},
					},
					{
						Name:  "edit",
						Usage: "Edit an existing app",
						Action: func(c *cli.Context) error {
							log.Println("Edit app feature coming soon...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "Delete an app",
						Action: func(c *cli.Context) error {
							log.Println("Delete app feature coming soon...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "List all apps",
						Action: func(c *cli.Context) error {
							log.Println("List apps feature coming soon...")
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
