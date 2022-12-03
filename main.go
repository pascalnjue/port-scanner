package main

import (
	"fmt"
	"net"
	"sort"
)

var (
	host = "localhost"
	//host = "scanme.nmap.org"
)

// worker is launched as a goroutine to process work from the work channel
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// main is the entry point for the application.
func main() {
	workers := 100
	portsToScan := 65535
	ports := make(chan int, workers)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= portsToScan; i++ {
			ports <- i
		}
	}()

	for i := 0; i < portsToScan; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// close the channels
	close(ports)
	close(results)

	// display the results
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("port %d is open\n", port)
	}
}
