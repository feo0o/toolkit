package toolkit

import "net"

// IsIPv4 check the string s if a valid IPv4, only for dotted decimal format
func IsIPv4(s string) bool {
    ip := net.ParseIP(s)
    if ip == nil {
        return false
    }
    ip = ip.To4()
    if ip == nil {
        return false
    }
    return true
}

// IsIPv6 check the string s if a valid IPv6 address, only for dotted hexadecimal format
func IsIPv6(s string) bool {
    ip := net.ParseIP(s)
    if ip == nil {
        return false
    }
    ipv4 := ip.To4()
    if ipv4 != nil {
        return false
    }
    ipv6 := ip.To16()
    if ipv6 == nil {
        return false
    }
    return true
}

// IsIP check the string s if a valid IP address
// support dotted decimal format for IPv4 and dotted hexadecimal format for IPv6
func IsIP(s string) bool {
    return IsIPv4(s) || IsIPv6(s)
}
