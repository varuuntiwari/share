// Package send contains the functions used to send data to the receiver.
// It sends the control to the netscan package after extracting a valid IP address
// to verify if it is on a valid IP range, which should be a private network. These
// ranges are desribed in the netscan package.
package send

import (
	"fmt"
	"net"

	"github.com/varuuntiwari/share/netscan"
	"github.com/varuuntiwari/share/vars"
)

var choice int

// SendFileCheck detects a local network to send the file in and scan for hosts
// ready to receive the file.
func SendFileCheck() {
	// get all interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(interfaces); i++ {
		fmt.Printf("%v. %v\n", (i + 1), interfaces[i].Name)
	}
	fmt.Print("Select interface to find receiver on: ")
	fmt.Scanf("%v")
	fmt.Scanf("%d", &choice)

	// get interface selected
	iface, _ := net.InterfaceByName(interfaces[(choice - 1)].Name)
	vars.SelectedInterface = iface.Name
	// get all addresses from interface
	addrs, _ := iface.Addrs()
	for _, addr := range addrs {
		// if address is IPv4, return it to scan network
		if ip := addr.(*net.IPNet).IP.To4(); ip != nil {
			exists, IPRange := netscan.IPinRange(ip)
			if exists && netscan.ScanRange(IPRange) {
				fmt.Print("\033[32m")
				fmt.Println("IPs open for connection:")
				fmt.Print("\033[0m")
				for i := 0; i < len(vars.ConnectedHosts); i++ {
					fmt.Printf("%d. %v\n", i+1, vars.ConnectedHosts[i])
				}
				var choice int
				fmt.Print("Select IP to send data to: ")
				fmt.Scanf("%v")
				fmt.Scanf("%d", &choice)

				vars.TargetIP = vars.ConnectedHosts[choice-1]
				sendFile()
			}
		}
	}

}
