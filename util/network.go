package util

import (
	"fmt"
	"net"
	"strings"
)

// GetHostIp get host ip from udp telnet
func GetHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	addr := conn.LocalAddr().String()
	ip := strings.Split(addr, ":")[0]
	return ip
}
