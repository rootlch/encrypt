package main

import (
	"flag"
	"io"
	"log"
	"myprog/encrypt"
	"os"
	"strings"
)

var (
	filename  = flag.String("file", "", "Path to the file that needs to be encrypted")
	key       = flag.String("key", "", "Encryption key")
	overwrite = flag.Bool("overw", false, "Determine if file is overwritten or not")
)

const EXTENSION = "xor"

func init() {
	flag.Parse()
}

func main() {
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x, _ := encrypt.NewXor(*key)
	encoded := x.NewReader(file)

	if !*overwrite {
		file, err = os.Create(encFilename(file.Name()))
	} else {
		file, err = os.OpenFile(*filename, os.O_WRONLY, 0666)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	io.Copy(file, encoded)
}

func encFilename(s string) string {
	sp := strings.Split(s, ".")
	if len(sp) <= 1 {
		return s + EXTENSION
	}

	sp[len(sp)-1] = EXTENSION + "." + sp[len(sp)-1]
	return strings.Join(sp, "")
}
