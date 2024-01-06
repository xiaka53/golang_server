package tool

import (
	"net"
	"time"
)

// ping域名
func PingUrl(domain string) bool {
	conn, err := net.DialTimeout("tcp", domain+":80", 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
