package main

import (
	"fmt"
	"gowebclients/goclienttest"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting Go Web Clients......")
	args := os.Args[1:]
	url := args[0]
	count, err := strconv.Atoi(args[1])
	concurrent, conerr := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if conerr != nil {
		fmt.Println(conerr)
		os.Exit(1)
	}

	goclienttest.RunClientTest(count, concurrent, url)
}
