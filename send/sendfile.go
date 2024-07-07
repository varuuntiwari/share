package send

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/varuuntiwari/share/vars"
)

var (
	portStr = strconv.Itoa(vars.Port)
)

// sendFile sends a test string to the selected IP address which is stored in
// TargetIP in vars/vars.go. This is a temporary function and will be replaced
// by the actual file sending function in the next release.
func sendFile() {
	target := vars.TargetIP + ":" + portStr
	conn, err := net.Dial("tcp", target)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Read file and dump into byte slice
	data, err := os.ReadFile("./screenshots/poc_on_wsl.png")
	if err != nil {
		panic(err)
	}
	
	// Write the data to the connection
	n, err := conn.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`sent %v bytes to %v\n`, n, vars.TargetIP)
}