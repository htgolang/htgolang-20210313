package runner

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Callback func()

type Runner struct {
	closed chan struct{}
}

func New() *Runner {
	return &Runner{
		closed: make(chan struct{}),
	}
}

func (r *Runner) Run(callback Callback) {
	go func() {
		callback()
		close(r.closed)
	}()
}

func (r *Runner) Wait() {
	// 等待
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-interrupt:
		logrus.Info("interrupt")
	case <-r.closed:
		logrus.Error("runner close")
	}
}
