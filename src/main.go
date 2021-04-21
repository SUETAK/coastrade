package main

import (
	"log"
	"fmt"

	"github.com/geek-line/sum"
)

func main() {
	fmt.Println("Hello Go!")
	result := sum.Sum(2, 3)
	log.Print(result)
}