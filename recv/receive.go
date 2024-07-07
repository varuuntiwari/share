// Package recv contains the functions uesd to listen for any senders and receive
// the data along with handling any write functions to the disk.
package recv

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/varuuntiwari/share/vars"
)

// ReceiveFileCheck sets up a listener to check for any senders and subsequently
// receive data from it once the check is complete, printing out the data. It uses
// the port specified in vars/vars.go to listen on and uses a buffer of 1024 bytes
// which is used to read the text sent.
// It returns nothing as the function is complete in itself from start to finish.
func ReceiveFileCheck() {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(vars.Port))
	if err != nil {
		panic(err)
	}

	fmt.Println("Stand-by for sender to scan...")
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	buf, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s ready to send data...\n", string(buf))

	conn, err = ln.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Println("Receiving data...")

	file, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("received_file.png", file, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("File received and saved as received_file.png.")
	conn.Close()
}