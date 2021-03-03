package toolkit

import "testing"

var (
    s1 = "hello world"
    s2 = "0.0.0.0"
    s3 = "1.1.1.256"
    s4 = "0xC0.0x00.0x02.0xEB"
    s5 = "2001:db8:2de:0:0:0:0:e13"
    s6 = "2001:0db8::1428:57ab"
    s7 = "2001::25de::cade"
)

func TestIsIPv4(t *testing.T) {
    t.Log(IsIPv4(s1))
    t.Log(IsIPv4(s2))
    t.Log(IsIPv4(s3))
    t.Log(IsIPv4(s4))
    t.Log(IsIPv4(s5))
    t.Log(IsIPv4(s6))
    t.Log(IsIPv4(s7))
}

func TestIsIPv6(t *testing.T) {
    t.Log(IsIPv6(s1))
    t.Log(IsIPv6(s2))
    t.Log(IsIPv6(s3))
    t.Log(IsIPv6(s4))
    t.Log(IsIPv6(s5))
    t.Log(IsIPv6(s6))
    t.Log(IsIPv6(s7))
}

func TestIsIP(t *testing.T) {
    t.Log(IsIP(s1))
    t.Log(IsIP(s2))
    t.Log(IsIP(s3))
    t.Log(IsIP(s4))
    t.Log(IsIP(s5))
    t.Log(IsIP(s6))
    t.Log(IsIP(s7))
}
