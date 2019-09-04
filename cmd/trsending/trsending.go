package main

import (
	"aqwari.net/xml/cmd/trsending/sftrrepo"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	schemaInstanceNS = "http://www.w3.org/2001/XMLSchema-instance"
)

func createLEI(in string) *sftrrepo.LEIIdentifier {
	out := sftrrepo.LEIIdentifier(in)
	return &out
}

func main() {
	logger := log.New(os.Stderr, "", 0)
	filepath := "testfile.xml"

	file := sftrrepo.Document{
		SctiesFincgRptgTxRpt: sftrrepo.SecuritiesFinancingReportingTransactionReportV01{
			TradData: []sftrrepo.TradeTransactionReport6Choice{
				{
					New: &sftrrepo.TradeNewTransaction6{
						TechRcrdId: nil,
						CtrPtyData: sftrrepo.CounterpartyData48{
							RptgDtTm: sftrrepo.ISODateTime(time.Date(2019, 2, 3, 2, 3, 3, 0, time.UTC)),
							RptSubmitgNtty: sftrrepo.OrganisationIdentification9Choice{
								LEI:    createLEI(strings.Repeat("2", 20)),
								ClntId: nil,
								AnyBIC: nil,
							},
							CtrPtyData: []sftrrepo.CounterpartyData51{ // 1-2 times to be valid
								{
									RptgCtrPty: sftrrepo.CounterpartyIdentification1{
										Id: sftrrepo.OrganisationIdentification9Choice{
											LEI:    createLEI(strings.Repeat("0", 20)),
											ClntId: nil,
											AnyBIC: nil,
										},
										Ntr:   nil,
										Brnch: nil,
										Sd:    nil,
									},
									OthrCtrPty: sftrrepo.CounterpartyIdentification2{
										Id: sftrrepo.OrganisationIdentification9Choice{
											LEI:    createLEI(strings.Repeat("1", 20)),
											ClntId: nil,
											AnyBIC: nil,
										},
										Brnch:  nil,
										CtryCd: nil,
									},
									NttyRspnsblForRpt: nil,
									TxSpcfcData:       nil,
								},
							},
						},
						LnData: sftrrepo.LoanData44{
							UnqTradIdr: sftrrepo.Max52Text("aaaaa"),
							EvtDt:      sftrrepo.ISODate(time.Date(2019, 2, 3, 2, 3, 3, 0, time.UTC)),
							CtrctTp:    sftrrepo.ExposureType6Code("REPO"),
							ExctnDtTm:  sftrrepo.ISODateTime(time.Date(2019, 2, 3, 2, 3, 3, 0, time.UTC)),
							TxLnData: sftrrepo.TransactionLoanData5Choice{
								RpTrad: &sftrrepo.LoanData1{
									ClrSts:          nil,
									TradgVn:         nil,
									MstrAgrmt:       nil,
									ValDt:           nil,
									MinNtcePrd:      nil,
									EarlstCallBckDt: nil,
									GnlColl:         nil,
									DlvryByVal:      nil,
									CollDlvryMtd:    nil,
									Term:            nil,
									IntrstRate:      nil,
									PrncplAmt:       nil,
								},
								BuySellBck: nil,
								SctiesLndg: nil,
								MrgnLndg:   nil,
							},
						},
						CollData:    nil,
						LvlTp:       "PSTN",
						SplmtryData: nil,
					},
				},
			},
		},
	}

	logger.Printf("writing file: %s", filepath)
	err := saveFile(file, filepath)
	if err != nil {
		logger.Fatalf("Could not create XML file %s", err)
	}

}

func saveFile(file sftrrepo.Document, filepath string) error {

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(xml.Header))
	if err != nil {
		return err
	}
	//
	//xmlEncoder := xml.NewEncoder(f)
	//xmlEncoder.Indent("", "    ")
	//StartElement, EndElement, CharData, Comment, ProcInst,

	//pi := xml.ProcInst{
	//	Target: "version",
	//	Inst:   nil,
	//}

	//
	////<xs:schema xmlns="urn:iso:std:iso:20022:tech:xsd:DRAFT2auth.052.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" elementFormDefault="qualified" targetNamespace="urn:iso:std:iso:20022:tech:xsd:DRAFT2auth.052.001.01">
	//se := xml.StartElement{
	//	Name: xml.Name{
	//		Space: "urn:iso:std:iso:20022:tech:xsd:DRAFT2auth.052.001.01",
	//		Local: "Document",
	//	},
	//	Attr: []xml.Attr{
	//		{
	//			Name: xml.Name{
	//				Space: "http://www.w3.org/2001/XMLSchema-instance",
	//				Local: "schemaLocation",
	//			},
	//			Value: "urn:iso:std:iso:20022:tech:xsd:DRAFT2auth.052.001.01 DRAFT2auth.052.001.01",
	//		},
	//		//{
	//		//	Name:  xml.Name{
	//		//		//Space: "space2",
	//		//		Local: "local2",
	//		//	},
	//		//	//Value: "value2",
	//		//},
	//	},
	//}

	//ee := se.End()
	//
	//err = xmlEncoder.EncodeToken(se)
	//if err != nil {
	//	return err
	//}

	//err = xmlEncoder.EncodeElement(file, se)
	//if err != nil {
	//	return err
	//}
	////
	////err = xmlEncoder.EncodeToken(ee)
	////if err != nil {
	////	return err
	////}
	//
	//err = xmlEncoder.Flush()
	//if err != nil {
	//	return err
	//}

	tmp := struct {
		sftrrepo.Document
		XMLName xml.Name
		Attr []xml.Attr `xml:",attr"`
	}{
		Document: file,
		XMLName:  xml.Name{
			Space: sftrrepo.Namespace0,
			Local: reflect.TypeOf(file).Name(),
		},
		Attr: []xml.Attr{
			{
				Name: xml.Name{
					Local: "xmlns:xsi",
				},
				Value: schemaInstanceNS,
			},
			{
				Name: xml.Name{
					Local: "xsi:schemaLocation",
				},
				Value: fmt.Sprintf("%s %s", sftrrepo.Namespace0, sftrrepo.Filename0),
			},
		},
	}
	xmlData, err := xml.MarshalIndent(tmp, "", `   `)
	if err != nil {
		return err
	}

	_, err = f.Write(xmlData)
	if err != nil {
		return err
	}

	//if err = iso.XsdValidator.ValidateAgainstXsd(ctx, originalFileLocation1); err != nil {
	//	return err
	//}
	return f.Close()
}

