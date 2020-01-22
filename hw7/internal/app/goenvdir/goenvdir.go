package goenvdir

import (
	"fmt"
)

//ReadDir сканирует указанный каталог и возвращает все переменные окружения, определенные в нем.
func ReadDir(dir string) (map[string]string, error) {
	fmt.Println("dir",dir)

	return nil, nil
}

//RunCmd запускает программу с аргументами (cmd) c переопределнным окружением.
func RunCmd(cmd []string, env map[string]string) int {
	fmt.Println("command arg1 arg2:", cmd)
	fmt.Println("env:", env)
	return 0
}
