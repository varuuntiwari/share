package netscan

import (
	"net"
	"strconv"
	"strings"
)

// Pre-defined IP ranges with mask which are accepted for use in local networks only
var (
	validRanges = []string{"192.168.0.0/16", "10.0.0.0/8", "172.16.0.0/12"}
)

// IPinRange checks if the IP detected is a local address or global address
func IPinRange(ip net.IP) {
	if ok, IPRange := localRangeCheck(ip); ok {
		scanRange(IPRange)
	}
}

// localRangeCheck checks if target IP exists in given range of IPs
func localRangeCheck(targetIP net.IP) (bool, string) {
	for _, x := range validRanges {
		t := strings.Split(x, "/")
		ip, m := net.ParseIP(t[0]), t[1]
		prefix, _ := strconv.Atoi(m)
		mask := net.CIDRMask(prefix, 32)
		network := net.IPNet{
			IP: ip,
			Mask: mask,
		}
		if network.Contains(targetIP) {
			return true, x
		}
	}
	return false, ""
}