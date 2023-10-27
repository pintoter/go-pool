package entity

import "encoding/xml"

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string       `xml:"name" json:"name"`
		Time        string       `xml:"stovetime" json:"time"`
		Ingridients []Ingridient `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

type Ingridient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}
