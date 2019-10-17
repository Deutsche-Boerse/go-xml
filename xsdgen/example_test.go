package xsdgen_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"aqwari.net/xml/xsdgen"
)

func tmpfile() *os.File {
	f, err := ioutil.TempFile("", "xsdgen_test")
	if err != nil {
		panic(err)
	}
	return f
}

func xsdfile(s string) (filename string) {
	file := tmpfile()
	defer file.Close()
	fmt.Fprintf(file, `
		<schema xmlns="http://www.w3.org/2001/XMLSchema"
		        xmlns:tns="http://www.example.com/"
		        xmlns:xs="http://www.w3.org/2001/XMLSchema"
		        xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/"
		        xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/"
		        targetNamespace="http://www.example.com/">
		  %s
		</schema>
	`, s)
	return file.Name()
}

func ExampleConfig_GenCLI() {
	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.IgnoreAttributes("id", "href", "offset"),
		xsdgen.IgnoreElements("comment"),
		xsdgen.PackageName("webapi"),
		xsdgen.Replace("_", ""),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
	)
	if err := cfg.GenCLI("webapi.xsd", "deps/soap11.xsd"); err != nil {
		log.Fatal(err)
	}
}

func ExampleLogOutput() {
	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.LogOutput(log.New(os.Stderr, "", 0)),
		xsdgen.LogLevel(2))
	if err := cfg.GenCLI("file.wsdl"); err != nil {
		log.Fatal(err)
	}
}

func ExampleIgnoreAttributes() {
	doc := xsdfile(`
	  <complexType name="ArrayOfString">
	    <any maxOccurs="unbounded" />
	    <attribute name="soapenc:arrayType" type="xs:string" />
	  </complexType>
	`)
	var cfg xsdgen.Config
	cfg.Option(xsdgen.IgnoreAttributes("arrayType"))

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	// const Namespace0 string = "http://www.example.com/"
	//
	// type ArrayOfString struct {
	// 	Items []string `xml:",innerxml"`
	// }
}

func ExampleIgnoreElements() {
	doc := xsdfile(`
	  <complexType name="Person">
	    <sequence>
	      <element name="name" type="xs:string" />
	      <element name="deceased" type="soapenc:boolean" />
	      <element name="private" type="xs:int" />
	    </sequence>
	  </complexType>
	`)
	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.IgnoreElements("private"),
		xsdgen.IgnoreAttributes("id", "href"))

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	// const Namespace0 string = "http://www.example.com/"
	//
	// type Person struct {
	// 	Name     string `json:"name,omitempty" xml:"name"`
	//	Deceased bool   `json:"deceased,omitempty" xml:"deceased"`
	// }
}

func ExamplePackageName() {
	doc := xsdfile(`
	  <simpleType name="zipcode">
	    <restriction base="xs:string">
	      <length value="10" />
	    </restriction>
	  </simpleType>
	`)
	var cfg xsdgen.Config
	cfg.Option(xsdgen.PackageName("postal"))

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package postal
	//
	//const Namespace0 string = "http://www.example.com/"
}

func ExampleReplace() {
	doc := xsdfile(`
	  <complexType name="ArrayOfString">
	    <any maxOccurs="unbounded" />
	    <attribute name="soapenc:arrayType" type="xs:string" />
	  </complexType>
	`)
	var cfg xsdgen.Config
	cfg.Option(xsdgen.Replace("ArrayOf(.*)", "${1}Array"))

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	//const Namespace0 string = "http://www.example.com/"
	//
	// type StringArray struct {
	// 	Items     []string `xml:",innerxml"`
	// 	ArrayType string   `xml:"arrayType,attr,omitempty"`
	// }
}

func ExampleHandleSOAPArrayType() {
	doc := xsdfile(`
	  <complexType name="BoolArray">
	    <complexContent>
	      <restriction base="soapenc:Array">
	        <attribute ref="soapenc:arrayType" wsdl:arrayType="xs:boolean[]"/>
	      </restriction>
	    </complexContent>
	  </complexType>`)

	var cfg xsdgen.Config
	cfg.Option(xsdgen.HandleSOAPArrayType())

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	//const Namespace0 string = "http://www.example.com/"
	//
	// type BoolArray struct {
	// 	Items  []bool `xml:",innerxml"`
	// 	Offset string `xml:"offset,attr,omitempty"`
	// 	Id     string `xml:"id,attr,omitempty"`
	// 	Href   string `xml:"href,attr,omitempty"`
	// }
}

