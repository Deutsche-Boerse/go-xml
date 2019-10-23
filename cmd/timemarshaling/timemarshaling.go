package main

import (
	"github.com/Deutsche-Boerse/go-xml/cmd/timemarshaling/timetest"
	"bufio"
	"encoding/xml"
	"log"
	"os"
	"time"
)

func main() {
	filename := "timetest.xml"
	unmarshaled, err := createStruct(filename)
	if err != nil {
		log.Fatalf("cannot unmarshal: %s", err)
	}
	log.Printf("%+v \n", unmarshaled)


	mapit(unmarshaled)
	marshalIt(unmarshaled)
}

func marshalIt(in timetest.ReportFile) {
	output, err := xml.MarshalIndent(in, "  ", "    ")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

func createStruct(filename string) (timetest.ReportFile, error) {

	var data timetest.ReportFile
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()
	dec := xml.NewDecoder(bufio.NewReader(file))
	err = dec.Decode(&data)

	return data, err


}

func mapit(in timetest.ReportFile) {
	o := mystruct{
		ZerTs: timeToDatetimeStringNoPointer(in.GetZerTs().GetTime()),
		OneTs: timeToDatetimeStringNoPointer(in.GetOneTs().GetTime()),
		ZerDt: timeToDateStringNoPointer(in.ZerDt.GetTime()),
		OneDt: timeToDateStringNoPointer(in.OneDt.GetTime()),
		ZerSt: in.GetZerSt(),
		OneSt: in.GetOneSt(),
		EvtDt: timeToDateStringNoPointer(in.EvtDt.GetTime()),
	}

	for _, v := range in.GetMulTs() {
		o.MulTs = append(o.MulTs, timeToDatetimeStringNoPointer(v.GetTime()))
	}

	log.Printf("%+v", o)

}

type mystruct struct {
	ZerTs string
	OneTs string
	ZerDt string
	MulTs []string
	OneDt string
	ZerSt string
	OneSt string
	EvtDt string
}



func timeToDateStringNoPointer(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return  t.Format("2006-01-02")
}



func timeToDatetimeStringNoPointer(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02T15:04:05Z")
}
