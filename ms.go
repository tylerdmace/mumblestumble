package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/tylerdmace/mumblestumble/config"
	"github.com/tylerdmace/mumblestumble/network"
	"github.com/tylerdmace/mumblestumble/version"
)

func main() {
	var cfg config.Config

	// Clean shutdowns
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		shutdown()
	}()

	printInfo()

	// Load config file
	cfg.LoadConfig()

	// TODO: Wallet

	// Bootstrap Network & Chain
	network.Bootstrap(cfg)
}

func printInfo() {
	fmt.Printf("Mumblestumble - Alpha (Build: %v - %v)\r\n", version.BuildVersion, version.BuildTimestamp)
}

func shutdown() {
	fmt.Printf("\r\nShutting down... ")

	// Any shutdown work goes here

	fmt.Printf("Done.\r\n")
	os.Exit(0)
}
