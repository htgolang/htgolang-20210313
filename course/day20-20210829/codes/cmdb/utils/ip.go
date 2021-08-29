package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PortRange struct {
	startPort int
	endPort   int
}

func (portRange *PortRange) Start() int {
	return portRange.startPort
}

func (portRange *PortRange) End() int {
	return portRange.endPort
}

func (portRange *PortRange) String() string {
	return fmt.Sprintf("%d-%d", portRange.startPort, portRange.endPort)
}

// IPV4
type IPRange struct {
	startIP string
	endIP   string
}

func (ipRange *IPRange) Start() int64 {
	return IP2Int(ipRange.startIP)
}

func (ipRange *IPRange) End() int64 {
	return IP2Int(ipRange.endIP)
}

func (ipRange *IPRange) String() string {
	return fmt.Sprintf("%s-%s", ipRange.startIP, ipRange.endIP)
}

func ParseIPRange(text string) []*IPRange {
	pattern := regexp.MustCompile(`[;,\s]+`)
	addrs := pattern.Split(text, -1)
	ipRanges := make([]*IPRange, 0, len(addrs))
	for _, addr := range addrs {
		addr := strings.TrimSpace(addr)
		if addr == "" {
			continue
		}
		if ipRange, err := parseIPAddr(addr); err == nil {
			ipRanges = append(ipRanges, ipRange)
		}
	}
	return ipRanges
}

func parseIPAddr(addr string) (*IPRange, error) {
	// start end:-
	// cidr: /
	// ip
	ipRange := new(IPRange)
	if strings.Contains(addr, "-") {
		nodes := strings.SplitN(addr, "-", 2)
		ipRange.startIP = nodes[0]
		ipRange.endIP = nodes[1]
	} else if strings.Contains(addr, "/") {
		nodes := strings.SplitN(addr, "/", 2)
		ipNum := IP2Int(nodes[0])
		mask, _ := strconv.Atoi(nodes[1])
		ipNum >>= 32 - mask

		ipRange.startIP = Int2Ip(ipNum << (32 - mask))

		for i := 0; i < 32-mask; i++ {
			ipNum <<= 1
			ipNum |= 1
		}
		ipRange.endIP = Int2Ip(ipNum)

		// 192.168.0.1/24
		// 256 => 8
		// [0000 0000].[0000 0000].[0000 0000].[0000 0000]
		// 32-24 0000 0000 1111 1111
		// ip int >> 32-mask << 32-mask
		// ip int >> 32-mask
		// 32-mask << 1 | 1

	} else {
		ipRange.startIP = addr
		ipRange.endIP = addr
	}

	return ipRange, nil
}

func ParsePorts(text string) []*PortRange {
	// start-end
	pattern := regexp.MustCompile(`[;,\s]+`)
	nodes := pattern.Split(text, -1)
	portRanges := make([]*PortRange, 0, len(nodes))
	for _, node := range nodes {
		node = strings.TrimSpace(node)
		if node == "" {
			continue
		}
		if portRange, err := parsePort(node); err == nil {
			portRanges = append(portRanges, portRange)
		}
	}
	return portRanges
}

func parsePort(port string) (*PortRange, error) {
	portRange := new(PortRange)
	if strings.Contains(port, "-") {
		nodes := strings.Split(port, "-")
		start, _ := strconv.Atoi(nodes[0])
		end, _ := strconv.Atoi(nodes[1])
		portRange.startPort = start
		portRange.endPort = end
	} else {
		portNum, _ := strconv.Atoi(port)
		portRange.startPort = portNum
		portRange.endPort = portNum
	}
	return portRange, nil
}

func IP2Int(addr string) int64 {
	// 1.2.3.4
	// 1234(256) 0-255
	// 1234(10) 0-9
	// 1 * 256 ^ 3 + 2 * 256 ^ 2 + 3 * 256 ^ 1 + 4 * 256 ^ 0
	// 1 * 10 ^ 3 + 2 * 10 ^ 2 + 3 * 10 ^ 1 + 4 * 10 ^ 0
	// num = 0 * 256 + 1
	// num = 1 * 256 + 2
	// num = (1 * 256 + 2) * 256 + 3 => 1 * 256 ^ 2 + 2 * 256 + 3
	// num = (1 * 256 ^ 2 + 2 * 256 + 3) *256 + 4 => 1 * 256 ^3 + 2 * 256 ^2 + 3 * 256 + 4
	nodes := strings.Split(addr, ".")
	var num int64
	for _, node := range nodes {
		temp, _ := strconv.Atoi(node)
		num = num*256 + int64(temp)
	}
	return num
}

func Int2Ip(num int64) string {
	count := 4
	var base int64 = 256
	nodes := make([]string, count)
	for i := 0; i < count; i++ {
		nodes[count-i-1] = strconv.FormatInt(num%base, 10)
		num /= base
	}
	return strings.Join(nodes, ".")
}
