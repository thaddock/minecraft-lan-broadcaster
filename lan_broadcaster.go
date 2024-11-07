package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Hello there!\n")

	conn, err := net.Dial("udp", "224.0.2.60:4445")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("[MOTD]They might be gnats[/MOTD][AD]22556[/AD]"))
  	if err != nil {
    		panic(err)
  	}
}
