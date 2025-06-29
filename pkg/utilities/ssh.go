package utilities

import (
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

func NewSSHConnection(path, username, address string, port uint) (*goph.Client, error) {
	auth, err := goph.Key(path, "")
	if err != nil {
		return nil, err
	}

	client, err := goph.NewConn(&goph.Config{
		User:     username,
		Addr:     address,
		Auth:     auth,
		Port:     port,
		Timeout:  goph.DefaultTimeout,
		Callback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		return nil, err
	}
	return client, nil
}
