package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	contents, err := ioutil.ReadFile("thefile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))
}
