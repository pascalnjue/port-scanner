package main

import (
	"github.com/pascalnjue/port-scanner/pkg/scanner"
	"github.com/pascalnjue/port-scanner/pkg/utils"
	"log"
	"os"
)

// main is the entry point for the application.
func main() {
	// prompt the user for the host and ports to scan then scan the ports
	openPorts := scanner.CheckOpenPorts(utils.PromptForScannerOptions(os.Stdin))
	log.Println("open ports:", openPorts)
}
