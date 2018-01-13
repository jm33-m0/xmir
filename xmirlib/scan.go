package xmirlib

import (
	"log"
	"os"
	"strings"
	"sync"
)

// ScanList : Scan IPs from a text file
func ScanList(filepath string, outfile string, serverType string) {
	lines, err := FileToLines(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// open outfile
	outf, err := OpenFileStream(outfile)
	if err != nil {
		log.Println("Error opening ", outfile+"\n", err)
		return
	}
	defer CloseFileStream(outf)

	var wg sync.WaitGroup
	i := 0
	for _, line := range lines {
		host := strings.Trim(line, "\n")
		go func() {
			wg.Add(1)
			runFingerprinting(host, serverType, outf)
			wg.Done()
		}()
		i++
		if i == 200 && &wg != nil {
			i = 0
			wg.Wait()
		}
	}
}

func runFingerprinting(target string, serverType string, outf *os.File) {
	switch serverType {
	case "Joomla":
		if IsJoomla(target, 80) ||
			IsJoomla(target, 443) {
			log.Print("[+] Joomla found! ", target)
			AppendToFile(outf, target)
		}
	case "WordPress":
		if IsWordPress(target, 80) ||
			IsWordPress(target, 443) {
			log.Print("[+] WordPress found! ", target)
			AppendToFile(outf, target)
		}
	default:
		log.Println("Please specify a server type, eg. Joomla")
	}
}
