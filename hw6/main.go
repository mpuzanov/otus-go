package main

import (
	"flag"
	"fmt"
	cp "hw6/internal/app/gocopy"
	"os"
)

var (
	from, to      string
	offset, limit int
)

func init() {
	flag.StringVar(&from, "from", "", "file copy from")
	flag.StringVar(&to, "to", "", "file copy to")
	flag.IntVar(&offset, "offset", 0, "offset in input file")
	flag.IntVar(&limit, "limit", 0, "limit copy file")
}

func main() {
	flag.Parse()
	if from == "" || to == "" {
		flag.PrintDefaults()
		return
	}
	if err := cp.Copy(from, to, limit, offset); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("copied successfully")
	}

}
