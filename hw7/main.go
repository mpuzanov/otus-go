package main

import (
	"os"

	"github.com/mpuzanov/otus-go/hw7/internal/goenvdir"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Trace(os.Args[:])

	if len(os.Args) < 3 {
		logrus.Info("usage: go-envdir /path/to/evndir command arg1 arg2")
		os.Exit(0)
	}

	env, err := goenvdir.ReadDir(os.Args[1])
	if err != nil {
		logrus.Fatal(err)
	}

	res := goenvdir.RunCmd(os.Args[2:], env)
	logrus.Trace("ExitCode:", res)
}