func ExampleSOAPArrayAsSlice() {
	doc := xsdfile(`
	  <complexType name="BoolArray">
	    <complexContent>
	      <restriction base="soapenc:Array">
	        <attribute ref="soapenc:arrayType" wsdl:arrayType="xs:boolean[]"/>
	      </restriction>
	    </complexContent>
	  </complexType>`)

	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
		xsdgen.LogOutput(log.New(os.Stderr, "", 0)),
		xsdgen.LogLevel(3),
		xsdgen.IgnoreAttributes("offset", "id", "href"))

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	// import "encoding/xml"
	//
	// const Namespace0 string = "http://www.example.com/"
	//
	// type BoolArray []bool
	//
	// func (a BoolArray) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// 	var output struct {
	// 		ArrayType string `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
	// 		Items     []bool `xml:" item"`
	// 	}
	// 	output.Items = []bool(a)
	// 	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{"", "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	// 	output.ArrayType = "ns1:boolean[]"
	// 	return e.EncodeElement(&output, start)
	// }
	// func (a *BoolArray) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	// 	var tok xml.Token
	// 	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
	// 		if tok, ok := tok.(xml.StartElement); ok {
	// 			var item bool
	// 			if err = d.DecodeElement(&item, &tok); err == nil {
	// 				*a = append(*a, item)
	// 			}
	// 		}
	// 		if _, ok := tok.(xml.EndElement); ok {
	// 			break
	// 		}
	// 	}
	// 	return err
	// }
}

func ExampleUseFieldNames() {
	doc := xsdfile(`
	  <complexType name="library">
	    <sequence>
	      <element name="book" maxOccurs="unbounded">
	        <complexType>
	          <all>
	            <element name="title" type="xs:string" />
	            <element name="published" type="xs:date" />
	            <element name="author" type="xs:string" />
	          </all>
	        </complexType>
	      </element>
	    </sequence>
	  </complexType>`)

	var cfg xsdgen.Config
	cfg.Option(xsdgen.UseFieldNames())

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	//import (
	//	"bytes"
	//	"encoding/xml"
	//	"time"
	//)
	//
	//const Namespace0 string = "http://www.example.com/"
	//
	//type Book struct {
	//	Title     string  `json:"title,omitempty" xml:"title"`
	//	Published XSDDate `json:"published,omitempty" xml:"published"`
	//	Author    string  `json:"author,omitempty" xml:"author"`
	//}
	//
	//type Library struct {
	//	Book []Book `json:"book,omitempty" xml:"book"`
	//}
	//
	//type XSDDate timeXSD
	//
	//func CreateXSDDatePointer(in time.Time) *XSDDate {
	//	if p := makeTimeXSDPointer(in); p != nil {
	//		out := XSDDate(*p)
	//		return &out
	//	}
	//	return nil
	//}
	//func CreateXSDDate(in time.Time) (out XSDDate) {
	//	if p := makeTimeXSDPointer(in); p != nil {
	//		return XSDDate(*p)
	//	}
	//	return
	//}
	//func (t *XSDDate) GetTime() (out time.Time) {
	//	if t == nil {
	//		return
	//	}
	//	return t.Time
	//}
	//func (t *XSDDate) UnmarshalText(text []byte) error {
	//	return _unmarshalTime(text, (*timeXSD)(t), "2006-01-02")
	//}
	//func (t XSDDate) MarshalText() ([]byte, error) {
	//	return []byte((t.Time).Format("2006-01-02")), nil
	//}
	//func (t XSDDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//	if t.Time.IsZero() {
	//		return nil
	//	}
	//	m, err := t.MarshalText()
	//	if err != nil {
	//		return err
	//	}
	//	return e.EncodeElement(m, start)
	//}
	//func (t XSDDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	//	if t.Time.IsZero() {
	//		return xml.Attr{}, nil
	//	}
	//	m, err := t.MarshalText()
	//	return xml.Attr{Name: name, Value: string(m)}, err
	//}
	//func _unmarshalTime(text []byte, t *timeXSD, format string) (err error) {
	//	s := string(bytes.TrimSpace(text))
	//	t.Time, err = time.Parse(format, s)
	//	if _, ok := err.(*time.ParseError); ok {
	//		t.Time, err = time.Parse(format+"Z07:00", s)
	//	}
	//	return err
	//}
	//
	//type timeXSD struct {
	//	time.Time
	//}
	//
	//func (t *timeXSD) GetTime() time.Time {
	//	return t.Time
	//}
	//func makeTimeXSDPointer(in time.Time) *timeXSD {
	//	if in.IsZero() {
	//		return nil
	//	}
	//	return &timeXSD{Time: in}
	//}

}

