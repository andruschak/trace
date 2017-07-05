package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n  sudo ./traceroute <domain|ip>\n")
		flag.PrintDefaults()
	}
}
