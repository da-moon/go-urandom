package urandom

import (
	"net"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	random := make([]rune, n)
	for i := range random {
		random[i] = letterRunes[seed.Intn(len(letterRunes))]
	}
	return string(random)
}

// IpV4Address returns a valid IPv4 address as string
func IpV4Address() string {
	return ipAddr(net.IPv4len)
}

// IpV6Address returns a valid IPv6 address as net.IP
func IpV6Address() string {
	return ipAddr(net.IPv6len)
}
func ipAddr(lenght int) string {
	var ip net.IP
	for i := 0; i < length; i++ {
		number := uint8(seed.Intn(255))
		ip = append(ip, number)
	}
	return ip.String()
}
