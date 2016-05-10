package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("readwriter.go") // idiom: return obj, err
	if err != nil {
		log.Fatal(err) // fail fast
	}
	defer fd.Close() // run stacked after returning from current function

	data, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as hex:\n%x\n", data)
	fmt.Printf("Data as string:\n%s\n", data)
}
