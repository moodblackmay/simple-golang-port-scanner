package main

import (
	"flag"
	"log"
	"scan/core"
)

func main() {
	ip := flag.String("ip", "", "ipv4 address")
	portRange := flag.Int("port", 1000, "port range")
	workerRange := flag.Int("T", 1, "threads")

	flag.Parse()

	portCheck := *portRange >= 0 && *portRange < 65535
	workerCheck := *workerRange >= 0 && *workerRange < 6

	if !workerCheck {
		log.Panicln("T parameter must be between 1-5")
	}

	if !portCheck {
		log.Panicln("port parameter must be between 1-65535")
	}

	core.Scan(*ip, *portRange, *workerRange)
}
