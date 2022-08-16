package main

import (
	"log"
	"net/http"
	"testing"
)

func Test_parseCountryTable(t *testing.T) {
	res, err := http.Get("https://whois.ipip.net/countries")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	parseCountryTable(res.Body)
}
