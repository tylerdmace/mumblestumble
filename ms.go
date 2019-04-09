package main

import (
	"fmt"

	"github.com/tylerdmace/mumblestumble/network"
	"github.com/tylerdmace/mumblestumble/version"
)

func main() {
	fmt.Printf("Mumblestumble - Alpha (Build: %v - %v)\r\n", version.BuildVersion, version.BuildTimestamp)

	// Handle the loading of wallet data (creating if doesn't exist, seed words, file encryption & password, etc)
	fmt.Printf("Loading wallet... ")
	fmt.Printf("Done.\r\n")

	// Handle the loading of chain state from disk or network seeds (for now, we will only load from network)
	fmt.Printf("Syncing the chain... ")

	seeds := network.GetSeeds(0x01) // Grab testnet seeds

	for i := 0; i < len(seeds); i++ {
		// Check connectivity with ping ops
		network.SendOp(seeds[i], 0x02)
	}

	fmt.Printf("Done.\r\n")
}
