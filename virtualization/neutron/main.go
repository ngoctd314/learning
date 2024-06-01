package main

import "fmt"

type host struct {
	mac  int
	port int
}

type lan struct {
	sw    l2Switch
	hosts []host
}

func (l *lan) addHost(h host, port int) error {
	if _, ok := l.sw.portBlackList[port]; ok {
		return fmt.Errorf("port %d is inused", port)
	}
	if _, ok := l.sw.macToPort[h.mac]; ok {
		return fmt.Errorf("mac %d is inused", h.mac)
	}

	l.hosts = append(l.hosts, h)
	l.sw.macToPort[h.mac] = port
	l.sw.portBlackList[port] = struct{}{}

	return nil
}

func (l *lan) canConnect(h1 host, h2 host) bool {

	return true
}

type l2Switch struct {
	macToPort     map[int]int // map MAC address to port
	portBlackList map[int]struct{}
}

func main() {
}
