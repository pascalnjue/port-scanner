package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// TestPromptForScannerOptions tests the PromptForScannerOptions function.
func TestPromptForScannerOptions(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		startPort string
		endPort   string
	}{
		{
			name:      "test prompt for single port",
			host:      "localhost",
			startPort: "80",
			endPort:   "",
		},
		{
			name:      "test prompt for multiple ports",
			host:      "localhost",
			startPort: "80",
			endPort:   "100",
		},
		{
			name:      "test prompt for swapped ports",
			host:      "localhost",
			startPort: "100",
			endPort:   "80",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := fmt.Sprintf(
				"%s\n%s\n%s\n\n",
				tt.host, tt.startPort, tt.endPort,
			)
			reader := strings.NewReader(input)

			host, ports := PromptForScannerOptions(reader)

			if host != tt.host {
				t.Errorf("PromptForScannerOptions() got host:%v, wanted host:%v", host, tt.host)
			}

			var portsToScan []int
			startPort, _ := strconv.Atoi(tt.startPort)
			if tt.endPort == "" {
				portsToScan = append(portsToScan, startPort)
			} else {
				endPort, _ := strconv.Atoi(tt.endPort)
				if startPort < endPort {
					for i := startPort; i <= endPort; i++ {
						portsToScan = append(portsToScan, i)
					}
				} else {
					for i := endPort; i <= startPort; i++ {
						portsToScan = append(portsToScan, i)
					}
				}
			}

			if !reflect.DeepEqual(ports, portsToScan) {
				t.Errorf("PromptForScannerOptions() got ports:%v, wanted ports:%v", ports, portsToScan)
			}

			t.Log("host:", host)
			t.Log("ports to scan:", ports)
		})
	}
}
