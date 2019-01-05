package common

import (
	"fmt"
	"net"
)

func ByteFormat(length uint64) string {
	size := float64(length)
	units := []string{"B", "K", "M", "G", "Conf"}

	for _, unit := range units {
		if size < 1024 {
			return fmt.Sprintf("%7.2f%s", size, unit)
		}
		size = size / 1024.0
	}
	return "Size larger than 1024TB."
}

func In(num byte, list []byte) bool {
	for _, e := range list {
		if e == num {
			return true
		}
	}
	return false
}

var ipAddress string

func GetIp() string {
	if ipAddress == "" {
		interfaces, _ := net.Interfaces()
		for _, i := range interfaces {
			address, _ := i.Addrs()
			for _, addr := range address {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				ipAddress = ip.String()
			}
		}
	}
	return ipAddress
}
