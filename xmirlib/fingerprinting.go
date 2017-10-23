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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	targetURL := "http://" + host + ":" + strconv.Itoa(port) + "/"
	log.Print(targetURL)
	resp, err := client.Get(targetURL)
	if err != nil {
		log.Print(err)
		return false
	}
	resp.Close = true
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	if strings.Contains(string(respBody), `<meta name="generator" content="Joomla! 1.5 - Open Source Content Management" />`) {
		return true
	}
	return false
}
