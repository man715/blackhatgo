package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("portquz.net:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open", j)
		}(i)
	}
}
