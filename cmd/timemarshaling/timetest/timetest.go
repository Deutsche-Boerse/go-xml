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

func CreateEvtDtPointer(in time.Time) *EvtDt {
	if p := makeTimeXSDPointer(in); p != nil {
		out := EvtDt(*p)
		return &out
	}
	return nil
}
func CreateEvtDt(in time.Time) (out EvtDt) {
	if p := makeTimeXSDPointer(in); p != nil {
		return EvtDt(*p)
	}
	return
}
func (t *EvtDt) GetTime() time.Time {
	return t.Time
}
func (t *EvtDt) UnmarshalText(text []byte) error {
	return (*XSDDate)(t).UnmarshalText(text)
}
func (t EvtDt) MarshalText() ([]byte, error) {
	return XSDDate(t).MarshalText()
}

type ISODateTime XSDDateTime

func CreateISODateTimePointer(in time.Time) *ISODateTime {
	if p := makeTimeXSDPointer(in); p != nil {
		out := ISODateTime(*p)
		return &out
	}
	return nil
}
func CreateISODateTime(in time.Time) (out ISODateTime) {
	if p := makeTimeXSDPointer(in); p != nil {
		return ISODateTime(*p)
	}
	return
}
func (t *ISODateTime) GetTime() time.Time {
	return t.Time
}
func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*XSDDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return XSDDateTime(t).MarshalText()
}

type ReportFile struct {
	ZerTs    *XSDDateTime  `json:"zerTs,omitempty" xml:"zerTs,omitempty"`
	OneTs    XSDDateTime   `json:"oneTs,omitempty" xml:"oneTs"`
	MulTs    []XSDDateTime `json:"mulTs,omitempty" xml:"mulTs,omitempty"`
	ZerDt    *XSDDate      `json:"zerDt,omitempty" xml:"zerDt,omitempty"`
	OneDt    XSDDate       `json:"oneDt,omitempty" xml:"oneDt"`
	ZerSt    *string       `json:"zerSt,omitempty" xml:"zerSt,omitempty"`
	OneSt    string        `json:"oneSt,omitempty" xml:"oneSt"`
	RptgDtTm ISODateTime   `json:"rptgDtTm,omitempty" xml:"RptgDtTm"`
	EvtDt    EvtDt         `json:"evtDt,omitempty" xml:"evtDt"`
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
func (t *ReportFile) GetRptgDtTm() (out ISODateTime) {
	if t == nil {
		return
	}
	return t.RptgDtTm
}
func (t *ReportFile) GetEvtDt() (out EvtDt) {
	if t == nil {
		return
	}
	return t.EvtDt
}

type XSDDate timeXSD

func CreateXSDDatePointer(in time.Time) *XSDDate {
	if p := makeTimeXSDPointer(in); p != nil {
		out := XSDDate(*p)
		return &out
	}
	return nil
}
func CreateXSDDate(in time.Time) (out XSDDate) {
	if p := makeTimeXSDPointer(in); p != nil {
		return XSDDate(*p)
	}
	return
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

func CreateXSDDateTimePointer(in time.Time) *XSDDateTime {
	if p := makeTimeXSDPointer(in); p != nil {
		out := XSDDateTime(*p)
		return &out
	}
	return nil
}
func CreateXSDDateTime(in time.Time) (out XSDDateTime) {
	if p := makeTimeXSDPointer(in); p != nil {
		return XSDDateTime(*p)
	}
	return
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
	return []byte((t.Time).Format("2006-01-02T15:04:05Z07:00")), nil
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
func makeTimeXSDPointer(in time.Time) *timeXSD {
	if in.IsZero() {
		return nil
	}
	return &timeXSD{Time: in}
}
