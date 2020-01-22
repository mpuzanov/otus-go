package gocopy

import (
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

//Copy функция копирования файлов
func Copy(from string, to string, limit int, offset int) error {

	sfi, err := os.Stat(from)
	if err != nil { //проверка на существование файла
		return err
	}
	if !sfi.Mode().IsRegular() {
		// проверка что файл не каталог, не устройство и т.п.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	offset64 := int64(offset)
	limit64 := int64(limit)
	sizeFile := sfi.Size()
	if limit64 <= 0 || limit64 > sizeFile {
		limit64 = sizeFile - offset64
	}
	if offset64 > limit64 {
		return fmt.Errorf("offset error")
	}

	fileIn, err := os.Open(from) // открываем файл для копирования
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileTo, err := os.Create(to) // открываем файл для записи
	if err != nil {
		return err
	}
	defer fileTo.Close()

	buffersize := 512
	buf := make([]byte, buffersize)

	reader := io.NewSectionReader(fileIn, offset64, limit64)

	bar := pb.Full.Start64(limit64)
	barReader := bar.NewProxyReader(reader)

	for {
		read, err := barReader.Read(buf) // читаем из файла

		if err != nil && err != io.EOF {
			return err
		}
		if read == 0 {
			break
		}
		// записываем в файл
		//fmt.Println(string(buf[:read]), offset, read, limit)
		if _, err := fileTo.Write(buf[:read]); err != nil {
			return err
		}
	}
	bar.Finish()

	return nil
}
