package main

import (
	"flag"
	"github.com/jm33-m0/xmir/xmirlib"
	"log"
)

func main() {
	serverType := flag.String("server", "", "Specify a server software")
	masscanFile := flag.String("file", "", "Masscan result file in XML format")
	outfile := flag.String("outfile", "", "IP list file for XML parser output")
	result := flag.String("result", "", "Scan result")
	flag.Parse()

	if *masscanFile == "" || *serverType == "" {
		flag.Usage()
		return
	}

	var filter string
	switch *serverType {
	case "Joomla":
		log.Println("Using Joomla")
		filter = ""
	case "WordPress":
		log.Println("Using WordPress")
		filter = ""
	default:
		log.Fatal("Not supported")
	}

	log.Println("[*] xmir started")

	log.Println("Parsing masscan result...")
	xmirlib.XML2List(*masscanFile, *outfile, filter)

	log.Println("Analysing targets...")
	xmirlib.ScanList(*outfile, *serverType, *result)
}
