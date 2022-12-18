package scanner

import (
	"fmt"
	"log"
	"net"
	"sort"
)

// CheckOpenPorts checks if the given ports are open on the given host.
// It returns a slice of open ports.
func CheckOpenPorts(host string, portsToScan []int) []int {
	// Create two channels, one for the ports and one for the results
	ports := make(chan int, 100)
	results := make(chan int)

	// Launch the worker goroutines
	for i := 0; i < cap(ports); i++ {
		go worker(host, ports, results)
	}

	// Send the ports to be scanned to the ports channel
	go func() {
		for _, port := range portsToScan {
			ports <- port
		}
	}()

	// Wait for the results and add them to the openPorts slice
	var openPorts []int
	for i := 0; i < len(portsToScan); i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	// close the channels when all the work is done
	close(ports)
	close(results)

	// Sort the open ports and return them
	sort.Ints(openPorts)
	return openPorts
}

// worker is launched as a goroutine and attempts to connect to the given host and port.
// If the connection is successful, it sends the port to the results channel.
// If the connection fails, it sends 0 to the results channel.
func worker(host string, ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		results <- port

		err = conn.Close()
		if err != nil {
			log.Printf("error closing connection on port %d: %v", port, err)
		}
	}
}
