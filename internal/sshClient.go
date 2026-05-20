package internal

import (
	"bytes"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	client *ssh.Client
}

func ConnectWithPassword(user string, host string, port int, password string) (Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	address := fmt.Sprintf("%s@%s:%d", user, host, port)

	conn, err := ssh.Dial("tcp", address, config)

	if err != nil {
		return Client{}, err
	}

	return Client{conn}, nil
}

func (c *Client) Run(command string) (string, string, error) {
	session, err := c.client.NewSession()

	if err != nil {
		return "", "", err
	}

	defer c.Close()

	var stdout bytes.Buffer
	var stdin bytes.Buffer

	session.Stdin = &stdin
	session.Stdout = &stdout

	err = session.Run(command)

	return stdin.String(), stdout.String(), nil
}

func (c *Client) Close() error {
	err := c.client.Close()

	if err != nil {
		return err
	}

	return nil
}
