package xsdgen

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"
	"time"

	"aqwari.net/xml/xmltree"
)

func TestBooks(t *testing.T) {
	type Document struct {
		Books BooksForm `xml:"urn:books books"`
	}
	var document Document
	input, err := ioutil.ReadFile("testdata/books.xml")
	if err != nil {
		t.Fatal(err)
	}
	input = append([]byte("<Document>\n"), input...)
	input = append(input, []byte("</Document>")...)
	if err := xml.Unmarshal(input, &document); err != nil {
		t.Fatal("unmarshal: ", err)
	}
	output, err := xml.Marshal(&document)
	if err != nil {
		t.Fatal("marshal: ", err)
	}
	inputTree, err := xmltree.Parse(input)
	if err != nil {
		t.Fatal("testdata/books.xml: ", err)
	}
	outputTree, err := xmltree.Parse(output)
	if err != nil {
		t.Fatal("remarshal: ", err)
	}
	if !xmltree.Equal(inputTree, outputTree) {
		t.Errorf("got \n%s\n, wanted \n%s\n", xmltree.MarshalIndent(outputTree, "", "  "), xmltree.MarshalIndent(inputTree, "", "  "))
	}
}

type BookForm struct {
	Name    string    `xml:"name,attr"`
	Author  string    `xml:"urn:books author"`
	Title   string    `xml:"urn:books title"`
	Genre   string    `xml:"urn:books genre"`
	Price   float32   `xml:"urn:books price"`
	Pubdate time.Time `xml:"urn:books pub_date"`
	Review  string    `xml:"urn:books review"`
}

func (t *BookForm) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T BookForm
	var layout struct {
		*T
		Pubdate xsdDate `xml:"urn:books pub_date"`
	}
	layout.T = (*T)(t)
	layout.Pubdate = xsdDate(layout.T.Pubdate)
	return e.EncodeElement(layout, start)
}
func (t *BookForm) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T BookForm
	var overlay struct {
		*T
		Pubdate xsdDate `xml:"urn:books pub_date"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	overlay.T.Pubdate = time.Time(overlay.Pubdate)
	return nil
}

type BooksForm struct {
	Book []BookForm `xml:"urn:books book"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
