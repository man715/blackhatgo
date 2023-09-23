package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	// loop through the range of ports coming in from the channel
	for p := range ports {
		// set up the address and port
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		// make the connection
		conn, err := net.Dial("tcp", address)
		// handle any errors
		if err != nil {
			// if there is an error send 0 to the results channel
			results <- 0
			// continue to the next iteration
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// Create the channels to be used
	ports := make(chan int, 10) // buffer 10
	results := make(chan int)
	var openports []int

	// create as many workers as the capacity of the ports
	for i := 0; i < cap(ports); i++ {
		// create a go routine for each iteration
		go worker(ports, results)
	}

	// go routine to send the port numbers to be scanned to the ports channel
	go func() {
		for i := 1; i <= 25; i++ {
			ports <- i
		}
	}()

	// get the results from the workers from the results channel
	for i := 0; i < 25; i++ {
		port := <-results
		// if the result is 0 then there was an error
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(results)
	close(ports)

	// sort the received ports
	sort.Ints(openports)

	// iterate through the openports slice and print each one out.
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}
