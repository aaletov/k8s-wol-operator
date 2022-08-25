package cmd

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type Executor interface {
	Run(cmd string) (out []byte, err error)
}

type GoExecutor struct {
	Logger *logrus.Logger
}

func NewExecutor(logger *logrus.Logger) (e Executor) {
	e = GoExecutor{logger}
	return
}

func (e GoExecutor) Run(command string) (out []byte, err error) {
	splited := strings.Fields(command)
	cmd := exec.Command(splited[0], splited[1:]...)
	out, err = cmd.CombinedOutput()
	return
}

type MockExecutor struct {
	Out   []byte
	Error error
}

func (e MockExecutor) Run(cmd string) (out []byte, err error) {
	out = e.Out
	err = e.Error
	return
}
