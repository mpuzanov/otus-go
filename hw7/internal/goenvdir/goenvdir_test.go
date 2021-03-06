package goenvdir

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
	//logrus.SetLevel(logrus.TraceLevel)
}

func TestReadDir(t *testing.T) {

	//=============================
	testDir := "testDir"
	//создаём каталог
	_, err := os.Stat(testDir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(testDir, 0755)
		if errDir != nil {
			logrus.Fatal(err)
		}
	}
	//файлы
	file := filepath.Join(testDir, "ENV1")
	err = ioutil.WriteFile(file, []byte("123"), 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	file = filepath.Join(testDir, "ENV2")
	err = ioutil.WriteFile(file, []byte("var2"), 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	file = filepath.Join(testDir, "ENV_5")
	err = ioutil.WriteFile(file, []byte(" 111 		"), 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	//=============================

	testCases := []struct {
		desc string
		path string
		want map[string]string
		err  error
	}{
		{
			desc: "Test 1",
			path: testDir,
			want: map[string]string{"ENV1": "123", "ENV2": "var2", "ENV_5": "111"},
			err:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := ReadDir(tC.path)
			if err != tC.err {
				t.Errorf("%s, error: %v", tC.desc, err)
			}
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("%s, got=%v, want=%v", tC.desc, got, tC.want)
			}
		})
	}
	//=============================
	//удалить каталог с файлами
	err = os.RemoveAll(testDir)
	if err != nil {
		logrus.Fatal(err)
	}
}

//TestReadDirInDir Тест вложенного каталога
func TestReadDirInDir(t *testing.T) {

	//=============================
	testDir := "testDir"
	testDirIn := "testDir/tmp"
	//создаём каталог
	_, err := os.Stat(testDir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(testDir, 0755)
		if errDir != nil {
			logrus.Fatal(err)
		}
	}
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(testDirIn, 0755)
		if errDir != nil {
			logrus.Fatal(err)
		}
	}

	testCases := []struct {
		desc string
		path string
		want map[string]string
		err  error
	}{
		{
			desc: "Тест вложенного каталога",
			path: testDir,
			want: map[string]string{},
			err:  ErrEnvNotIsDir,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := ReadDir(tC.path)
			if err != tC.err {
				t.Errorf("%s, error: %v", tC.desc, err)
			}
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("%s, got=%v, want=%v", tC.desc, got, tC.want)
			}
		})
	}
	//=============================
	//удалить каталог с файлами
	err = os.RemoveAll(testDir)
	if err != nil {
		logrus.Fatal(err)
	}
}

//TestReadDirInDir Тестирование плохого имени файла для переменной
func TestReadFailEnv(t *testing.T) {

	//=============================
	testDir := "testDir"
	//создаём каталог
	_, err := os.Stat(testDir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(testDir, 0755)
		if errDir != nil {
			logrus.Fatal(err)
		}
	}
	file := filepath.Join(testDir, "ENV3=")
	err = ioutil.WriteFile(file, []byte("var3"), 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	testCases := []struct {
		desc string
		path string
		want map[string]string
		err  error
	}{
		{
			desc: "Тест плохого имени файла для переменной",
			path: testDir,
			want: map[string]string{},
			err:  ErrFailNameEnv,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := ReadDir(tC.path)
			if !errors.Is(err, tC.err) {
				t.Errorf("%s error: %v", tC.desc, err)
			}
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("%s, got=%v, want=%v", tC.desc, got, tC.want)
			}
		})
	}
	//=============================
	//удалить каталог с файлами
	err = os.RemoveAll(testDir)
	if err != nil {
		logrus.Fatal(err)
	}
}

//TestReadDirInDir Тест пустого файла для переменной
func TestReadEmptyFile(t *testing.T) {

	//=============================
	testDir := "testDir"
	//создаём каталог
	_, err := os.Stat(testDir)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(testDir, 0755)
		if errDir != nil {
			logrus.Fatal(err)
		}
	}
	file := filepath.Join(testDir, "PASSWORD")
	err = ioutil.WriteFile(file, []byte(""), 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	testCases := []struct {
		desc string
		path string
		env  string
		want map[string]string
		err  error
	}{
		{
			desc: "Тест пустого файла для переменной",
			path: testDir,
			env:  "PASSWORD",
			want: map[string]string{},
			err:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			os.Setenv(tC.env, "12345")

			got, err := ReadDir(tC.path)
			if !errors.Is(err, tC.err) {
				t.Errorf("%s error: %v", tC.desc, err)
			}
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("%s, got=%v, want=%v", tC.desc, got, tC.want)
			}
			v, ok := os.LookupEnv(tC.env)
			if ok {
				t.Errorf("%s, got=%v, want=%v", tC.desc, v, tC.want)
			}
		})
	}
	//=============================
	//удалить каталог с файлами
	err = os.RemoveAll(testDir)
	if err != nil {
		logrus.Fatal(err)
	}
}

func TestRunCmd(t *testing.T) {
	testCases := []struct {
		desc    string
		command []string
		env     map[string]string
		want    string
	}{
		{
			desc:    "Test 1",
			command: []string{"printenv", "ENV1"},
			env:     map[string]string{"ENV1": "value1", "ENV2": "value2"},
			want:    "value1\n",
		},
		{
			desc:    "Test 2",
			command: []string{"printenv", "ENV2"},
			env:     map[string]string{"ENV1": "value1", "ENV2": "value2"},
			want:    "value2\n",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			saved := out
			defer func() { out = saved }()
			out = new(bytes.Buffer) // Перехватываем вывод

			if got := RunCmd(tc.command, tc.env); got != 0 {
				t.Errorf("%s, ExitCode %v != 0", tc.desc, got)
			}
			got := out.(*bytes.Buffer).String()
			//log.Println("out:", got)
			if got != tc.want {
				t.Errorf("%s, got=%v, want=%v", tc.desc, got, tc.want)
			}
		})
	}
}
