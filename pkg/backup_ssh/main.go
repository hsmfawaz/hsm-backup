package backup_ssh

import (
	"fmt"
	"log"

	"github.com/hsmfawaz/hsm-backup/pkg/utilities"
	"github.com/melbahja/goph"
)

type SSHBackup struct {
	connection *goph.Client
	path       string
	host       string
	username   string
	port       int
}

func New(path, host, username string, port int) *SSHBackup {
	return &SSHBackup{
		path:     path,
		host:     host,
		username: username,
		port:     port,
	}
}

func (b *SSHBackup) Connect() error {
	if b.host == "" || b.username == "" || b.port == 0 {
		return fmt.Errorf("host, username, and port must be provided")
	}

	log.Printf("Connecting to %s@%s:%d...\n", b.username, b.host, b.port)
	connection, err := utilities.NewSSHConnection(b.path, b.username, b.host, uint(b.port))
	if err != nil {
		return fmt.Errorf("failed to connect to SSH host: %w", err)
	}

	b.connection = connection
	log.Println("Connected successfully.")

	return nil
}

func (b *SSHBackup) GetStats(path string) int64 {
	log.Println("Getting server stats...")

	cmd := fmt.Sprintf("df --output=avail '%s' | tail -1", path)
	out, err := b.connection.Run(cmd)
	if err != nil {
		return 0
	}

	log.Println("Server stats retrieved successfully.")
	size, _ := getFreeDiskSpaceInKB(string(out))

	return size
}
