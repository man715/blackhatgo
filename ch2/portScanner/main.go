package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("portquiz.net:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}

}
