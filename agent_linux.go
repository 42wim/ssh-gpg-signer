package main

import (
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh/agent"
)

func getAgent() agent.Agent {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Fatal(err)
	}

	return agent.NewClient(sshAgent)
}
