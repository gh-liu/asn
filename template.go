package main

import (
	"html/template"
	"os"
)

type Item struct {
	Country
	ASNs []ASN
}

func parse(items []Item) error {
	file, err := os.OpenFile("out/asn.text", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	t := template.Must(template.ParseFiles("asn.tmpl"))
	err = t.Execute(file, items)
	if err != nil {
		return err
	}
	return nil
}
