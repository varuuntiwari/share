package main

import (
	"fmt"

	"github.com/varuuntiwari/share/recv"
	"github.com/varuuntiwari/share/send"
)

func welcomeScreen() {
	fmt.Print(`
▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
██░▄▄▄░█░████░▄▄▀█░▄▄▀█░▄▄
██▄▄▄▀▀█░▄▄░█░▀▀░█░▀▀▄█░▄▄
██░▀▀▀░█▄██▄█▄██▄█▄█▄▄█▄▄▄
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
 by varuuntiwari`)
 	fmt.Println(`
 	1. Send
 	2. Receive`)
}

var (
	opt int64
)

func main() {
	welcomeScreen()
	fmt.Print(": ")
	fmt.Scanf("%d", &opt)

	if opt == 1 {
		send.SendFileCheck()
	} else if opt == 2 {
		recv.ReceiveFileCheck()
	}
}