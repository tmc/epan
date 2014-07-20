package main

import (
	"fmt"

	"github.com/tmc/epan"
)

func main() {
	for _, protocol := range epan.ListProtocols() {
		fmt.Println(protocol)
	}
}
