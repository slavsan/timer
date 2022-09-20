package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// https://code-maven.com/display-notification-from-the-mac-command-line

var re = regexp.MustCompile(`^[a-zA-Z0-9 ]*$`)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("invalid arguments count")
	}

	duration, err := time.ParseDuration(os.Args[1])
	check(err)

	message := strings.Join(os.Args[2:], " ")

	if !re.MatchString(message) {
		log.Fatal("invalid message, only alphanumeric input is allowed")
	}

	select {
	case <-time.After(duration):
		notify(message)
	}
}

func notify(msg string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "Time's up!" sound name "Funk"`, msg))

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		fmt.Printf("stdout: %s\n", outb.String())
		fmt.Printf("stderr: %s\n", errb.String())
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
