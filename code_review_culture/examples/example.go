package main

func main() {
	ips := getIPs()

	var processedIPs []ipv6
	for _, ip := range ips {
		ipv6 := ipv6From4(ip)
		processedIPs = append(processedIPs, ipv6)
	}
}

type ipv6 struct{}

type ipv4 struct{}

func getIPs() []ipv4 {
	return nil
}

func ipv6From4(ips ipv4) ipv6 {
	return ipv6{}
}
