package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
)

const (
	DefaultListenAddr = ":2224"
	PathProject       = "/go/src/github.com/noreng-jg/sshserver"
)

const (
	TypeItPty  = "itpty"
	TypeCmdPty = "cmdpty"
	TypeSCP    = "scp"
	TypeNotPty = "notpty"
)

func HandlePty(s ssh.Session) {
	pty, _, _ := s.Pty()

	commands := s.RawCommand()

	logrus.Info(fmt.Sprintf("The commands are: %s\n", commands))
	logrus.Info(fmt.Sprintf("The pty term is: %s\n", func() string {
		if pty.Term != "" {
			return pty.Term
		}
		return "empty"
	}()))

	switch {
	case pty.Term != "":
		io.WriteString(s, fmt.Sprintf("I am a %s\n", TypeItPty))
	case pty.Term == "" && strings.Contains(commands, "scp -t"):
		io.WriteString(s, fmt.Sprintf("I am a %s\n", TypeSCP))
	case pty.Term == "" && commands != "":
		io.WriteString(s, fmt.Sprintf("I am a %s\n", TypeCmdPty))
	default:
		io.WriteString(s, fmt.Sprint("I am a %s\n", TypeNotPty))
	}
}

func NewServer() *ssh.Server {
	server := &ssh.Server{
		Addr:    DefaultListenAddr,
		Handler: ssh.Handler(HandlePty),
	}

	if err := server.SetOption(ssh.HostKeyFile(fmt.Sprintf("%s/.rsa", PathProject))); err != nil {
		logrus.Error(err)
		return nil
	}

	return server
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Info("Listening on port ", DefaultListenAddr)

	s := NewServer()
	if s == nil {
		logrus.Error("Failed to create server")
		return
	}

	logrus.Fatal(s.ListenAndServe())
}
