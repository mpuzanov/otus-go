package goenvdir

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

var out io.Writer = os.Stdout

//ReadDir сканирует указанный каталог и возвращает все переменные окружения, определенные в нем.
func ReadDir(dir string) (map[string]string, error) {
	logrus.Trace("ReadDir: ", dir)

	res := make(map[string]string)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return res, err
	}

	for _, file := range files {
		fileName := filepath.Join(dir, file.Name())
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			return res, err
		}
		logrus.Trace(file.Name(), "=", string(content))
		res[file.Name()] = string(content)

	}
	logrus.Trace(res)
	return res, nil
}

//RunCmd запускает программу с аргументами (cmd) c переопределённым окружением.
func RunCmd(cmd []string, env map[string]string) int {
	logrus.Tracef("command: %q, args: %v, env: %v\n", cmd[0], cmd[1:], env)

	execCmd := exec.Command(cmd[0], cmd[1:]...)

	//execCmd.Env = os.Environ()
	for key, element := range env {
		execCmd.Env = append(execCmd.Env, fmt.Sprintf("%s=%s", key, element))
	}
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = out
	execCmd.Stderr = os.Stderr

	if err := execCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}
	return 0
}
