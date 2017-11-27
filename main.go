package main

import (
	"flag"
	"fmt"
	"github.com/jm33-m0/xmir/xmirlib"
	"log"
	"os"
)

func main() {
	serverType := flag.String("server", "", "Specify a server software")
	masscanFile := flag.String("xml", "", "Masscan result file in XML format")
	outfile := flag.String("tolist", "", "IP list file for XML parser output")
	result := flag.String("result", "", "Where to save our scan result")
	flag.Parse()

	if flag.NFlag() != 4 {
		flag.Usage()
		fmt.Println("eg. ./xmir -xml masscan.xml -tolist " +
			"list.txt -server Joomla -result result.txt")
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

	// if ip list doesn't exist, parse the XML file to get one
	if _, err := os.Stat(*outfile); os.IsNotExist(err) {
		log.Println("Parsing masscan result...")
		xmirlib.XML2List(*masscanFile, *outfile, filter)
	}

	log.Println("Analysing targets...")
	xmirlib.ScanList(*outfile, *result, *serverType)
}
