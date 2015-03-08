package gutils

import (
	"fmt"
	"net"
	"time"
)

type UDPQuitFunc func()

// from address and content
type UDPProcessFunc func(net.Addr, string)

/*
listen on a udp socket, quit when reading in ":quit\n" and call quit
/*func, pass anything else to the processing function
*/
func UdpServe(addr string, qf UDPQuitFunc, pf UDPProcessFunc) error {
	udpAddress, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		fmt.Println("error resolving UDP address on ", addr)
		fmt.Println(err)
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		fmt.Println("error listening on UDP address ", addr)
		fmt.Println(err)
		return err
	}

	defer conn.Close()

	var buf [1024]byte
	for {
		time.Sleep(100 * time.Millisecond)
		n, address, err := conn.ReadFrom(buf[:])
		if err != nil {
			fmt.Println("error reading data from connection")
			fmt.Println(err)
			continue
		}
		if address != nil {
			fmt.Println("got message from ", address, " with n = ", n)
			if n > 0 {
				content := string(buf[0:n])
				fmt.Println("from address ", address, " got message: ", content, n)
				if content == ":quit\n" {
					qf()
					break
				} else {
					pf(address, content)
				}
			}
		}
	}
	return nil
}
