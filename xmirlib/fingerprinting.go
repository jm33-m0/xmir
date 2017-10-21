package xmirlib

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// IsJoomla : Does this server run Joomla?
func IsJoomla(host string, port int) bool {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("http://" + host + ":" + strconv.Itoa(port))
	if err != nil {
		log.Print(err)
	}
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
