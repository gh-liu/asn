package main

import "testing"

func Test_parse(t *testing.T) {
	var items []Item
	items = append(items, Item{
		Country: Country{
			Name: "china",
			Code: "CN",
		},
		ASNs: []ASN{{
			Name: "1",
			ASN:  "001",
		}, {
			Name: "2",
			ASN:  "002",
		}},
	})
	items = append(items, Item{
		Country: Country{
			Name: "china2",
			Code: "CN2",
		},
		ASNs: []ASN{{
			Name: "12",
			ASN:  "0012",
		}, {
			Name: "22",
			ASN:  "0022",
		}},
	})
	items = append(items, Item{
		Country: Country{
			Name: "china3",
			Code: "CN3",
		},
		ASNs: []ASN{{
			Name: "13",
			ASN:  "0013",
		}, {
			Name: "23",
			ASN:  "0023",
		}},
	})
	parse(items)
}
