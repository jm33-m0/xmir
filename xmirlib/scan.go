package xmirlib

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

// ScanList : Scan IPs from a text file
func ScanList(filepath string, serverType string) {
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
			runFingerprinting(host, serverType)
			wg.Done()
		}()
		i++
		if i == 0 || i == 100 {
			i = 0
			wg.Wait()
		}
	}
}

func runFingerprinting(target string, serverType string) {
	switch serverType {
	case "Joomla":
		if IsJoomla(target, 80) ||
			IsJoomla(target, 443) {
			log.Print("[+] Joomla found! ", target)
		}
	case "WordPress":
		if IsWordPress(target, 80) ||
			IsWordPress(target, 443) {
			log.Print("[+] WordPress found! ", target)
		}
	default:
		log.Println("Please specify a server type, eg. Joomla")
	}
}

// FileToLines : Read lines from a text file
func FileToLines(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err == nil {
		defer f.Close()

		var lines []string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}
		return lines, nil
	}
	return nil, err
}
