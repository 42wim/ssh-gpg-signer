// +build !windows

package main

import (
	"net"
	"os"

	"golang.org/x/crypto/ssh/agent"
)

func getAgent() agent.Agent {
	sshAgent, _ := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(sshAgent)
}
