package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func read(filename string) ([]byte, error) {
	fd, err := os.Open(filename) // idiom: return obj, err
	if err != nil {
		log.Fatal(err) // fail fast
	}
	defer fd.Close() // run stacked after returning from current function

	return ioutil.ReadAll(fd)
}

func main() {
	data, err := read("./src/basics/readwriter.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex:\n%x\n", data)
	fmt.Printf("Data as string:\n%s\n", data)
} //END
