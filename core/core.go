package core

import (
	"fmt"
	"net"
	"time"
)

func Scan(ip string, portRange int, workerRange int) {
	ports := make(chan int)

	for worker := 0; worker < workerRange; worker++ {
		go func() {
			for port := range ports {
				isOpen(port, ip)
			}
		}()
	}

	for port := 0; port < portRange; port++ {
		ports <- port
	}
}

func isOpen(port int, ip string) {
	port += 1
	address := fmt.Sprintf("%s:%d", ip, port)
	timeOut := time.Millisecond * 250

	server, err := net.DialTimeout("tcp", address, timeOut)

	if err == nil {
		fmt.Printf("Port:%d is Open\n", port)
		_ = server.Close()
	}
}
