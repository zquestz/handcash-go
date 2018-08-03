package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/zquestz/handcash-go/api"
)

func main() {
	setupSignalHandlers()

	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: handcash-go <handle> [network]")
		return
	}

	handle := os.Args[1]
	network := "mainnet"

	if len(os.Args) == 3 {
		network = os.Args[2]
	}

	err := process(handle, network)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func process(handle string, network string) error {
	if err := api.SetNetwork(network); err != nil {
		return err
	}

	resp, err := api.Receive(handle)
	if err != nil {
		return err
	}

	json, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", json)

	return nil
}

func setupSignalHandlers() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go interuptSignal(c)
}

func interuptSignal(c <-chan os.Signal) {
	<-c
	os.Exit(0)
}
