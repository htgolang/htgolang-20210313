package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("127.0.0.1", "8888"))
	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupAddr("204.79.197.200"))
	fmt.Println(net.LookupHost("localhost"))
	fmt.Println(net.LookupHost("www.baidu.com"))

	// ip=>IP
	// 字符串 ipv4 点分十进制
	// ipv6 ::
	// net.ParseIP()
	var ip net.IP = net.ParseIP("1.1.1.1")
	fmt.Println(ip)
	ip = net.ParseIP("::1")
	fmt.Println(ip)
	ip = net.ParseIP("x")
	fmt.Println(ip)

	// ip段 cidr格式 ip/mask
	// ipstart-ipend 192.168.0.2-192.168.0.255
	// 192.168.0.2-255
	ip, ipnet, err := net.ParseCIDR("192.168.0.0/24")
	fmt.Println(ip)
	fmt.Println(ipnet)
	fmt.Println(err)

	ip, ipnet, err = net.ParseCIDR("192.168.0.10/32")
	fmt.Println(ip)
	fmt.Println(ipnet)
	fmt.Println(err)

	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.10")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.11")))

	ip, ipnet, err = net.ParseCIDR("192.168.0.0/24")
	// 192.168.0.0 -192.168.0.255
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.10")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.0.11")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.11")))

	// Addr 网络地址
	addrs, err := net.InterfaceAddrs()

	fmt.Println("addrs:")
	for idx, addr := range addrs {
		fmt.Println(idx, addr.Network(), addr.String())
	}

	fmt.Println("interfaces:")
	// Interfaces 网卡地址
	intfs, err := net.Interfaces()

	for idx, intf := range intfs {
		fmt.Println(idx,
			intf.Index,
			intf.Name,
			intf.MTU,
			intf.Flags,
			intf.HardwareAddr.String(),
		)
		fmt.Println(intf.Addrs())
		fmt.Println(intf.MulticastAddrs())
	}
}
