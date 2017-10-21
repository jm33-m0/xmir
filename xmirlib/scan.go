package xmirlib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ScanList : Scan IPs from a text file
func ScanList(filepath string) {
	lines, err := FileToLines(filepath)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		host := strings.Trim(line, "\n")
		if IsJoomla(host, 80) {
			log.Print("Found one! ", host)
		}
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
