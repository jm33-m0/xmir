package xmirlib

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// IsJoomla : Does this server run Joomla?
func IsJoomla(host string, port int) bool {

	respBody := getPage(host, port)
	if respBody == nil {
		return false
	}

	if strings.Contains(string(respBody), `<meta name="generator" content="Joomla!`) {
		return true
	}
	return false
}

// IsWordPress : Does this server run WordPress?
func IsWordPress(host string, port int) bool {

	respBody := getPage(host, port)
	if respBody == nil {
		return false
	}

	if strings.Contains(string(respBody), `<meta name="generator" content="WordPress`) {
		return true
	}
	return false
}

// Get webpage from target http server for further analysis
func getPage(host string, port int) []byte {

	targetURL := "http://" + host + ":" + strconv.Itoa(port) + "/"
	if port == 443 {
		targetURL = "https://" + host + ":" + strconv.Itoa(port) + "/"
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		// log.Print(err)
		return nil
	}
	resp.Close = true
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Print(err)
		return nil
	}
	return respBody
}
