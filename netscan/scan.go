package netscan

import (
	"net"
	"strconv"

	"github.com/varuuntiwari/share/vars"
)

// Scan the IP range to check for open ports in the devices present on the local network
func scanRange(ipRange string) {
	vars.TargetIP = "localhost"
	sendFile()
}

// Send the file
func sendFile() {
	target := vars.TargetIP + ":" + strconv.Itoa(vars.Port)
	conn, err := net.Dial("tcp", target)
	if err != nil {
		panic(err)
	}

	data := []byte("my own data, Varun")

	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}
}