package system

import (
	"bufio"
	"github.com/k20human/roby2000/pkg/logger"
	"go.uber.org/zap"
	"io"
	"os/exec"
)

type Driver struct {
	logger *zap.Logger
}

func New() (*Driver, error) {
	var err error
	var d Driver

	if d.logger, err = logger.New(); err != nil {
		return nil, err
	}

	return &d, err
}

func (d *Driver) Call(bin string, args []string) error {
	cmd := exec.Command(bin, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go d.copyOutput(stdout, false)
	go d.copyOutput(stderr, true)

	err = cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (d *Driver) copyOutput(r io.Reader, isError bool) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if isError {
			d.logger.Error(scanner.Text())
		} else {
			d.logger.Info(scanner.Text())
		}
	}
}
