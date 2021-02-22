package toolkit

import (
	"fmt"
	"time"
)

func nowInCST() time.Time {
	return time.Now().In(time.FixedZone("CST", 8*3600))
}

// NowCST return now time for CST zone in normal string
func NowCST() string {
	return nowInCST().Format("2006-01-02 15:04:05")
}

// NowDateCST return now date for CST zone in normal string
func NowDateCST() string {
	return nowInCST().Format("2006-01-02")
}

// NowCSTUnix return now time for CST zone in unix string
func NowCSTUnix() string {
	return fmt.Sprintf("%d", nowInCST().Unix())
}
