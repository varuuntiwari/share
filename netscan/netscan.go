// Package netscan is used along with the send package to scan the local network
// and find hosts ready to receive the file. It also sends the test string for now.
// This functionality will be moved back the send package in the next release to maintain
// the modularity of the packages.
package netscan

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/varuuntiwari/share/vars"
)

// var portStr is declared here to avoid multiple conversions from int to string
// in various places in this code.
var (
	portStr = strconv.Itoa(vars.Port)
)

// nextIP returns the next IP address from the IP passed to it.
func nextIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// ScanRange retrieves the local CIDR to find the IP range to scan for open ports.
// After finding a valid IP range, it will scan all IPs in a multi-threaded approach
// and add the IPs which respond to the variable ConnectedHosts in vars/vars.go.
func ScanRange(ipRange net.IPNet) bool {
	var wg sync.WaitGroup
	var mu sync.Mutex

	localCIDR := GetLocalCIDR()
	ip, ipNet, err := net.ParseCIDR(localCIDR.String())
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		panic(err)
	}

	// Iterate over the IP addresses within the CIDR range
	fmt.Printf("Scanning for open ports over the IP range %v...\n", localCIDR.String())
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); nextIP(ip) {
		target := ip.String() + ":" + portStr
		sendCheck := []byte(strings.Split(GetLocalCIDR().String(), "/")[0])

		// Multi-threaded approach to scan for open ports
		wg.Add(1)
		go pingIP(target, &wg, &mu, sendCheck)
	}
	wg.Wait()
	return true
}

func pingIP(ip string, wg *sync.WaitGroup, mu *sync.Mutex, sendCheck []byte) {
	defer wg.Done()
	conn, err := net.DialTimeout("tcp", ip, time.Second * 5)

	if err == nil {
		_, err := conn.Write(sendCheck)
		conn.Close()
		mu.Lock()
		if err == nil {
			vars.ConnectedHosts = append(vars.ConnectedHosts, strings.Split(ip, ":")[0])
		} else {
			fmt.Println("[-] Error sending data to IP: ", strings.Split(ip, ":")[0])
		}
		defer mu.Unlock()
	}
}