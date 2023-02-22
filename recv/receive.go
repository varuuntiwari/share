package recv

import (
	"fmt"
	"net"
	"strconv"

	"github.com/varuuntiwari/share/vars"
)

func ReceiveFileCheck() {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(vars.Port))
	if err != nil {
		panic(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received data: %s\n", string(buf[:n]))
}