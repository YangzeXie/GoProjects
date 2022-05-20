package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		sep = " "
		s = strconv.Itoa(index) + sep + arg
		fmt.Println(s)
	}
}
