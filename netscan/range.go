package netscan

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/varuuntiwari/share/vars"
)

// Pre-defined IP ranges with mask which are accepted for use in local networks only
var (
	validRanges = []string{"192.168.0.0/16", "10.0.0.0/8", "172.16.0.0/12"}
)

// IPinRange checks if the IP detected is a local address or global address
func IPinRange(ip net.IP) {
	if ok, IPRange := localRangeCheck(ip); ok {
		scanRange(IPRange)
	}
}

// localRangeCheck checks if target IP exists in given range of IPs
func localRangeCheck(targetIP net.IP) (bool, net.IPNet) {
	for _, x := range validRanges {
		t := strings.Split(x, "/")
		ip, m := net.ParseIP(t[0]), t[1]
		prefix, _ := strconv.Atoi(m)
		mask := net.CIDRMask(prefix, 32)
		network := net.IPNet{
			IP: ip,
			Mask: mask,
		}
		if network.Contains(targetIP) {
			return true, network
		}
	}
	return false, net.IPNet{}
}


// Get the IP of local system on the selected network interface
func GetLocalCIDR() net.Addr {
	// Get the network interface by name
	localInterface, err := net.InterfaceByName(vars.SelectedInterface)
	if err != nil {
		fmt.Println("Error getting interface:", err)
		return nil
	}

	// Get the addresses associated with the network interface
	addrs, err := localInterface.Addrs()
	if err != nil {
		fmt.Println("Error getting addresses:", err)
		return nil
	}

	// Print the IP addresses
	for _, addr := range addrs {
		if ip := addr.(*net.IPNet).IP.To4(); ip != nil {
			return addr
		}
	}
	return nil
}