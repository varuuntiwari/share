package send

import (
	"fmt"
	"net"

	"github.com/varuuntiwari/share/netscan"
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
	// get all addresses from interface
	addrs, _ := iface.Addrs()
	for _, addr := range addrs {
		// if address is IPv4, return it to scan network
		if ip := addr.(*net.IPNet).IP.To4(); ip != nil {
			netscan.IPinRange(ip)
		}
	}
}