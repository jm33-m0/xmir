package xmirlib

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// IsJoomla : Does this server run Joomla?
func IsJoomla(host string, port int) bool {

	targetURL := "http://" + host + ":" + strconv.Itoa(port) + "/"
	if port == 443 {
		targetURL = "https://" + host + ":" + strconv.Itoa(port) + "/"
	}
	respBody := getPage(targetURL)
	if respBody == nil {
		return false
	}

	if strings.Contains(string(respBody), `<meta name="generator" content="Joomla! 1.5 - Open Source Content Management" />`) {
		return true
	}
	return false
}

// Get webpage from target http server for further analysis
func getPage(targetURL string) []byte {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		log.Print(err)
		return nil
	}
	resp.Close = true
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	return respBody
}
