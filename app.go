package main

import (
	"fmt"
	"os"
	"yaGotp/core"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Error: invalid input parameters")
		os.Exit(0)
	}

	pin := core.Pin{}
	if !pin.CreatePin(args[0]) {
		fmt.Println("Error: invalid pin")
		os.Exit(0)
	}

	secret := core.Secret{}
	if !secret.CreateSecret(args[1]) {
		fmt.Print("Error: invalid secret")
		os.Exit(0)
	}

	fmt.Println(core.ComputeOtp(pin, secret))
}
