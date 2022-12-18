package utils

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

// PromptForScannerOptions prompts the user for the host and ports to scan.
func PromptForScannerOptions(device io.Reader) (string, []int) {
	// prompt the user for the host to scan
	reader := bufio.NewReader(device)
	host := promptForString(reader, "Enter the host to scan:", true)

	var portsToScan []int
	for {
		// prompt the user for the ports to scan
		startPort := promptForInt(
			reader,
			"Enter the lowest port to scan (between 1 and 65535):",
			true,
		)
		endPort := promptForInt(
			reader,
			"Enter the highest port or leave blank to scan only the lowest port (between 1 and 65535)."+
				"\nLeave blank to scan only one port:",
			false,
		)

		if endPort == 0 {
			portsToScan = append(portsToScan, startPort)
			break
		} else {
			// check if the start port is greater than the end port
			if startPort > endPort {
				log.Println("The start port is greater than the end port:", startPort, ">", endPort)

				// prompt the user to swap the ports
				swapPorts := promptForString(
					reader,
					"Do you want to swap the ports?\n"+
						"Press return to swap the ports or enter 'n' to re-enter the ports:",
					false,
				)

				// swap if the user did not type anything or typed 'y'
				if swapPorts == "" || strings.ToLower(swapPorts) == "y" {
					for i := endPort; i <= startPort; i++ {
						portsToScan = append(portsToScan, i)
					}
					break
				}
			} else {
				for i := startPort; i <= endPort; i++ {
					portsToScan = append(portsToScan, i)
				}
				break
			}
		}
		println("-----------------------------")
	}

	println("-----------------------------")
	log.Println("host to scan:", host)
	log.Println("number of ports to scan:", len(portsToScan))
	println("-----------------------------")

	return host, portsToScan
}

// promptForInt prompts the user persistently for an integer.
func promptForInt(reader *bufio.Reader, promptText string, forceInput bool) int {
	var input int
	for {
		log.Print(promptText)

		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading submitted input:", err)
		} else {
			readString = strings.TrimSpace(readString)

			// if forcing input, check if the input is empty
			if forceInput && readString == "" {
				log.Println("Error: please enter a value")
				continue
			}

			// if not forcing input and the user did not enter anything, assume 0
			if !forceInput && readString == "" {
				readString = "0"
			}

			// if forcing input, check if the input is "0"
			if forceInput && readString == "0" {
				log.Println("Error: please enter a value in the range 1-65535")
				continue
			}

			convInput, convErr := strconv.Atoi(readString)
			if convErr != nil {
				log.Println("Error: could not covert input value to an integer:", convErr)
			} else {
				input = convInput
				break
			}
		}
	}
	return input
}

// promptForString prompts the user persistently for a string.
func promptForString(reader *bufio.Reader, promptText string, forceInput bool) string {
	var input string
	for {
		log.Print(promptText)

		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error: please enter a valid value:", err)
		} else {
			readString = strings.TrimSpace(readString)
			if readString == "" && forceInput {
				log.Println("Error: please enter a non-empty value")
				continue
			}
			input = readString
			break
		}
	}
	return input
}
