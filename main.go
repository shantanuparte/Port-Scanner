package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	scanned_ports := 0
	fmt.Printf("\t---PORT SCANNER---\n\n")

	if (len(os.Args)) < 3 {
		fmt.Println("Not enough arguments: <host name> <port> <port> ...")
	}

	
	args := os.Args[2:]
	address := os.Args[1]

	start := time.Now()
	fmt.Printf("HOST: %v\n\n",address)
	
	for _, something := range args{
		wg.Add(1)
		scanned_ports++
		go func ()  {
			defer wg.Done()
			portScanner(something,address)
		}()
	}

	wg.Wait()
	timeGone := time.Since(start)
	fmt.Printf("\n\nScanned: %v ports\n",scanned_ports)
	fmt.Printf("Total Time: %v\n", timeGone)

}

func portScanner(port string, address string) {
	timeout := time.Second * 2
	network := fmt.Sprintf("%s:%s", address, port)
	conn, err := net.DialTimeout("tcp", network, timeout)
	if err != nil {
		fmt.Printf("PORT: %s\tCLOSED\n", port)
		return
	}
	fmt.Printf("PORT: %s\tOPEN\n", port)
	conn.Close()
}
