package main

import (
	"github.com/jm33-m0/xmir/xmirlib"
	"log"
)

func main() {
	log.Println("let's begin")
	xmirlib.ScanList("iplist.txt", "Joomla")
}
