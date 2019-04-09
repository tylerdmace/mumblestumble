package main

import (
	"fmt"

	"github.com/tylerdmace/mumblestumble/config"
	"github.com/tylerdmace/mumblestumble/network"
	"github.com/tylerdmace/mumblestumble/version"
)

func main() {
	var cfg config.Config

	printInfo()

	// Load config file
	cfg.LoadConfig()

	fmt.Printf("Listening at %v:%v\r\n", cfg.LocalAddress, cfg.LocalPort)

	// TODO: Wallet

	// Bootstrap Network & Chain
	seeds := network.GetSeeds((byte)(cfg.Network)) // TODO: Support DNS seed servers

	for i := 0; i < len(seeds); i++ {
		// Check connectivity of seed nodes
		network.SendOp(seeds[i], 0x02)
	}
}

func printInfo() {
	fmt.Printf("Mumblestumble - Alpha (Build: %v - %v)\r\n", version.BuildVersion, version.BuildTimestamp)
}
