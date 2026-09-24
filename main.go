package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
	"strconv"
)

func main() {

	var wg sync.WaitGroup
	scanned_ports := 0

	fmt.Printf("\t---PORT SCANNER---\n\n")

	timePtr := flag.Int("time", 2, "Custom time defined by the user")
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Usage: scanner [-time seconds] <host> <port> <port> ...")
		return
	}

	address := args[0]
	ports := args[1:]

	fmt.Printf("HOST: %v\n\n", address)

	start := time.Now()

	for _, port := range ports {

		portNum, err := strconv.Atoi(port)

		if err != nil || portNum < 1|| portNum > 65535{
			fmt.Printf("Invalid port: %s\n",port)
			continue
		}
		wg.Add(1)
		scanned_ports++

		go func(port string) {
			defer wg.Done()
			portScanner(port, address, *timePtr)
		}(port)
	}

	wg.Wait()

	timeGone := time.Since(start)

	fmt.Printf("\n\nScanned: %v ports\n", scanned_ports)
	fmt.Printf("Total Time: %v\n", timeGone)
}

func portScanner(port string, address string, user_time int) {

	timeout := time.Second * time.Duration(user_time)

	network := fmt.Sprintf("%s:%s", address, port)

	conn, err := net.DialTimeout("tcp", network, timeout)

	if err != nil {
		fmt.Printf("PORT: %s\tCLOSED\n", port)
		return
	}

	fmt.Printf("PORT: %s\tOPEN\n", port)

	conn.Close()
}
