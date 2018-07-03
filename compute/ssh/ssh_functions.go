package test

import (
	"fmt"
	"net"
	"os"
	"testing"

	"golang.org/x/crypto/ssh"
)

func createSSHConfig(t *testing.T) *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User: os.Args[1],
		Auth: []ssh.AuthMethod{
			PublicKeyFile(os.Args[2]),
		},
		HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error {
			return nil
		},
	}

	return sshConfig
}

func createSSHTarget(t *testing.T, publicIP string) string {
	host := publicIP
	port := "22"
	target := fmt.Sprintf("%s:%s", host, port)

	return target
}

func createSSHSession(t *testing.T, target string, sshConfig *ssh.ClientConfig) (*ssh.Session, error) {
	connection, err := ssh.Dial("tcp", target, sshConfig)
	if err != nil {
		return nil, err
	}

	session, err2 := connection.NewSession()
	if err2 != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err3 := session.RequestPty("xterm", 80, 40, modes)
	if err3 != nil {
		session.Close()
		return nil, err
	}

	return session, nil
}

func runSSHCommands(t *testing.T, commands string, session *ssh.Session) error {
	err := session.Run(commands)
	if err != nil {
		return err
	}

	return nil
}
