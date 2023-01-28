package shellConnector

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func ExecuteCommand(sshHost string, sshPort int, sshUser string, sshPassword string, command string) (string, error) {
	// create ssh config
	config := &ssh.ClientConfig{
		Timeout:         time.Second, // ssh connection timeout
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	// connect ssh
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), config)
	if err != nil {
		log.Printf("ssh.Dial failed: %v", err)
		return "", err
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
		}
	}(client)

	// create session
	session, err := client.NewSession()
	if err != nil {
		log.Printf("client.NewSession failed: %v", err)
		return "", err
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
		}
	}(session)

	// execute command
	b, err := session.CombinedOutput(command)
	if err != nil {
		log.Printf("session.CombinedOutput failed: %v", err)
		return "", err
	}
	log.Printf("result: %s", b)
	return string(b), nil
}
