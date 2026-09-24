package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	
	if(len(os.Args)) < 3{
		fmt.Println("Not enough arguments: <host name> <port> <port> ...")
	}
	args := os.Args[2:]
	address := os.Args[1]

	start := time.Now()
	for _, a := range args{
		portScanner(a,address)
	}

	timeGone := time.Since(start)
	fmt.Printf("\n\nTotal Time: %v\n",timeGone)
	
}


func portScanner(port string, address string){
	timeout := time.Second*2
	network := fmt.Sprintf("%s:%s",address,port)
	conn, err := net.DialTimeout("tcp",network, timeout)
	if err != nil{
		fmt.Printf("PORT: %s\tError: %v\n",port,err)
		return
	}
	fmt.Printf("PORT: %s\tConnection succesful\n",port)
	conn.Close()
}

