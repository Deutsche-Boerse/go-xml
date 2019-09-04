package sftrrepo

import (
	"bytes"
	"encoding/xml"
	"time"
)

const (
	Namespace0 string = "urn:iso:std:iso:20022:tech:xsd:DRAFT2auth.052.001.01"
	Filename0  string = "DRAFT2auth.052.001.01.xsd"
)

type ActiveCurrencyAnd13DecimalAmount struct {
	Value string             `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value string                       `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AgreementType1Choice struct {
	Tp    *ExternalAgreementType1Code `json:"tp,omitempty" xml:"Tp,omitempty"`
	Prtry *Max35Text                  `json:"prtry,omitempty" xml:"Prtry,omitempty"`
}

type AgriculturalCommodityDairy1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType20Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type AgriculturalCommodityForestry1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType21Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type AgriculturalCommodityGrain2 struct {
	BasePdct     AssetClassProductType1Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType5Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType30Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type AgriculturalCommodityLiveStock1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType22Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type AgriculturalCommodityOilSeed1 struct {
	BasePdct     AssetClassProductType1Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType1Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType1Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type AgriculturalCommodityOliveOil2 struct {
	BasePdct     AssetClassProductType1Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType3Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType29Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type AgriculturalCommodityOther1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type AgriculturalCommodityPotato1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType45Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type AgriculturalCommoditySeafood1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType23Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type AgriculturalCommoditySoft1 struct {
	BasePdct     AssetClassProductType1Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType2Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType2Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type AmountAndDirection61 struct {
	Amt ActiveCurrencyAnd13DecimalAmount `json:"amt,omitempty" xml:"Amt"`
	Sgn *bool                            `json:"sgn,omitempty" xml:"Sgn,omitempty"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type AssetClassCommodity5Choice struct {
	Agrcltrl          *AssetClassCommodityAgricultural5Choice         `json:"agrcltrl,omitempty" xml:"Agrcltrl,omitempty"`
	Nrgy              *AssetClassCommodityEnergy2Choice               `json:"nrgy,omitempty" xml:"Nrgy,omitempty"`
	Envttl            *AssetClassCommodityEnvironmental2Choice        `json:"envttl,omitempty" xml:"Envttl,omitempty"`
	Frtlzr            *AssetClassCommodityFertilizer3Choice           `json:"frtlzr,omitempty" xml:"Frtlzr,omitempty"`
	Frght             *AssetClassCommodityFreight3Choice              `json:"frght,omitempty" xml:"Frght,omitempty"`
	IndstrlPdct       *AssetClassCommodityIndustrialProduct1Choice    `json:"indstrlPdct,omitempty" xml:"IndstrlPdct,omitempty"`
	Metl              *AssetClassCommodityMetal1Choice                `json:"metl,omitempty" xml:"Metl,omitempty"`
	OthrC10           *AssetClassCommodityOtherC102Choice             `json:"othrC10,omitempty" xml:"OthrC10,omitempty"`
	Ppr               *AssetClassCommodityPaper3Choice                `json:"ppr,omitempty" xml:"Ppr,omitempty"`
	Plprpln           *AssetClassCommodityPolypropylene3Choice        `json:"plprpln,omitempty" xml:"Plprpln,omitempty"`
	Infltn            *AssetClassCommodityInflation1                  `json:"infltn,omitempty" xml:"Infltn,omitempty"`
	MultiCmmdtyExtc   *AssetClassCommodityMultiCommodityExotic1       `json:"multiCmmdtyExtc,omitempty" xml:"MultiCmmdtyExtc,omitempty"`
	OffclEcnmcSttstcs *AssetClassCommodityOfficialEconomicStatistics1 `json:"offclEcnmcSttstcs,omitempty" xml:"OffclEcnmcSttstcs,omitempty"`
	Othr              *AssetClassCommodityOther1                      `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityAgricultural5Choice struct {
	GrnOilSeed *AgriculturalCommodityOilSeed1   `json:"grnOilSeed,omitempty" xml:"GrnOilSeed,omitempty"`
	Soft       *AgriculturalCommoditySoft1      `json:"soft,omitempty" xml:"Soft,omitempty"`
	Ptt        *AgriculturalCommodityPotato1    `json:"ptt,omitempty" xml:"Ptt,omitempty"`
	OlvOil     *AgriculturalCommodityOliveOil2  `json:"olvOil,omitempty" xml:"OlvOil,omitempty"`
	Dairy      *AgriculturalCommodityDairy1     `json:"dairy,omitempty" xml:"Dairy,omitempty"`
	Frstry     *AgriculturalCommodityForestry1  `json:"frstry,omitempty" xml:"Frstry,omitempty"`
	Sfd        *AgriculturalCommoditySeafood1   `json:"sfd,omitempty" xml:"Sfd,omitempty"`
	LiveStock  *AgriculturalCommodityLiveStock1 `json:"liveStock,omitempty" xml:"LiveStock,omitempty"`
	Grn        *AgriculturalCommodityGrain2     `json:"grn,omitempty" xml:"Grn,omitempty"`
	Othr       *AgriculturalCommodityOther1     `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityEnergy2Choice struct {
	Elctrcty  *EnergyCommodityElectricity1     `json:"elctrcty,omitempty" xml:"Elctrcty,omitempty"`
	NtrlGas   *EnergyCommodityNaturalGas2      `json:"ntrlGas,omitempty" xml:"NtrlGas,omitempty"`
	Oil       *EnergyCommodityOil2             `json:"oil,omitempty" xml:"Oil,omitempty"`
	Coal      *EnergyCommodityCoal1            `json:"coal,omitempty" xml:"Coal,omitempty"`
	IntrNrgy  *EnergyCommodityInterEnergy1     `json:"intrNrgy,omitempty" xml:"IntrNrgy,omitempty"`
	RnwblNrgy *EnergyCommodityRenewableEnergy1 `json:"rnwblNrgy,omitempty" xml:"RnwblNrgy,omitempty"`
	LghtEnd   *EnergyCommodityLightEnd1        `json:"lghtEnd,omitempty" xml:"LghtEnd,omitempty"`
	Dstllts   *EnergyCommodityDistillates1     `json:"dstllts,omitempty" xml:"Dstllts,omitempty"`
	Othr      *EnergyCommodityOther1           `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityEnvironmental2Choice struct {
	Emssns   *EnvironmentalCommodityEmission2      `json:"emssns,omitempty" xml:"Emssns,omitempty"`
	Wthr     *EnvironmentalCommodityWeather1       `json:"wthr,omitempty" xml:"Wthr,omitempty"`
	CrbnRltd *EnvironmentalCommodityCarbonRelated1 `json:"crbnRltd,omitempty" xml:"CrbnRltd,omitempty"`
	Othr     *EnvironmentCommodityOther1           `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityFertilizer3Choice struct {
	Ammn             *FertilizerCommodityAmmonia1                `json:"ammn,omitempty" xml:"Ammn,omitempty"`
	DmmnmPhspht      *FertilizerCommodityDiammoniumPhosphate1    `json:"dmmnmPhspht,omitempty" xml:"DmmnmPhspht,omitempty"`
	Ptsh             *FertilizerCommodityPotash1                 `json:"ptsh,omitempty" xml:"Ptsh,omitempty"`
	Slphr            *FertilizerCommoditySulphur1                `json:"slphr,omitempty" xml:"Slphr,omitempty"`
	Urea             *FertilizerCommodityUrea1                   `json:"urea,omitempty" xml:"Urea,omitempty"`
	UreaAndAmmnmNtrt *FertilizerCommodityUreaAndAmmoniumNitrate1 `json:"ureaAndAmmnmNtrt,omitempty" xml:"UreaAndAmmnmNtrt,omitempty"`
	Othr             *FertilizerCommodityOther1                  `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityFreight3Choice struct {
	Dry       *FreightCommodityDry2           `json:"dry,omitempty" xml:"Dry,omitempty"`
	Wet       *FreightCommodityWet2           `json:"wet,omitempty" xml:"Wet,omitempty"`
	CntnrShip *FreightCommodityContainerShip1 `json:"cntnrShip,omitempty" xml:"CntnrShip,omitempty"`
	Othr      *FreightCommodityOther1         `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityIndustrialProduct1Choice struct {
	Cnstrctn *IndustrialProductCommodityConstruction1  `json:"cnstrctn,omitempty" xml:"Cnstrctn,omitempty"`
	Manfctg  *IndustrialProductCommodityManufacturing1 `json:"manfctg,omitempty" xml:"Manfctg,omitempty"`
}

type AssetClassCommodityInflation1 struct {
	BasePdct AssetClassProductType12Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

type AssetClassCommodityMetal1Choice struct {
	NonPrcs *MetalCommodityNonPrecious1 `json:"nonPrcs,omitempty" xml:"NonPrcs,omitempty"`
	Prcs    *MetalCommodityPrecious1    `json:"prcs,omitempty" xml:"Prcs,omitempty"`
}

type AssetClassCommodityMultiCommodityExotic1 struct {
	BasePdct AssetClassProductType13Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

type AssetClassCommodityOfficialEconomicStatistics1 struct {
	BasePdct AssetClassProductType14Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

type AssetClassCommodityOther1 struct {
	BasePdct AssetClassProductType15Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

type AssetClassCommodityOtherC102Choice struct {
	Dlvrbl    *OtherC10CommodityDeliverable2    `json:"dlvrbl,omitempty" xml:"Dlvrbl,omitempty"`
	NonDlvrbl *OtherC10CommodityNonDeliverable2 `json:"nonDlvrbl,omitempty" xml:"NonDlvrbl,omitempty"`
}

type AssetClassCommodityPaper3Choice struct {
	CntnrBrd *PaperCommodityContainerBoard1 `json:"cntnrBrd,omitempty" xml:"CntnrBrd,omitempty"`
	Nwsprnt  *PaperCommodityNewsprint1      `json:"nwsprnt,omitempty" xml:"Nwsprnt,omitempty"`
	Pulp     *PaperCommodityPulp1           `json:"pulp,omitempty" xml:"Pulp,omitempty"`
	RcvrdPpr *PaperCommodityRecoveredPaper1 `json:"rcvrdPpr,omitempty" xml:"RcvrdPpr,omitempty"`
	Othr     *PaperCommodityRecoveredPaper2 `json:"othr,omitempty" xml:"Othr,omitempty"`
}

type AssetClassCommodityPolypropylene3Choice struct {
	Plstc *PolypropyleneCommodityPlastic1 `json:"plstc,omitempty" xml:"Plstc,omitempty"`
	Othr  *PolypropyleneCommodityOther1   `json:"othr,omitempty" xml:"Othr,omitempty"`
}

// May be one of ALUM, ALUA, CBLT, COPR, IRON, MOLY, NASC, NICK, STEL, TINN, ZINC, OTHR, LEAD
type AssetClassDetailedSubProductType10Code string

// May be one of GOLD, OTHR, PLDM, PTNM, SLVR
type AssetClassDetailedSubProductType11Code string

// May be one of FWHT, SOYB, RPSD, OTHR, CORN, RICE
type AssetClassDetailedSubProductType1Code string

// May be one of LAMP, OTHR
type AssetClassDetailedSubProductType29Code string

// May be one of ROBU, CCOA, BRWN, WHSG, OTHR
type AssetClassDetailedSubProductType2Code string

// May be one of MWHT, OTHR
type AssetClassDetailedSubProductType30Code string

// May be one of GASP, LNGG, NCGG, TTFG, NBPG, OTHR
type AssetClassDetailedSubProductType31Code string

// May be one of BAKK, BDSL, BRNT, BRNX, CNDA, COND, DSEL, DUBA, ESPO, ETHA, FUEL, FOIL, GOIL, GSLN, HEAT, JTFL, KERO, LLSO, MARS, NAPH, NGLO, TAPI, WTIO, URAL, OTHR
type AssetClassDetailedSubProductType32Code string

// May be one of DBCR, OTHR
type AssetClassDetailedSubProductType33Code string

// May be one of TNKR, OTHR
type AssetClassDetailedSubProductType34Code string

// May be one of BSLD, FITR, PKLD, OFFP, OTHR
type AssetClassDetailedSubProductType5Code string

// May be one of CERE, ERUE, EUAE, EUAA, OTHR
type AssetClassDetailedSubProductType8Code string

// May be one of OTHC
type AssetClassProductType11Code string

// May be one of INFL
type AssetClassProductType12Code string

// May be one of MCEX
type AssetClassProductType13Code string

// May be one of OEST
type AssetClassProductType14Code string

// May be one of OTHR
type AssetClassProductType15Code string

// May be one of AGRI
type AssetClassProductType1Code string

// May be one of NRGY
type AssetClassProductType2Code string

// May be one of ENVR
type AssetClassProductType3Code string

// May be one of FRGT
type AssetClassProductType4Code string

// May be one of FRTL
type AssetClassProductType5Code string

// May be one of INDP
type AssetClassProductType6Code string

// May be one of METL
type AssetClassProductType7Code string

// May be one of PAPR
type AssetClassProductType8Code string

// May be one of POLY
type AssetClassProductType9Code string

// May be one of EMIS
type AssetClassSubProductType10Code string

// May be one of NPRM
type AssetClassSubProductType15Code string

// May be one of PRME
type AssetClassSubProductType16Code string

// May be one of PLST
type AssetClassSubProductType18Code string

// May be one of GROS
type AssetClassSubProductType1Code string

// May be one of DIRY
type AssetClassSubProductType20Code string

// May be one of FRST
type AssetClassSubProductType21Code string

// May be one of LSTK
type AssetClassSubProductType22Code string

// May be one of SEAF
type AssetClassSubProductType23Code string

// May be one of COAL
type AssetClassSubProductType24Code string

// May be one of DIST
type AssetClassSubProductType25Code string

// May be one of INRG
type AssetClassSubProductType26Code string

// May be one of LGHT
type AssetClassSubProductType27Code string

// May be one of RNNG
type AssetClassSubProductType28Code string

// May be one of CRBR
type AssetClassSubProductType29Code string

// May be one of SOFT
type AssetClassSubProductType2Code string

// May be one of WTHR
type AssetClassSubProductType30Code string

// May be one of DRYF
type AssetClassSubProductType31Code string

// May be one of WETF
type AssetClassSubProductType32Code string

// May be one of CSTR
type AssetClassSubProductType33Code string

// May be one of MFTG
type AssetClassSubProductType34Code string

// May be one of CBRD
type AssetClassSubProductType35Code string

// May be one of NSPT
type AssetClassSubProductType36Code string

// May be one of PULP
type AssetClassSubProductType37Code string

// May be one of RCVP
type AssetClassSubProductType38Code string

// May be one of AMMO
type AssetClassSubProductType39Code string

// May be one of OOLI
type AssetClassSubProductType3Code string

// May be one of DAPH
type AssetClassSubProductType40Code string

// May be one of PTSH
type AssetClassSubProductType41Code string

// May be one of SLPH
type AssetClassSubProductType42Code string

// May be one of UREA
type AssetClassSubProductType43Code string

// May be one of UAAN
type AssetClassSubProductType44Code string

// May be one of POTA
type AssetClassSubProductType45Code string

// May be one of CSHP
type AssetClassSubProductType46Code string

// May be one of DLVR
type AssetClassSubProductType47Code string

// May be one of NDLV
type AssetClassSubProductType48Code string

// May be one of OTHR
type AssetClassSubProductType49Code string

// May be one of GRIN
type AssetClassSubProductType5Code string

// May be one of ELEC
type AssetClassSubProductType6Code string

// May be one of NGAS
type AssetClassSubProductType7Code string

// May be one of OILP
type AssetClassSubProductType8Code string

// May be one of WIBO, TREA, TIBO, TLBO, SWAP, STBO, PRBO, PFAN, NIBO, MAAA, MOSP, LIBO, LIBI, JIBA, ISDA, GCFR, FUSW, EUCH, EUUS, EURI, EONS, EONA, CIBO, CDOR, BUBO, BBSW
type BenchmarkCurveName2Code string

type BenchmarkCurveName8Choice struct {
	Indx *BenchmarkCurveName2Code `json:"indx,omitempty" xml:"Indx,omitempty"`
	Nm   *Max350Text              `json:"nm,omitempty" xml:"Nm,omitempty"`
}

type Branch2Choice struct {
	Id   *OrganisationIdentification9Choice `json:"id,omitempty" xml:"Id,omitempty"`
	Ctry *CountryCode                       `json:"ctry,omitempty" xml:"Ctry,omitempty"`
}

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type Cleared2Choice struct {
	Clrd    *ClearingPartyAndTime7 `json:"clrd,omitempty" xml:"Clrd,omitempty"`
	NonClrd *NoReasonCode          `json:"nonClrd,omitempty" xml:"NonClrd,omitempty"`
}

type Cleared8Choice struct {
	Clrd    *ClearingPartyAndTime7 `json:"clrd,omitempty" xml:"Clrd,omitempty"`
	NonClrd *NoReasonCode          `json:"nonClrd,omitempty" xml:"NonClrd,omitempty"`
}

type ClearingPartyAndTime7 struct {
	CCP        *LEIIdentifier `json:"ccp,omitempty" xml:"CCP,omitempty"`
	ClrDtTm    *ISODateTime   `json:"clrDtTm,omitempty" xml:"ClrDtTm,omitempty"`
	RptTrckgNb *Max52Text     `json:"rptTrckgNb,omitempty" xml:"RptTrckgNb,omitempty"`
	PrtflCd    *Max52Text     `json:"prtflCd,omitempty" xml:"PrtflCd,omitempty"`
}

type Collateral15 struct {
	CollValDt         *ISODate                        `json:"collValDt,omitempty" xml:"CollValDt,omitempty"`
	AsstTp            []CollateralType7Choice         `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	NetXpsrCollstnInd *bool                           `json:"netXpsrCollstnInd,omitempty" xml:"NetXpsrCollstnInd,omitempty"`
	BsktIdr           *SecurityIdentification26Choice `json:"bsktIdr,omitempty" xml:"BsktIdr,omitempty"`
}

type Collateral27 struct {
	CollValDt *ISODate `json:"collValDt,omitempty" xml:"CollValDt,omitempty"`
}

type CollateralData3 struct {
	AsstTp  []CollateralType7Choice         `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	BsktIdr *SecurityIdentification26Choice `json:"bsktIdr,omitempty" xml:"BsktIdr,omitempty"`
}

type CollateralData7 struct {
	TxCollData TransactionCollateralData1Choice `json:"txCollData,omitempty" xml:"TxCollData"`
}

// May be one of SICA, SIUR, TTCA
type CollateralDeliveryMethod1Code string

type CollateralFlag6 struct {
	HrcutOrMrgn string `json:"hrcutOrMrgn,omitempty" xml:"HrcutOrMrgn"`
}

type CollateralFlag6Choice struct {
	Collsd   *CollaterisedData4 `json:"collsd,omitempty" xml:"Collsd,omitempty"`
	Uncollsd *NoReasonCode      `json:"uncollsd,omitempty" xml:"Uncollsd,omitempty"`
}

// May be one of INVG, NIVG, NOTR, NOAP
type CollateralQualityType1Code string

// May be one of GIVE, TAKE
type CollateralRole1Code string

type CollateralType7Choice struct {
	Scty   []HaircutPortfolioSecurityIdentification1 `json:"scty,omitempty" xml:"Scty,omitempty"`
	Csh    []ActiveOrHistoricCurrencyAndAmount       `json:"csh,omitempty" xml:"Csh,omitempty"`
	Cmmdty []Commodity3                              `json:"cmmdty,omitempty" xml:"Cmmdty,omitempty"`
}

type CollaterisedData4 struct {
	CollValDt         *ISODate                        `json:"collValDt,omitempty" xml:"CollValDt,omitempty"`
	AsstTp            *CollateralType7Choice          `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	NetXpsrCollstnInd *bool                           `json:"netXpsrCollstnInd,omitempty" xml:"NetXpsrCollstnInd,omitempty"`
	BsktIdr           *SecurityIdentification26Choice `json:"bsktIdr,omitempty" xml:"BsktIdr,omitempty"`
}

type Commodity3 struct {
	Clssfctn *AssetClassCommodity5Choice        `json:"clssfctn,omitempty" xml:"Clssfctn,omitempty"`
	Qty      *Quantity13                        `json:"qty,omitempty" xml:"Qty,omitempty"`
	UnitPric *SecuritiesTransactionPrice2Choice `json:"unitPric,omitempty" xml:"UnitPric,omitempty"`
	MktVal   *string                            `json:"mktVal,omitempty" xml:"MktVal,omitempty"`
}

type ContractTerm2Choice struct {
	Opn *NoReasonCode       `json:"opn,omitempty" xml:"Opn,omitempty"`
	Fxd *FixedTermContract2 `json:"fxd,omitempty" xml:"Fxd,omitempty"`
}

type CounterpartyData48 struct {
	RptgDtTm       ISODateTime                       `json:"rptgDtTm,omitempty" xml:"RptgDtTm"`
	RptSubmitgNtty OrganisationIdentification9Choice `json:"rptSubmitgNtty,omitempty" xml:"RptSubmitgNtty"`
	CtrPtyData     []CounterpartyData51              `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
}

type CounterpartyData49 struct {
	RptgDtTm       ISODateTime                       `json:"rptgDtTm,omitempty" xml:"RptgDtTm"`
	RptSubmitgNtty OrganisationIdentification9Choice `json:"rptSubmitgNtty,omitempty" xml:"RptSubmitgNtty"`
	CtrPtyData     []CounterpartyData50              `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
}

type CounterpartyData50 struct {
	RptgCtrPty CounterpartyIdentification1 `json:"rptgCtrPty,omitempty" xml:"RptgCtrPty"`
	OthrCtrPty CounterpartyIdentification2 `json:"othrCtrPty,omitempty" xml:"OthrCtrPty"`
}

type CounterpartyData51 struct {
	RptgCtrPty        CounterpartyIdentification1         `json:"rptgCtrPty,omitempty" xml:"RptgCtrPty"`
	OthrCtrPty        CounterpartyIdentification2         `json:"othrCtrPty,omitempty" xml:"OthrCtrPty"`
	NttyRspnsblForRpt *OrganisationIdentification9Choice  `json:"nttyRspnsblForRpt,omitempty" xml:"NttyRspnsblForRpt,omitempty"`
	TxSpcfcData       *TransactionCounterpartyData3Choice `json:"txSpcfcData,omitempty" xml:"TxSpcfcData,omitempty"`
}

type CounterpartyIdentification1 struct {
	Id    OrganisationIdentification9Choice `json:"id,omitempty" xml:"Id"`
	Ntr   *CounterpartyTradeNature3Choice   `json:"ntr,omitempty" xml:"Ntr,omitempty"`
	Brnch *Branch2Choice                    `json:"brnch,omitempty" xml:"Brnch,omitempty"`
	Sd    *CollateralRole1Code              `json:"sd,omitempty" xml:"Sd,omitempty"`
}

type CounterpartyIdentification2 struct {
	Id     OrganisationIdentification9Choice `json:"id,omitempty" xml:"Id"`
	Brnch  *Branch2Choice                    `json:"brnch,omitempty" xml:"Brnch,omitempty"`
	CtryCd *CountryCode                      `json:"ctryCd,omitempty" xml:"CtryCd,omitempty"`
}

type CounterpartyTradeNature3Choice struct {
	FI  *FinancialPartyClassification1Choice `json:"fi,omitempty" xml:"FI,omitempty"`
	NFI []NACEDomainIdentifier               `json:"nfi,omitempty" xml:"NFI,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	SctiesFincgRptgTxRpt SecuritiesFinancingReportingTransactionReportV01 `json:"sctiesFincgRptgTxRpt,omitempty" xml:"SctiesFincgRptgTxRpt"`
}

type EnergyCommodityCoal1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType24Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnergyCommodityDistillates1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType25Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnergyCommodityElectricity1 struct {
	BasePdct     AssetClassProductType2Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType6Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType5Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type EnergyCommodityInterEnergy1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType26Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnergyCommodityLightEnd1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType27Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnergyCommodityNaturalGas2 struct {
	BasePdct     AssetClassProductType2Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType7Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType31Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type EnergyCommodityOil2 struct {
	BasePdct     AssetClassProductType2Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType8Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType32Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type EnergyCommodityOther1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnergyCommodityRenewableEnergy1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType28Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnvironmentCommodityOther1 struct {
	BasePdct AssetClassProductType3Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnvironmentalCommodityCarbonRelated1 struct {
	BasePdct AssetClassProductType3Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType29Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type EnvironmentalCommodityEmission2 struct {
	BasePdct     AssetClassProductType3Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType10Code        `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType8Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type EnvironmentalCommodityWeather1 struct {
	BasePdct AssetClassProductType3Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType30Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

// May be one of SBSC, MGLD, REPO, SLEB
type ExposureType6Code string

// Must be at least 1 items long
type ExternalAgreementType1Code string

// Must be at least 1 items long
type ExternalSecuritiesLendingType1Code string

type FertilizerCommodityAmmonia1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType39Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FertilizerCommodityDiammoniumPhosphate1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType40Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FertilizerCommodityOther1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FertilizerCommodityPotash1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType41Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FertilizerCommoditySulphur1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType42Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FertilizerCommodityUrea1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType43Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FertilizerCommodityUreaAndAmmoniumNitrate1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType44Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FinancialPartyClassification1Choice struct {
	Clssfctn           []FinancialPartySectorType2Code `json:"clssfctn,omitempty" xml:"Clssfctn,omitempty"`
	InvstmtFndClssfctn *FundType2Code                  `json:"invstmtFndClssfctn,omitempty" xml:"InvstmtFndClssfctn,omitempty"`
}

// May be one of AIFD, CSDS, CCPS, CDTI, INUN, ORPI, INVF, REIN, UCIT
type FinancialPartySectorType2Code string

type FixedRate2 struct {
	Rate       string                                 `json:"rate,omitempty" xml:"Rate"`
	DayCntBsis InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
}

type FixedRate7 struct {
	Rate       *string                                `json:"rate,omitempty" xml:"Rate,omitempty"`
	DayCntBsis InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
	MrgnLnAmt  *ActiveOrHistoricCurrencyAndAmount     `json:"mrgnLnAmt,omitempty" xml:"MrgnLnAmt,omitempty"`
}

type FixedTermContract2 struct {
	MtrtyDt     *ISODate                    `json:"mtrtyDt,omitempty" xml:"MtrtyDt,omitempty"`
	TermntnOptn *RepoTerminationOption2Code `json:"termntnOptn,omitempty" xml:"TermntnOptn,omitempty"`
}

type FloatingInterestRate10 struct {
	RefRate    BenchmarkCurveName8Choice `json:"refRate,omitempty" xml:"RefRate"`
	Term       InterestRateContractTerm2 `json:"term,omitempty" xml:"Term"`
	PmtFrqcy   InterestRateContractTerm2 `json:"pmtFrqcy,omitempty" xml:"PmtFrqcy"`
	RstFrqcy   InterestRateContractTerm2 `json:"rstFrqcy,omitempty" xml:"RstFrqcy"`
	BsisPtSprd string                    `json:"bsisPtSprd,omitempty" xml:"BsisPtSprd"`
}

type FloatingInterestRate12 struct {
	RefRate      BenchmarkCurveName8Choice              `json:"refRate,omitempty" xml:"RefRate"`
	Term         InterestRateContractTerm2              `json:"term,omitempty" xml:"Term"`
	PmtFrqcy     InterestRateContractTerm2              `json:"pmtFrqcy,omitempty" xml:"PmtFrqcy"`
	RstFrqcy     InterestRateContractTerm2              `json:"rstFrqcy,omitempty" xml:"RstFrqcy"`
	BsisPtSprd   string                                 `json:"bsisPtSprd,omitempty" xml:"BsisPtSprd"`
	RateAdjstmnt []RateAdjustment1                      `json:"rateAdjstmnt,omitempty" xml:"RateAdjstmnt,omitempty"`
	DayCntBsis   InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
}

type FloatingInterestRate16 struct {
	RefRate    *BenchmarkCurveName8Choice             `json:"refRate,omitempty" xml:"RefRate,omitempty"`
	Term       *InterestRateContractTerm2             `json:"term,omitempty" xml:"Term,omitempty"`
	PmtFrqcy   *InterestRateContractTerm2             `json:"pmtFrqcy,omitempty" xml:"PmtFrqcy,omitempty"`
	RstFrqcy   *InterestRateContractTerm2             `json:"rstFrqcy,omitempty" xml:"RstFrqcy,omitempty"`
	BsisPtSprd *string                                `json:"bsisPtSprd,omitempty" xml:"BsisPtSprd,omitempty"`
	DayCntBsis InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
}

type FreightCommodityContainerShip1 struct {
	BasePdct AssetClassProductType4Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType46Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FreightCommodityDry2 struct {
	BasePdct     AssetClassProductType4Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType31Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType33Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type FreightCommodityOther1 struct {
	BasePdct AssetClassProductType4Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type FreightCommodityWet2 struct {
	BasePdct     AssetClassProductType4Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType32Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType34Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

// May be one of ETFT, MMFT, OTHR, REIT
type FundType2Code string

type HaircutPortfolioSecurityIdentification1 struct {
	PrtflHrcutOrMrgn *string   `json:"prtflHrcutOrMrgn,omitempty" xml:"PrtflHrcutOrMrgn,omitempty"`
	Id               Security3 `json:"id,omitempty" xml:"Id"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISORestrictedYear time.Time

func (t *ISORestrictedYear) UnmarshalText(text []byte) error {
	return (*xsdGYear)(t).UnmarshalText(text)
}
func (t ISORestrictedYear) MarshalText() ([]byte, error) {
	return xsdGYear(t).MarshalText()
}

type IndustrialProductCommodityConstruction1 struct {
	BasePdct AssetClassProductType6Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType33Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type IndustrialProductCommodityManufacturing1 struct {
	BasePdct AssetClassProductType6Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType34Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014
type InterestComputationMethod1Code string

type InterestComputationMethodFormat6Choice struct {
	Cd    *InterestComputationMethod1Code `json:"cd,omitempty" xml:"Cd,omitempty"`
	Prtry *Max35Text                      `json:"prtry,omitempty" xml:"Prtry,omitempty"`
}

type InterestRate14Choice struct {
	Fxd  *FixedRate2             `json:"fxd,omitempty" xml:"Fxd,omitempty"`
	Fltg *FloatingInterestRate12 `json:"fltg,omitempty" xml:"Fltg,omitempty"`
}

type InterestRate15Choice struct {
	Amt        *ActiveOrHistoricCurrencyAndAmount `json:"amt,omitempty" xml:"Amt,omitempty"`
	IntrstRate *InterestRate7Choice               `json:"intrstRate,omitempty" xml:"IntrstRate,omitempty"`
}

type InterestRate7Choice struct {
	Fxd  *FixedRate7             `json:"fxd,omitempty" xml:"Fxd,omitempty"`
	Fltg *FloatingInterestRate16 `json:"fltg,omitempty" xml:"Fltg,omitempty"`
}

type InterestRateContractTerm2 struct {
	Unit RateBasis1Code `json:"unit,omitempty" xml:"Unit"`
	Val  string         `json:"val,omitempty" xml:"Val"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LoanData1 struct {
	ClrSts          *Cleared8Choice                `json:"clrSts,omitempty" xml:"ClrSts,omitempty"`
	TradgVn         *MICIdentifier                 `json:"tradgVn,omitempty" xml:"TradgVn,omitempty"`
	MstrAgrmt       *MasterAgreement1              `json:"mstrAgrmt,omitempty" xml:"MstrAgrmt,omitempty"`
	ValDt           *ISODate                       `json:"valDt,omitempty" xml:"ValDt,omitempty"`
	MinNtcePrd      *string                        `json:"minNtcePrd,omitempty" xml:"MinNtcePrd,omitempty"`
	EarlstCallBckDt *ISODate                       `json:"earlstCallBckDt,omitempty" xml:"EarlstCallBckDt,omitempty"`
	GnlColl         *SpecialCollateral1Code        `json:"gnlColl,omitempty" xml:"GnlColl,omitempty"`
	DlvryByVal      *bool                          `json:"dlvryByVal,omitempty" xml:"DlvryByVal,omitempty"`
	CollDlvryMtd    *CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd,omitempty"`
	Term            []ContractTerm2Choice          `json:"term,omitempty" xml:"Term,omitempty"`
	IntrstRate      *InterestRate14Choice          `json:"intrstRate,omitempty" xml:"IntrstRate,omitempty"`
	PrncplAmt       *PrincipalAmount1              `json:"prncplAmt,omitempty" xml:"PrncplAmt,omitempty"`
}

type LoanData11 struct {
	UnqTradIdr Max52Text `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt      ISODate   `json:"evtDt,omitempty" xml:"EvtDt"`
	TermntnDt  ISODate   `json:"termntnDt,omitempty" xml:"TermntnDt"`
}

type LoanData13 struct {
	UnqTradIdr Max52Text `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
}

type LoanData2 struct {
	ClrSts    *Cleared2Choice                    `json:"clrSts,omitempty" xml:"ClrSts,omitempty"`
	TradgVn   *MICIdentifier                     `json:"tradgVn,omitempty" xml:"TradgVn,omitempty"`
	MstrAgrmt *MasterAgreement1                  `json:"mstrAgrmt,omitempty" xml:"MstrAgrmt,omitempty"`
	ValDt     *ISODate                           `json:"valDt,omitempty" xml:"ValDt,omitempty"`
	MtrtyDt   *ISODate                           `json:"mtrtyDt,omitempty" xml:"MtrtyDt,omitempty"`
	GnlColl   *SpecialCollateral1Code            `json:"gnlColl,omitempty" xml:"GnlColl,omitempty"`
	PrncplAmt *PrincipalAmount1                  `json:"prncplAmt,omitempty" xml:"PrncplAmt,omitempty"`
	UnitPric  *SecuritiesTransactionPrice2Choice `json:"unitPric,omitempty" xml:"UnitPric,omitempty"`
}

type LoanData27 struct {
	PrncplAmt PrincipalAmount1 `json:"prncplAmt,omitempty" xml:"PrncplAmt"`
	MtrtyDt   ISODate          `json:"mtrtyDt,omitempty" xml:"MtrtyDt"`
}

type LoanData29 struct {
	UnqTradIdr *Max52Text       `json:"unqTradIdr,omitempty" xml:"UnqTradIdr,omitempty"`
	MstrAgrmt  MasterAgreement1 `json:"mstrAgrmt,omitempty" xml:"MstrAgrmt"`
}

type LoanData32 struct {
	CollDlvryMtd     *CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd,omitempty"`
	OutsdngMrgnLnAmt *string                        `json:"outsdngMrgnLnAmt,omitempty" xml:"OutsdngMrgnLnAmt,omitempty"`
	ShrtMktValAmt    *string                        `json:"shrtMktValAmt,omitempty" xml:"ShrtMktValAmt,omitempty"`
	Ccy              *ActiveOrHistoricCurrencyCode  `json:"ccy,omitempty" xml:"Ccy,omitempty"`
	MrgnLnAttr       []InterestRate15Choice         `json:"mrgnLnAttr,omitempty" xml:"MrgnLnAttr,omitempty"`
}

type LoanData37 struct {
	DlvryByVal   bool                          `json:"dlvryByVal,omitempty" xml:"DlvryByVal"`
	CollDlvryMtd CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd"`
	Term         []ContractTerm2Choice         `json:"term,omitempty" xml:"Term,omitempty"`
	IntrstRate   InterestRate14Choice          `json:"intrstRate,omitempty" xml:"IntrstRate"`
	PrncplAmt    PrincipalAmount1              `json:"prncplAmt,omitempty" xml:"PrncplAmt"`
}

type LoanData39 struct {
	ClrSts        *Cleared2Choice                `json:"clrSts,omitempty" xml:"ClrSts,omitempty"`
	TradgVn       *MICIdentifier                 `json:"tradgVn,omitempty" xml:"TradgVn,omitempty"`
	MstrAgrmt     *MasterAgreement1              `json:"mstrAgrmt,omitempty" xml:"MstrAgrmt,omitempty"`
	ValDt         *ISODate                       `json:"valDt,omitempty" xml:"ValDt,omitempty"`
	GnlColl       *SpecialCollateral1Code        `json:"gnlColl,omitempty" xml:"GnlColl,omitempty"`
	DlvryByVal    *bool                          `json:"dlvryByVal,omitempty" xml:"DlvryByVal,omitempty"`
	CollDlvryMtd  *CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd,omitempty"`
	Term          []ContractTerm2Choice          `json:"term,omitempty" xml:"Term,omitempty"`
	AsstTp        *SecurityCommodity2Choice      `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	LnVal         *string                        `json:"lnVal,omitempty" xml:"LnVal,omitempty"`
	RbtRate       *RebateRate1Choice             `json:"rbtRate,omitempty" xml:"RbtRate,omitempty"`
	LndgFee       *string                        `json:"lndgFee,omitempty" xml:"LndgFee,omitempty"`
	ExclsvArrgmnt *bool                          `json:"exclsvArrgmnt,omitempty" xml:"ExclsvArrgmnt,omitempty"`
}

type LoanData4 struct {
	EvtDt    ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	TxLnData TransactionLoanData4Choice `json:"txLnData,omitempty" xml:"TxLnData"`
}

type LoanData40 struct {
	UnqTradIdr  Max52Text                          `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt       ISODate                            `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp     ExposureType6Code                  `json:"ctrctTp,omitempty" xml:"CtrctTp"`
	ClrSts      Cleared8Choice                     `json:"clrSts,omitempty" xml:"ClrSts"`
	TradgVn     MICIdentifier                      `json:"tradgVn,omitempty" xml:"TradgVn"`
	MstrAgrmt   *MasterAgreement1                  `json:"mstrAgrmt,omitempty" xml:"MstrAgrmt,omitempty"`
	ExctnDtTm   ISODateTime                        `json:"exctnDtTm,omitempty" xml:"ExctnDtTm"`
	ValDt       ISODate                            `json:"valDt,omitempty" xml:"ValDt"`
	TermntnDt   ISODate                            `json:"termntnDt,omitempty" xml:"TermntnDt"`
	GnlColl     *SpecialCollateral1Code            `json:"gnlColl,omitempty" xml:"GnlColl,omitempty"`
	UnitPric    *SecuritiesTransactionPrice2Choice `json:"unitPric,omitempty" xml:"UnitPric,omitempty"`
	TxSpcfcData TransactionLoanData6Choice         `json:"txSpcfcData,omitempty" xml:"TxSpcfcData"`
}

type LoanData41 struct {
	DlvryByVal    bool                           `json:"dlvryByVal,omitempty" xml:"DlvryByVal"`
	CollDlvryMtd  *CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd,omitempty"`
	Term          []ContractTerm2Choice          `json:"term,omitempty" xml:"Term,omitempty"`
	AsstTp        SecurityCommodity2Choice       `json:"asstTp,omitempty" xml:"AsstTp"`
	RbtRate       RebateRate1Choice              `json:"rbtRate,omitempty" xml:"RbtRate"`
	LnVal         string                         `json:"lnVal,omitempty" xml:"LnVal"`
	LndgFee       string                         `json:"lndgFee,omitempty" xml:"LndgFee"`
	ExclsvArrgmnt bool                           `json:"exclsvArrgmnt,omitempty" xml:"ExclsvArrgmnt"`
}

type LoanData42 struct {
	UnqTradIdr  Max52Text                  `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt       ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp     *ExposureType6Code         `json:"ctrctTp,omitempty" xml:"CtrctTp,omitempty"`
	TxSpcfcData TransactionLoanData5Choice `json:"txSpcfcData,omitempty" xml:"TxSpcfcData"`
}

type LoanData43 struct {
	UnqTradIdr  Max52Text                  `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt       ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp     *ExposureType6Code         `json:"ctrctTp,omitempty" xml:"CtrctTp,omitempty"`
	ExctnDtTm   *ISODateTime               `json:"exctnDtTm,omitempty" xml:"ExctnDtTm,omitempty"`
	TermntnDt   *ISODate                   `json:"termntnDt,omitempty" xml:"TermntnDt,omitempty"`
	TxSpcfcData TransactionLoanData5Choice `json:"txSpcfcData,omitempty" xml:"TxSpcfcData"`
}

type LoanData44 struct {
	UnqTradIdr Max52Text                  `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt      ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp    ExposureType6Code          `json:"ctrctTp,omitempty" xml:"CtrctTp"`
	ExctnDtTm  ISODateTime                `json:"exctnDtTm,omitempty" xml:"ExctnDtTm"`
	TxLnData   TransactionLoanData5Choice `json:"txLnData,omitempty" xml:"TxLnData"`
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

type MasterAgreement1 struct {
	Tp                AgreementType1Choice `json:"tp,omitempty" xml:"Tp"`
	Vrsn              *ISORestrictedYear   `json:"vrsn,omitempty" xml:"Vrsn,omitempty"`
	OthrMstrAgrmtDtls *Max50Text           `json:"othrMstrAgrmtDtls,omitempty" xml:"OthrMstrAgrmtDtls,omitempty"`
}

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max50Text string

// Must be at least 1 items long
type Max52Text string

type MetalCommodityNonPrecious1 struct {
	BasePdct     AssetClassProductType7Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType15Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType10Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

type MetalCommodityPrecious1 struct {
	BasePdct     AssetClassProductType7Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType16Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType11Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

// May be one of PSTN, TCTN
type ModificationLevel1Code string

// Must match the pattern [A-U]{1,1}
type NACEDomainIdentifier string

// May be one of NORE
type NoReasonCode string

// May be one of NTAV
type NotAvailable1Code string

type OrganisationIdentification9Choice struct {
	LEI    *LEIIdentifier           `json:"lei,omitempty" xml:"LEI,omitempty"`
	ClntId *Max50Text               `json:"clntId,omitempty" xml:"ClntId,omitempty"`
	AnyBIC *AnyBICDec2014Identifier `json:"anyBIC,omitempty" xml:"AnyBIC,omitempty"`
}

type OtherC10CommodityDeliverable2 struct {
	BasePdct AssetClassProductType11Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType47Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type OtherC10CommodityNonDeliverable2 struct {
	BasePdct AssetClassProductType11Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType48Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PaperCommodityContainerBoard1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType35Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PaperCommodityNewsprint1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType36Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PaperCommodityPulp1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType37Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PaperCommodityRecoveredPaper1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType38Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PaperCommodityRecoveredPaper2 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PolypropyleneCommodityOther1 struct {
	BasePdct AssetClassProductType9Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

type PolypropyleneCommodityPlastic1 struct {
	BasePdct AssetClassProductType9Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType18Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

type PrincipalAmount1 struct {
	ValDtAmt   string                       `json:"valDtAmt,omitempty" xml:"ValDtAmt"`
	MtrtyDtAmt string                       `json:"mtrtyDtAmt,omitempty" xml:"MtrtyDtAmt"`
	Ccy        ActiveOrHistoricCurrencyCode `json:"ccy,omitempty" xml:"Ccy"`
}

type Quantity13 struct {
	Val         string             `json:"val,omitempty" xml:"Val"`
	UnitOfMeasr UnitOfMeasure1Code `json:"unitOfMeasr,omitempty" xml:"UnitOfMeasr"`
}

type QuantityNominalValue1Choice struct {
	Qty     *string                            `json:"qty,omitempty" xml:"Qty,omitempty"`
	NmnlVal *ActiveOrHistoricCurrencyAndAmount `json:"nmnlVal,omitempty" xml:"NmnlVal,omitempty"`
}

type RateAdjustment1 struct {
	Rate       string  `json:"rate,omitempty" xml:"Rate"`
	AdjstmntDt ISODate `json:"adjstmntDt,omitempty" xml:"AdjstmntDt"`
}

// May be one of DAYS, MNTH, WEEK, YEAR
type RateBasis1Code string

type RebateRate1Choice struct {
	Fxd  *string                 `json:"fxd,omitempty" xml:"Fxd,omitempty"`
	Fltg *FloatingInterestRate10 `json:"fltg,omitempty" xml:"Fltg,omitempty"`
}

// May be one of EGRN, EGAE, ETSB, NOAP
type RepoTerminationOption2Code string

type SecuritiesFinancingReportingTransactionReportV01 struct {
	TradData    []TradeTransactionReport6Choice `json:"tradData,omitempty" xml:"TradData"`
	SplmtryData []SupplementaryData1            `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type SecuritiesLendingType3Choice struct {
	Cd    *ExternalSecuritiesLendingType1Code `json:"cd,omitempty" xml:"Cd,omitempty"`
	Prtry *Max35Text                          `json:"prtry,omitempty" xml:"Prtry,omitempty"`
}

type SecuritiesTransactionPrice2Choice struct {
	MntryVal *AmountAndDirection61 `json:"mntryVal,omitempty" xml:"MntryVal,omitempty"`
	Pctg     *string               `json:"pctg,omitempty" xml:"Pctg,omitempty"`
	Yld      *string               `json:"yld,omitempty" xml:"Yld,omitempty"`
	BsisPts  *string               `json:"bsisPts,omitempty" xml:"BsisPts,omitempty"`
}

type Security3 struct {
	Id                *ISINOct2015Identifier             `json:"id,omitempty" xml:"Id,omitempty"`
	ClssfctnTp        *CFIOct2015Identifier              `json:"clssfctnTp,omitempty" xml:"ClssfctnTp,omitempty"`
	QtyOrNmnlVal      *QuantityNominalValue1Choice       `json:"qtyOrNmnlVal,omitempty" xml:"QtyOrNmnlVal,omitempty"`
	UnitPric          *SecuritiesTransactionPrice2Choice `json:"unitPric,omitempty" xml:"UnitPric,omitempty"`
	MktVal            *string                            `json:"mktVal,omitempty" xml:"MktVal,omitempty"`
	HrcutOrMrgn       *string                            `json:"hrcutOrMrgn,omitempty" xml:"HrcutOrMrgn,omitempty"`
	Qlty              *CollateralQualityType1Code        `json:"qlty,omitempty" xml:"Qlty,omitempty"`
	Mtrty             *ISODate                           `json:"mtrty,omitempty" xml:"Mtrty,omitempty"`
	Issr              *SecurityIssuer1                   `json:"issr,omitempty" xml:"Issr,omitempty"`
	Tp                *SecuritiesLendingType3Choice      `json:"tp,omitempty" xml:"Tp,omitempty"`
	AvlblForCollReuse *bool                              `json:"avlblForCollReuse,omitempty" xml:"AvlblForCollReuse,omitempty"`
}

type Security4 struct {
	Id            *ISINOct2015Identifier             `json:"id,omitempty" xml:"Id,omitempty"`
	Clssfctn      *CFIOct2015Identifier              `json:"clssfctn,omitempty" xml:"Clssfctn,omitempty"`
	QtyOrNmnlVal  *QuantityNominalValue1Choice       `json:"qtyOrNmnlVal,omitempty" xml:"QtyOrNmnlVal,omitempty"`
	Qlty          *CollateralQualityType1Code        `json:"qlty,omitempty" xml:"Qlty,omitempty"`
	Mtrty         *ISODate                           `json:"mtrty,omitempty" xml:"Mtrty,omitempty"`
	Issr          *SecurityIssuer1                   `json:"issr,omitempty" xml:"Issr,omitempty"`
	Tp            []SecuritiesLendingType3Choice     `json:"tp,omitempty" xml:"Tp,omitempty"`
	UnitPric      *SecuritiesTransactionPrice2Choice `json:"unitPric,omitempty" xml:"UnitPric,omitempty"`
	ExclsvArrgmnt *bool                              `json:"exclsvArrgmnt,omitempty" xml:"ExclsvArrgmnt,omitempty"`
	MktVal        *string                            `json:"mktVal,omitempty" xml:"MktVal,omitempty"`
}

type SecurityCommodity2Choice struct {
	Scty   *Security4  `json:"scty,omitempty" xml:"Scty,omitempty"`
	Cmmdty *Commodity3 `json:"cmmdty,omitempty" xml:"Cmmdty,omitempty"`
}

type SecurityIdentification26Choice struct {
	Id       *ISINOct2015Identifier `json:"id,omitempty" xml:"Id,omitempty"`
	NotAvlbl *NotAvailable1Code     `json:"notAvlbl,omitempty" xml:"NotAvlbl,omitempty"`
}

type SecurityIssuer1 struct {
	LEI          LEIIdentifier `json:"lei,omitempty" xml:"LEI"`
	JursdctnCtry CountryCode   `json:"jursdctnCtry,omitempty" xml:"JursdctnCtry"`
}

type SettlementParties31Choice struct {
	CntrlSctiesDpstryPtcpt *LEIIdentifier `json:"cntrlSctiesDpstryPtcpt,omitempty" xml:"CntrlSctiesDpstryPtcpt,omitempty"`
	IndrctPtcpt            *LEIIdentifier `json:"indrctPtcpt,omitempty" xml:"IndrctPtcpt,omitempty"`
}

// May be one of GENE, SPEC
type SpecialCollateral1Code string

type SupplementaryData1 struct {
	PlcAndNm *Max350Text                `json:"plcAndNm,omitempty" xml:"PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `json:"envlp,omitempty" xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeEarlyTermination2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData49   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      LoanData11           `json:"lnData,omitempty" xml:"LnData"`
	CollData    CollateralFlag6      `json:"collData,omitempty" xml:"CollData"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeError2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData49   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	UnqTxIdr    Max52Text            `json:"unqTxIdr,omitempty" xml:"UnqTxIdr"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeNewTransaction6 struct {
	TechRcrdId  *Max35Text                        `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48                `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      LoanData44                        `json:"lnData,omitempty" xml:"LnData"`
	CollData    *TransactionCollateralData1Choice `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code            `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1              `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeTransactionCollateralUpdate2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData49   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      LoanData4            `json:"lnData,omitempty" xml:"LnData"`
	CollData    CollateralData7      `json:"collData,omitempty" xml:"CollData"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeTransactionCorrection6 struct {
	TechRcrdId  *Max35Text             `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48     `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      *LoanData43            `json:"lnData,omitempty" xml:"LnData,omitempty"`
	CollData    *CollateralData7       `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1   `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeTransactionModification7 struct {
	TechRcrdId  *Max35Text                        `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48                `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      *LoanData42                       `json:"lnData,omitempty" xml:"LnData,omitempty"`
	CollData    *TransactionCollateralData8Choice `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code            `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1              `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeTransactionPositionComponent2 struct {
	TechRcrdId  *Max35Text             `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48     `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      *LoanData40            `json:"lnData,omitempty" xml:"LnData,omitempty"`
	CollData    *CollateralData3       `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1   `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TradeTransactionReport6Choice struct {
	New          *TradeNewTransaction6               `json:"new,omitempty" xml:"New,omitempty"`
	Mod          *TradeTransactionModification7      `json:"mod,omitempty" xml:"Mod,omitempty"`
	Err          *TradeError2                        `json:"err,omitempty" xml:"Err,omitempty"`
	EarlyTermntn *TradeEarlyTermination2             `json:"earlyTermntn,omitempty" xml:"EarlyTermntn,omitempty"`
	PosCmpnt     *TradeTransactionPositionComponent2 `json:"posCmpnt,omitempty" xml:"PosCmpnt,omitempty"`
	CollUpd      *TradeTransactionCollateralUpdate2  `json:"collUpd,omitempty" xml:"CollUpd,omitempty"`
	Crrctn       *TradeTransactionCorrection6        `json:"crrctn,omitempty" xml:"Crrctn,omitempty"`
	ValtnUpd     *TradeValuationUpdate2              `json:"valtnUpd,omitempty" xml:"ValtnUpd,omitempty"`
}

type TradeValuationUpdate2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	UnqTxIdr    Max52Text            `json:"unqTxIdr,omitempty" xml:"UnqTxIdr"`
	EvtDt       ISODate              `json:"evtDt,omitempty" xml:"EvtDt"`
	MktVal      string               `json:"mktVal,omitempty" xml:"MktVal"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

type TransactionCollateralData1Choice struct {
	RpTrad     *Collateral15          `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *Collateral15          `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *CollateralFlag6Choice `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
	MrgnLndg   *Security3             `json:"mrgnLndg,omitempty" xml:"MrgnLndg,omitempty"`
}

type TransactionCollateralData8Choice struct {
	RpTrad     *Collateral27          `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *Collateral27          `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *CollateralFlag6Choice `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
}

type TransactionCounterpartyData3Choice struct {
	RpTrad     *TransactionCounterpartyData7 `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *TransactionCounterpartyData7 `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *TransactionCounterpartyData7 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
}

type TransactionCounterpartyData7 struct {
	Bnfcry     *OrganisationIdentification9Choice `json:"bnfcry,omitempty" xml:"Bnfcry,omitempty"`
	TrptyAgt   *LEIIdentifier                     `json:"trptyAgt,omitempty" xml:"TrptyAgt,omitempty"`
	Brkr       *LEIIdentifier                     `json:"brkr,omitempty" xml:"Brkr,omitempty"`
	ClrMmb     *LEIIdentifier                     `json:"clrMmb,omitempty" xml:"ClrMmb,omitempty"`
	SttlmPties *SettlementParties31Choice         `json:"sttlmPties,omitempty" xml:"SttlmPties,omitempty"`
	AgtLndr    *LEIIdentifier                     `json:"agtLndr,omitempty" xml:"AgtLndr,omitempty"`
}

type TransactionLoanData4Choice struct {
	RpTrad     *LoanData29 `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *LoanData29 `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *LoanData29 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
	MrgnLndg   *LoanData13 `json:"mrgnLndg,omitempty" xml:"MrgnLndg,omitempty"`
}

type TransactionLoanData5Choice struct {
	RpTrad     *LoanData1  `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *LoanData2  `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *LoanData39 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
	MrgnLndg   *LoanData32 `json:"mrgnLndg,omitempty" xml:"MrgnLndg,omitempty"`
}

type TransactionLoanData6Choice struct {
	RpTrad     *LoanData37 `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *LoanData27 `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *LoanData41 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
}

// May be one of PIEC, TONS, FOOT, GBGA, USGA, GRAM, INCH, KILO, PUND, METR, CMET, MMET, LITR, CELI, MILI, GBOU, USOU, GBQA, USQA, GBPI, USPI, MILE, KMET, YARD, SQKI, HECT, ARES, SMET, SCMT, SMIL, SQMI, SQYA, SQFO, SQIN, ACRE
type UnitOfMeasure1Code string

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006")), nil
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
