package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {


	args := os.Args[2:]
	timeout := time.Second*2

	for _, port := range args {
		link := fmt.Sprintf("%s:%s", os.Args[1], port)
		conn, err := net.DialTimeout("tcp", link, timeout)
		if err != nil {
			fmt.Printf("Failed To connect: %s:%v\n", port, err)
			continue
		}

		fmt.Printf("Connected: %s\n", port)
		conn.Close()

	}

}
