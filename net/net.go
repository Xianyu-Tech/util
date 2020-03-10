package netutil

import (
	"net"
)

func GetNetIps() ([]string, error) {
	intrAddrs, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}

	netIps := make([]string, 0)

	for _, intrAddr := range intrAddrs {
		if ipNet, ok := intrAddr.(*net.IPNet); ok == true {
			if ipNet.IP.IsLoopback() == false {
				if ipNet.IP.To4() != nil {
					netIps = append(netIps, ipNet.IP.String())
				}
			}
		}
	}

	return netIps, nil
}
