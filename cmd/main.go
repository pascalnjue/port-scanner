package main

import (
	"github.com/pascalnjue/port-scanner/pkg/utils"
	"log"
	"os"
)

// main is the entry point for the application.
func main() {
	// prompt the user for the host and ports to scan
	host, portsToScan := utils.PromptForScannerOptions(os.Stdin)

	log.Println("host:", host)
	log.Println("ports to scan:", portsToScan)
}
