package main

import (
	"fmt"
	"os"

	"github.com/ctopher7/sirclo/v2/berat"
)

func main() {
	switch os.Args[1] {
	case "berat":
		beratServer()
	default:
		fmt.Println("wrong argument")
	}
}

func beratServer() {
	err := berat.Server(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
}
