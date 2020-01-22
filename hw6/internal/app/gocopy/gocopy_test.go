package gocopy_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	cp "github.com/mpuzanov/otus-go/hw6/internal/app/gocopy"
)

func TestCopy(t *testing.T) {

	fileNameTest := "in_file_test"
	dataTest := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	testCases := []struct {
		desc     string
		fromFile string
		toFile   string
		limit    int
		offset   int
		want     []byte
		err      error
	}{
		{
			desc:     "Test full copy file",
			fromFile: fileNameTest,
			toFile:   "out_file",
			limit:    0,
			offset:   0,
			want:     dataTest,
			err:      nil,
		},
		{
			desc:     "Test offset",
			fromFile: fileNameTest,
			toFile:   "out_file_offset",
			limit:    0,
			offset:   5,
			want:     []byte{5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			err:      nil,
		},
		{
			desc:     "Test offset limit",
			fromFile: fileNameTest,
			toFile:   "out_file_offset_limit",
			limit:    10,
			offset:   5,
			want:     []byte{5, 6, 7, 8, 9, 0, 1, 2, 3, 4},
			err:      nil,
		},
		{
			desc:     "Test err",
			fromFile: fileNameTest,
			toFile:   "out_file_err",
			limit:    10,
			offset:   50, //больше размера файла
			want:     nil,
			err:      errors.New("any error"),
		},
		{
			desc:     "Test file not exists",
			fromFile: "file_not_exists",
			toFile:   "out_file_err",
			limit:    0,
			offset:   0,
			want:     nil,
			err:      errors.New("any error"),
		},
		{
			desc:     "Test copy device",
			fromFile: "/dev/urandom",
			toFile:   "out_file_err",
			limit:    0,
			offset:   0,
			want:     nil,
			err:      errors.New("any error"),
		},
	}

	err := createFile(fileNameTest, dataTest)
	if err != nil {
		t.Error(fmt.Errorf("ошибка создания тестового файла %v", err))
	}
	defer func() {
		//Удаляем файл
		err = os.Remove(fileNameTest)
		if err != nil {
			t.Error(fmt.Errorf("ошибка удаления тестового файла %v", err))
		}
	}()

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			err := cp.Copy(tC.fromFile, tC.toFile, tC.limit, tC.offset)
			if err != nil && tC.err == nil || err == nil && tC.err != nil {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, err, tC.err)
			}
			// если без ошибок скопировали - проверяем дальше
			if err == nil {
				got, err := ioutil.ReadFile(tC.toFile)
				if err != nil {
					t.Errorf("%s, error ReadFile %q %v", tC.desc, tC.toFile, err)
				}
				if !bytes.Equal(tC.want, got) {
					t.Errorf("%s, got=%v, expected=%v", tC.desc, got, tC.want)
				}
			}

			if _, err := os.Stat(tC.toFile); err == nil { // if tC.toFile exists Remove file
				err := os.Remove(tC.toFile)
				if err != nil {
					t.Error(fmt.Errorf("ошибка удаления файла %q %v", tC.toFile, err))
				}
			}
		})
	}
}

//createFile функция создания файла для тестов
func createFile(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
