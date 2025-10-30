package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// SSH client configuration
	config := &ssh.ClientConfig{
		User: "secadmin", // Replace with your SSH username
		Auth: []ssh.AuthMethod{
			ssh.Password("Infinera2!"), // Replace with your SSH password
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // WARNING: Insecure for production, use a proper HostKeyCallback
		Timeout:         5 * time.Second,
	}

	// Connect to the SSH server
	conn, err := ssh.Dial("tcp", net.JoinHostPort("100.113.37.2", "22"), config) // Replace with remote host IP
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Create a new SSH session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Run a command on the remote server
	command := "swversion" // Replace with the command you want to execute
	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}

	fmt.Printf("Command output:\n%s\n", output)
}
