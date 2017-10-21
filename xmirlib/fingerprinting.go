package xmirlib

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// IsJoomla : Does this server run Joomla?
func IsJoomla(host string, port int) bool {
	resp, err := http.Get("http://" + host + ":" + strconv.Itoa(port))
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
