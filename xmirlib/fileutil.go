package xmirlib

import (
	"bufio"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

// Address : host>address
type Address struct {
	Addr     string `xml:"addr,attr"`
	Addrtype string `xml:"addrtype,attr"`
}

// State : host>ports>port>state
type State struct {
	State     string `xml:"state,attr"`
	Reason    string `xml:"reason,attr"`
	ReasonTTL string `xml:"reason_ttl,attr"`
}

// Service : host>ports>port>service
type Service struct {
	Name   string `xml:"name,attr"`
	Banner string `xml:"banner,attr"`
}

// Ports : host>ports
type Ports []struct {
	Protocol string `xml:"protocol,attr"`
	Portid   string `xml:"portid,attr"`

	State   State   `xml:"state"`
	Service Service `xml:"service"`
}

// Host : host field in XML
type Host struct {
	XMLName xml.Name `xml:"host"`
	Endtime string   `xml:"endtime,attr"`

	Address Address `xml:"address"`
	Ports   Ports   `xml:"ports>port"`
}

// XML2List : Parse masscan result, pick useful items and save them to a list file
func XML2List(xmlfile string, outfile string, filter string) {

	xmlStream, err := os.Open(xmlfile)
	if err != nil {
		log.Println("Failed to open XML file")
		return
	}
	defer xmlStream.Close()

	decoder := xml.NewDecoder(xmlStream)

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "host" {
				var h Host
				decoder.DecodeElement(&h, &se)

				// since mostly we have just one port to detect
				address := h.Address.Addr
				// port := h.Ports[0].Portid

				banner := h.Ports[0].Service.Banner
				// test value reading
				// log.Print(address, " : ", port, "\n", banner)

				// write desired host to file
				if searchHost(filter, banner) {
					AppendToFile(outfile, address)
				}
			}
		default:
		}
	}
}

// AppendToFile : append a line to target file
func AppendToFile(filepath string, line string) {

	// open outfile
	out, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(filepath, " : Failed to open file\n", err)
		return
	}

	// write appendly
	if _, err = out.Write([]byte(line + "\n")); err != nil {
		log.Print(filepath, " : Write err: ", err, "\nWriting ", line)
		return
	}
	out.Close()
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

func searchHost(filter string, banner string) bool {
	if strings.Contains(banner, filter) ||
		filter == "" {
		return true
	}

	return false
}
