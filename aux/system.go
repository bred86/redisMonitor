package aux

import (
	"net"
	"os"
)

// GetHostname - (string)
func GetHostname() string {
	hostname, errorHostname := os.Hostname()
	if errorHostname != nil {
		panic(errorHostname)
	}

	return hostname
}

// GetLocalIP - (string)
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
