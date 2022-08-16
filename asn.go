package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// ASN .
type ASN struct {
	Name string
	ASN  string
}

func parseAsnTable(r io.Reader) ([]ASN, error) {
	root, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var asns []ASN

	root.Find("tbody").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			var asn ASN
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				if indexth == 0 {
					asn.ASN = tablecell.Text()[2:]
				}
				if indexth == 1 {
					asn.Name = tablecell.Text()
				}
			})
			asns = append(asns, asn)
		})
	})

	return asns, nil
}
