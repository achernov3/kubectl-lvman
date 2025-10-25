package sshclient

import (
	"fmt"
	"kubectl-lvman/internal/config"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func ExecCMD(cfg *config.Config, cmd, host string) ([]byte, error) {
	privateKey, err := os.ReadFile(cfg.SSHKey)
	if err != nil {
		return nil, fmt.Errorf("reading SSH key: %w", err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: cfg.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := net.JoinHostPort(host, cfg.Port)

	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("SSH dial failed: %w", err)
	}

	defer func() error {
		if err := client.Close(); err != nil {
			return fmt.Errorf("failed to close client connection: %w", err)
		}
		return nil
	}()

	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("creating session: %w", err)
	}

	defer func() error {
		if err := session.Close(); err != nil {
			return fmt.Errorf("failed to close ssh session: %w", err)
		}
		return nil
	}()

	stdoutBytes, err := session.CombinedOutput(cmd)
	if err != nil {
		return stdoutBytes, fmt.Errorf("command execution failed: %w. Output: %s",
			err, string(stdoutBytes))
	}

	return stdoutBytes, nil
}
