// +build windows

package main

import (
	"github.com/davidmz/go-pageant"
	"golang.org/x/crypto/ssh/agent"
)

func hasAgent() bool {
	return pageant.Available()
}

func getAgent() agent.Agent {
	return pageant.New()
}
