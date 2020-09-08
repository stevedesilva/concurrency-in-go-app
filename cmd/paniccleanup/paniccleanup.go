package main

import (
	"errors"
	"os"
)

func main() {

}

func OpenCSV(filename string) (file *os.File, err error) {
	defer func(){
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}
}

func RemoveEmptyLines(f *os.File) {
	panic(errors.New("failed parse"))
}