func ExampleOptionalDatetimeNoPointer() {
	doc := xsdfile(`
	  <complexType name="timeType">
	    <sequence>
	      <element name="clock" maxOccurs="unbounded">
	        <complexType>
	          <all>
	            <element name="timestamp" type="xs:dateTime" minOccurs="0" />
	          </all>
	        </complexType>
	      </element>
	    </sequence>
	  </complexType>`)

	var cfg xsdgen.Config
	cfg.Option(xsdgen.UseFieldNames())

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	//import (
	//	"bytes"
	//	"encoding/xml"
	//	"time"
	//)
	//
	//const Namespace0 string = "http://www.example.com/"
	//
	//type Clock struct {
	//	Timestamp *XSDDateTime `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	//}
	//
	//type TimeType struct {
	//	Clock []Clock `json:"clock,omitempty" xml:"clock"`
	//}
	//
	//type XSDDateTime timeXSD
	//
	//func CreateXSDDateTimePointer(in time.Time) *XSDDateTime {
	//	if p := makeTimeXSDPointer(in); p != nil {
	//		out := XSDDateTime(*p)
	//		return &out
	//	}
	//	return nil
	//}
	//func CreateXSDDateTime(in time.Time) (out XSDDateTime) {
	//	if p := makeTimeXSDPointer(in); p != nil {
	//		return XSDDateTime(*p)
	//	}
	//	return
	//}
	//func (t *XSDDateTime) GetTime() (out time.Time) {
	//	if t == nil {
	//		return
	//	}
	//	return t.Time
	//}
	//func (t *XSDDateTime) UnmarshalText(text []byte) error {
	//	return _unmarshalTime(text, (*timeXSD)(t), "2006-01-02T15:04:05.999999999")
	//}
	//func (t XSDDateTime) MarshalText() ([]byte, error) {
	//	return []byte((t.Time).Format("2006-01-02T15:04:05.999999999")), nil
	//}
	//func (t XSDDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//	if t.Time.IsZero() {
	//		return nil
	//	}
	//	m, err := t.MarshalText()
	//	if err != nil {
	//		return err
	//	}
	//	return e.EncodeElement(m, start)
	//}
	//func (t XSDDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	//	if t.Time.IsZero() {
	//		return xml.Attr{}, nil
	//	}
	//	m, err := t.MarshalText()
	//	return xml.Attr{Name: name, Value: string(m)}, err
	//}
	//func _unmarshalTime(text []byte, t *timeXSD, format string) (err error) {
	//	s := string(bytes.TrimSpace(text))
	//	t.Time, err = time.Parse(format, s)
	//	if _, ok := err.(*time.ParseError); ok {
	//		t.Time, err = time.Parse(format+"Z07:00", s)
	//	}
	//	return err
	//}
	//
	//type timeXSD struct {
	//	time.Time
	//}
	//
	//func (t *timeXSD) GetTime() time.Time {
	//	return t.Time
	//}
	//func makeTimeXSDPointer(in time.Time) *timeXSD {
	//	if in.IsZero() {
	//		return nil
	//	}
	//	return &timeXSD{Time: in}
	//}

}

func ExampleDecimalsAsString() {
	doc := xsdfile(`
	  <complexType name="monetary">
	    <sequence>
	      <element name="wallet" maxOccurs="unbounded">
	        <complexType>
	          <all>
	            <element name="amount" type="xs:string" />
	          </all>
	        </complexType>
	      </element>
	    </sequence>
	  </complexType>`)

	var cfg xsdgen.Config
	cfg.Option(xsdgen.UseFieldNames())
	cfg.Option(xsdgen.DecimalsAsString(false))

	out, err := cfg.GenSource(doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package ws
	//
	//const Namespace0 string = "http://www.example.com/"
	//
	//type Monetary struct {
	//	Wallet []Wallet `json:"wallet,omitempty" xml:"wallet"`
	//}
	//
	//type Wallet struct {
	//	Amount string `json:"amount,omitempty" xml:"amount"`
	//}

}
