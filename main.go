package main

import (
	"fmt"
	"log"
	"net/http"
)

var output string

func main() {
	res, err := http.Get("https://whois.ipip.net/countries")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	countries, err2 := parseCountryTable(res.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	var items []Item
	for _, c := range countries {
		url := fmt.Sprintf("https://whois.ipip.net/countries/%s", c.Code)
		res, err := http.Get(url)
		if err != nil {
			// log.Fatal(err)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			continue
		}

		asns, err3 := parseAsnTable(res.Body)
		if err3 != nil {
			continue
		}
		var item Item
		item.Country = c
		item.ASNs = asns

		items = append(items, item)
	}

	if err := parse(items); err != nil {
		panic(err)
	}
}
