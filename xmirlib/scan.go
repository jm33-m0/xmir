package xmirlib

import (
	"log"
	"strings"
	"sync"
)

// ScanList : Scan IPs from a text file
func ScanList(filepath string, outfile string, serverType string) {
	lines, err := FileToLines(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	i := 0
	for _, line := range lines {
		host := strings.Trim(line, "\n")
		go func() {
			wg.Add(1)
			runFingerprinting(host, serverType, outfile)
			wg.Done()
		}()
		i++
		if i == 1000 {
			i = 0
			wg.Wait()
		}
	}
}

func runFingerprinting(target string, serverType string, outfile string) {
	switch serverType {
	case "Joomla":
		if IsJoomla(target, 80) ||
			IsJoomla(target, 443) {
			log.Print("[+] Joomla found! ", target)
			AppendToFile(outfile, target)
		}
	case "WordPress":
		if IsWordPress(target, 80) ||
			IsWordPress(target, 443) {
			log.Print("[+] WordPress found! ", target)
			AppendToFile(outfile, target)
		}
	default:
		log.Println("Please specify a server type, eg. Joomla")
	}
}
