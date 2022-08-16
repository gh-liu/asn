package main

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Country .
type Country struct {
	Name string
	Code string
}

func parseCountryTable(r io.Reader) ([]Country, error) {
	root, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var countries []Country

	root.Find("tbody").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			var country Country
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				if indexth == 0 {
					country.Name = tablecell.ChildrenFiltered("a").Text()
					if href, exists := tablecell.ChildrenFiltered("a").Attr("href"); exists {
						country.Code = href[11:] // '/countries/US', trim '/countries/'
					}
				}
			})
			countries = append(countries, country)
		})
	})

	return countries, nil
}
