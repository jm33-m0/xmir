package main

import (
	"flag"
	"github.com/jm33-m0/xmir/xmirlib"
	"log"
)

func main() {
	serverType := flag.String("server", "WordPress", "Specify a server software")
	listFile := flag.String("list", "", "IP list file")
	scanAll := flag.Bool("all", false, "Whether to search for all available server software")
	flag.Parse()

	if *listFile == "" {
		flag.Usage()
		return
	}

	log.Println("[*] Scanner started")

	serverTypes := []string{"Joomla", "WordPress"}

	if *scanAll && *serverType == "" {
		for _, item := range serverTypes {
			xmirlib.ScanList(*listFile, item)
		}
		return
	}
	xmirlib.ScanList(*listFile, *serverType)
}
