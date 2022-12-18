package scanner

import "testing"

// TestScanPorts tests the CheckOpenPorts function.
func TestScanPorts(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		portsToScan []int
	}{
		{
			name:        "test scanning localhost ports",
			host:        "localhost",
			portsToScan: []int{80, 443, 22},
		},
		{
			name:        "test scanning scanme ports",
			host:        "scanme.nmap.org",
			portsToScan: []int{80, 443, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			openPorts := CheckOpenPorts(tt.host, tt.portsToScan)
			t.Logf("open %s ports: %v", tt.host, openPorts)
		})
	}
}
