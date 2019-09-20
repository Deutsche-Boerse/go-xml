package timetest

import (
	"bytes"
	"encoding/xml"
	"time"
)

const (
	Namespace0 string = "http://deutsche-boerse.com/DBRegHub"
	Filename0  string = "timetest.xsd"
)

type EvtDt XSDDate

func (t *EvtDt) GetTime() time.Time {
	return t.Time
}
func (t *EvtDt) UnmarshalText(text []byte) error {
	return (*XSDDate)(t).UnmarshalText(text)
}
func (t EvtDt) MarshalText() ([]byte, error) {
	return XSDDate(t).MarshalText()
}

type ReportFile struct {
	ZerTs *XSDDateTime  `json:"zerTs,omitempty" xml:"zerTs,omitempty"`
	OneTs XSDDateTime   `json:"oneTs,omitempty" xml:"oneTs"`
	MulTs []XSDDateTime `json:"mulTs,omitempty" xml:"mulTs,omitempty"`
	ZerDt *XSDDate      `json:"zerDt,omitempty" xml:"zerDt,omitempty"`
	OneDt XSDDate       `json:"oneDt,omitempty" xml:"oneDt"`
	ZerSt *string       `json:"zerSt,omitempty" xml:"zerSt,omitempty"`
	OneSt string        `json:"oneSt,omitempty" xml:"oneSt"`
	EvtDt EvtDt         `json:"evtDt,omitempty" xml:"evtDt"`
}

func (t *ReportFile) GetZerTs() (out *XSDDateTime) {
	if t == nil {
		return
	}
	return t.ZerTs
}
func (t *ReportFile) GetOneTs() (out *XSDDateTime) {
	if t == nil {
		return
	}
	return &t.OneTs
}
func (t *ReportFile) GetMulTs() (out []XSDDateTime) {
	if t == nil {
		return
	}
	return t.MulTs
}
func (t *ReportFile) GetZerDt() (out *XSDDate) {
	if t == nil {
		return
	}
	return t.ZerDt
}
func (t *ReportFile) GetOneDt() (out *XSDDate) {
	if t == nil {
		return
	}
	return &t.OneDt
}
func (t *ReportFile) GetZerSt() (out string) {
	if t == nil || t.ZerSt == nil {
		return
	}
	return *t.ZerSt
}
func (t *ReportFile) GetOneSt() (out string) {
	if t == nil {
		return
	}
	return t.OneSt
}
func (t *ReportFile) GetEvtDt() (out EvtDt) {
	if t == nil {
		return
	}
	return t.EvtDt
}

type XSDDate timeXSD

func CreateXSDDate(in time.Time) *XSDDate {
	if in.IsZero() {
		return nil
	}
	return &XSDDate{Time: in}
}
func (t *XSDDate) GetTime() (out time.Time) {
	if t == nil {
		return
	}
	return t.Time
}
func (t *XSDDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*timeXSD)(t), "2006-01-02")
}
func (t XSDDate) MarshalText() ([]byte, error) {
	return []byte((t.Time).Format("2006-01-02")), nil
}
func (t XSDDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Time.IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t XSDDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t.Time.IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *timeXSD, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	t.Time, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		t.Time, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type XSDDateTime timeXSD

func CreateXSDDateTime(in time.Time) *XSDDateTime {
	if in.IsZero() {
		return nil
	}
	return &XSDDateTime{Time: in}
}
func (t *XSDDateTime) GetTime() (out time.Time) {
	if t == nil {
		return
	}
	return t.Time
}
func (t *XSDDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*timeXSD)(t), "2006-01-02T15:04:05.999999999")
}
func (t XSDDateTime) MarshalText() ([]byte, error) {
	return []byte((t.Time).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t XSDDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Time.IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t XSDDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t.Time.IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type timeXSD struct {
	time.Time
}

func (t *timeXSD) GetTime() time.Time {
	return t.Time
}
