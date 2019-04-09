package main

import (
	"fmt"

	"github.com/tylerdmace/mumblestumble/version"
)

func main() {
	fmt.Printf("Mumblestumble - Alpha (Build: %v - %v)\r\n", version.BuildVersion, version.BuildTimestamp)
}
