package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/crypto/ssh/agent"
)

func rungpg() {
	cmd := exec.Command("gpg", os.Args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("This shouldn't be run manually, it only supports being called from git.")
		os.Exit(0)
	}
	// talk to the agent
	ag := getAgent()

	var uid string

	if runtime.GOOS == "windows" {
		// https://github.com/git-for-windows/git/blob/v2.31.1.windows.1/gpg-interface.c#L449-L452
		if len(os.Args) < 3 {
			rungpg()
			return
		}

		if os.Args[1] != "-bsau" {
			rungpg()
			return
		}

		uid = os.Args[2]
	} else {
		// os.Args[0] --status-fd=2 -bsau uid
		if len(os.Args) < 4 {
			rungpg()
			return
		}

		if os.Args[2] != "-bsau" {
			rungpg()
			return
		}

		uid = os.Args[3]
	}

	snr := bufio.NewScanner(os.Stdin)

	var lines []string

	for snr.Scan() {
		lines = append(lines, snr.Text())
	}

	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// maximum email address is 320 bytes, we add 80 bytes for the name
	uidContents := make([]byte, 400)

	copy(uidContents, []byte(uid))

	var contents []byte

	contents = append(contents, uidContents...)

	// the content we want to have signed
	for _, line := range lines {
		contents = append(contents, []byte(line)...)
		contents = append(contents, []byte("\n")...)
	}

	s, err := ag.(agent.ExtendedAgent).Extension("ssh-gpg-sign@42wim", contents)
	if err != nil {
		log.Fatal(err)
	}

	// See https://github.com/git/git/blob/v2.31.1/gpg-interface.c#L467-L470
	// also need to be on stderr (https://github.com/git/git/blob/v2.31.1/gpg-interface.c#L452)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "[GNUPG:] SIG_CREATED ")
	fmt.Fprintln(os.Stdout, string(s))
}
