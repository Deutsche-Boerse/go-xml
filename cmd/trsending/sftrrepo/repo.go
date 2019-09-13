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

func GetActiveCurrencyCodePointer(in string) *ActiveCurrencyCode {
	if in == "" {
		return nil
	}
	out := ActiveCurrencyCode(in)
	return &out
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value string                       `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

func GetActiveOrHistoricCurrencyCodePointer(in string) *ActiveOrHistoricCurrencyCode {
	if in == "" {
		return nil
	}
	out := ActiveOrHistoricCurrencyCode(in)
	return &out
}

type AgreementType1Choice struct {
	Tp    *ExternalAgreementType1Code `json:"tp,omitempty" xml:"Tp,omitempty"`
	Prtry *Max35Text                  `json:"prtry,omitempty" xml:"Prtry,omitempty"`
}

func (t *AgreementType1Choice) GetTp() (out ExternalAgreementType1Code) {
	if t == nil || t.Tp == nil {
		return
	}
	return *t.Tp
}
func (t *AgreementType1Choice) GetPrtry() (out Max35Text) {
	if t == nil || t.Prtry == nil {
		return
	}
	return *t.Prtry
}

type AgriculturalCommodityDairy1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType20Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *AgriculturalCommodityDairy1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityDairy1) GetSubPdct() (out AssetClassSubProductType20Code) {
	return t.SubPdct
}

type AgriculturalCommodityForestry1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType21Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *AgriculturalCommodityForestry1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityForestry1) GetSubPdct() (out AssetClassSubProductType21Code) {
	return t.SubPdct
}

type AgriculturalCommodityGrain2 struct {
	BasePdct     AssetClassProductType1Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType5Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType30Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *AgriculturalCommodityGrain2) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityGrain2) GetSubPdct() (out AssetClassSubProductType5Code) {
	return t.SubPdct
}
func (t *AgriculturalCommodityGrain2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType30Code) {
	return t.AddtlSubPdct
}

type AgriculturalCommodityLiveStock1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType22Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *AgriculturalCommodityLiveStock1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityLiveStock1) GetSubPdct() (out AssetClassSubProductType22Code) {
	return t.SubPdct
}

type AgriculturalCommodityOilSeed1 struct {
	BasePdct     AssetClassProductType1Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType1Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType1Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *AgriculturalCommodityOilSeed1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityOilSeed1) GetSubPdct() (out AssetClassSubProductType1Code) {
	return t.SubPdct
}
func (t *AgriculturalCommodityOilSeed1) GetAddtlSubPdct() (out AssetClassDetailedSubProductType1Code) {
	return t.AddtlSubPdct
}

type AgriculturalCommodityOliveOil2 struct {
	BasePdct     AssetClassProductType1Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType3Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType29Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *AgriculturalCommodityOliveOil2) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityOliveOil2) GetSubPdct() (out AssetClassSubProductType3Code) {
	return t.SubPdct
}
func (t *AgriculturalCommodityOliveOil2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType29Code) {
	return t.AddtlSubPdct
}

type AgriculturalCommodityOther1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *AgriculturalCommodityOther1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityOther1) GetSubPdct() (out AssetClassSubProductType49Code) {
	return t.SubPdct
}

type AgriculturalCommodityPotato1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType45Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *AgriculturalCommodityPotato1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommodityPotato1) GetSubPdct() (out AssetClassSubProductType45Code) {
	return t.SubPdct
}

type AgriculturalCommoditySeafood1 struct {
	BasePdct AssetClassProductType1Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType23Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *AgriculturalCommoditySeafood1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommoditySeafood1) GetSubPdct() (out AssetClassSubProductType23Code) {
	return t.SubPdct
}

type AgriculturalCommoditySoft1 struct {
	BasePdct     AssetClassProductType1Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType2Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType2Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *AgriculturalCommoditySoft1) GetBasePdct() (out AssetClassProductType1Code) {
	return t.BasePdct
}
func (t *AgriculturalCommoditySoft1) GetSubPdct() (out AssetClassSubProductType2Code) {
	return t.SubPdct
}
func (t *AgriculturalCommoditySoft1) GetAddtlSubPdct() (out AssetClassDetailedSubProductType2Code) {
	return t.AddtlSubPdct
}

type AmountAndDirection61 struct {
	Amt ActiveCurrencyAnd13DecimalAmount `json:"amt,omitempty" xml:"Amt"`
	Sgn *bool                            `json:"sgn,omitempty" xml:"Sgn,omitempty"`
}

func (t *AmountAndDirection61) GetAmt() (out ActiveCurrencyAnd13DecimalAmount) {
	return t.Amt
}
func (t *AmountAndDirection61) GetSgn() (out bool) {
	if t == nil || t.Sgn == nil {
		return
	}
	return *t.Sgn
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

func GetAnyBICDec2014IdentifierPointer(in string) *AnyBICDec2014Identifier {
	if in == "" {
		return nil
	}
	out := AnyBICDec2014Identifier(in)
	return &out
}

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

func (t *AssetClassCommodity5Choice) GetAgrcltrl() (out AssetClassCommodityAgricultural5Choice) {
	if t == nil || t.Agrcltrl == nil {
		return
	}
	return *t.Agrcltrl
}
func (t *AssetClassCommodity5Choice) GetNrgy() (out AssetClassCommodityEnergy2Choice) {
	if t == nil || t.Nrgy == nil {
		return
	}
	return *t.Nrgy
}
func (t *AssetClassCommodity5Choice) GetEnvttl() (out AssetClassCommodityEnvironmental2Choice) {
	if t == nil || t.Envttl == nil {
		return
	}
	return *t.Envttl
}
func (t *AssetClassCommodity5Choice) GetFrtlzr() (out AssetClassCommodityFertilizer3Choice) {
	if t == nil || t.Frtlzr == nil {
		return
	}
	return *t.Frtlzr
}
func (t *AssetClassCommodity5Choice) GetFrght() (out AssetClassCommodityFreight3Choice) {
	if t == nil || t.Frght == nil {
		return
	}
	return *t.Frght
}
func (t *AssetClassCommodity5Choice) GetIndstrlPdct() (out AssetClassCommodityIndustrialProduct1Choice) {
	if t == nil || t.IndstrlPdct == nil {
		return
	}
	return *t.IndstrlPdct
}
func (t *AssetClassCommodity5Choice) GetMetl() (out AssetClassCommodityMetal1Choice) {
	if t == nil || t.Metl == nil {
		return
	}
	return *t.Metl
}
func (t *AssetClassCommodity5Choice) GetOthrC10() (out AssetClassCommodityOtherC102Choice) {
	if t == nil || t.OthrC10 == nil {
		return
	}
	return *t.OthrC10
}
func (t *AssetClassCommodity5Choice) GetPpr() (out AssetClassCommodityPaper3Choice) {
	if t == nil || t.Ppr == nil {
		return
	}
	return *t.Ppr
}
func (t *AssetClassCommodity5Choice) GetPlprpln() (out AssetClassCommodityPolypropylene3Choice) {
	if t == nil || t.Plprpln == nil {
		return
	}
	return *t.Plprpln
}
func (t *AssetClassCommodity5Choice) GetInfltn() (out AssetClassCommodityInflation1) {
	if t == nil || t.Infltn == nil {
		return
	}
	return *t.Infltn
}
func (t *AssetClassCommodity5Choice) GetMultiCmmdtyExtc() (out AssetClassCommodityMultiCommodityExotic1) {
	if t == nil || t.MultiCmmdtyExtc == nil {
		return
	}
	return *t.MultiCmmdtyExtc
}
func (t *AssetClassCommodity5Choice) GetOffclEcnmcSttstcs() (out AssetClassCommodityOfficialEconomicStatistics1) {
	if t == nil || t.OffclEcnmcSttstcs == nil {
		return
	}
	return *t.OffclEcnmcSttstcs
}
func (t *AssetClassCommodity5Choice) GetOthr() (out AssetClassCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
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

func (t *AssetClassCommodityAgricultural5Choice) GetGrnOilSeed() (out AgriculturalCommodityOilSeed1) {
	if t == nil || t.GrnOilSeed == nil {
		return
	}
	return *t.GrnOilSeed
}
func (t *AssetClassCommodityAgricultural5Choice) GetSoft() (out AgriculturalCommoditySoft1) {
	if t == nil || t.Soft == nil {
		return
	}
	return *t.Soft
}
func (t *AssetClassCommodityAgricultural5Choice) GetPtt() (out AgriculturalCommodityPotato1) {
	if t == nil || t.Ptt == nil {
		return
	}
	return *t.Ptt
}
func (t *AssetClassCommodityAgricultural5Choice) GetOlvOil() (out AgriculturalCommodityOliveOil2) {
	if t == nil || t.OlvOil == nil {
		return
	}
	return *t.OlvOil
}
func (t *AssetClassCommodityAgricultural5Choice) GetDairy() (out AgriculturalCommodityDairy1) {
	if t == nil || t.Dairy == nil {
		return
	}
	return *t.Dairy
}
func (t *AssetClassCommodityAgricultural5Choice) GetFrstry() (out AgriculturalCommodityForestry1) {
	if t == nil || t.Frstry == nil {
		return
	}
	return *t.Frstry
}
func (t *AssetClassCommodityAgricultural5Choice) GetSfd() (out AgriculturalCommoditySeafood1) {
	if t == nil || t.Sfd == nil {
		return
	}
	return *t.Sfd
}
func (t *AssetClassCommodityAgricultural5Choice) GetLiveStock() (out AgriculturalCommodityLiveStock1) {
	if t == nil || t.LiveStock == nil {
		return
	}
	return *t.LiveStock
}
func (t *AssetClassCommodityAgricultural5Choice) GetGrn() (out AgriculturalCommodityGrain2) {
	if t == nil || t.Grn == nil {
		return
	}
	return *t.Grn
}
func (t *AssetClassCommodityAgricultural5Choice) GetOthr() (out AgriculturalCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
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

func (t *AssetClassCommodityEnergy2Choice) GetElctrcty() (out EnergyCommodityElectricity1) {
	if t == nil || t.Elctrcty == nil {
		return
	}
	return *t.Elctrcty
}
func (t *AssetClassCommodityEnergy2Choice) GetNtrlGas() (out EnergyCommodityNaturalGas2) {
	if t == nil || t.NtrlGas == nil {
		return
	}
	return *t.NtrlGas
}
func (t *AssetClassCommodityEnergy2Choice) GetOil() (out EnergyCommodityOil2) {
	if t == nil || t.Oil == nil {
		return
	}
	return *t.Oil
}
func (t *AssetClassCommodityEnergy2Choice) GetCoal() (out EnergyCommodityCoal1) {
	if t == nil || t.Coal == nil {
		return
	}
	return *t.Coal
}
func (t *AssetClassCommodityEnergy2Choice) GetIntrNrgy() (out EnergyCommodityInterEnergy1) {
	if t == nil || t.IntrNrgy == nil {
		return
	}
	return *t.IntrNrgy
}
func (t *AssetClassCommodityEnergy2Choice) GetRnwblNrgy() (out EnergyCommodityRenewableEnergy1) {
	if t == nil || t.RnwblNrgy == nil {
		return
	}
	return *t.RnwblNrgy
}
func (t *AssetClassCommodityEnergy2Choice) GetLghtEnd() (out EnergyCommodityLightEnd1) {
	if t == nil || t.LghtEnd == nil {
		return
	}
	return *t.LghtEnd
}
func (t *AssetClassCommodityEnergy2Choice) GetDstllts() (out EnergyCommodityDistillates1) {
	if t == nil || t.Dstllts == nil {
		return
	}
	return *t.Dstllts
}
func (t *AssetClassCommodityEnergy2Choice) GetOthr() (out EnergyCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
}

type AssetClassCommodityEnvironmental2Choice struct {
	Emssns   *EnvironmentalCommodityEmission2      `json:"emssns,omitempty" xml:"Emssns,omitempty"`
	Wthr     *EnvironmentalCommodityWeather1       `json:"wthr,omitempty" xml:"Wthr,omitempty"`
	CrbnRltd *EnvironmentalCommodityCarbonRelated1 `json:"crbnRltd,omitempty" xml:"CrbnRltd,omitempty"`
	Othr     *EnvironmentCommodityOther1           `json:"othr,omitempty" xml:"Othr,omitempty"`
}

func (t *AssetClassCommodityEnvironmental2Choice) GetEmssns() (out EnvironmentalCommodityEmission2) {
	if t == nil || t.Emssns == nil {
		return
	}
	return *t.Emssns
}
func (t *AssetClassCommodityEnvironmental2Choice) GetWthr() (out EnvironmentalCommodityWeather1) {
	if t == nil || t.Wthr == nil {
		return
	}
	return *t.Wthr
}
func (t *AssetClassCommodityEnvironmental2Choice) GetCrbnRltd() (out EnvironmentalCommodityCarbonRelated1) {
	if t == nil || t.CrbnRltd == nil {
		return
	}
	return *t.CrbnRltd
}
func (t *AssetClassCommodityEnvironmental2Choice) GetOthr() (out EnvironmentCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
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

func (t *AssetClassCommodityFertilizer3Choice) GetAmmn() (out FertilizerCommodityAmmonia1) {
	if t == nil || t.Ammn == nil {
		return
	}
	return *t.Ammn
}
func (t *AssetClassCommodityFertilizer3Choice) GetDmmnmPhspht() (out FertilizerCommodityDiammoniumPhosphate1) {
	if t == nil || t.DmmnmPhspht == nil {
		return
	}
	return *t.DmmnmPhspht
}
func (t *AssetClassCommodityFertilizer3Choice) GetPtsh() (out FertilizerCommodityPotash1) {
	if t == nil || t.Ptsh == nil {
		return
	}
	return *t.Ptsh
}
func (t *AssetClassCommodityFertilizer3Choice) GetSlphr() (out FertilizerCommoditySulphur1) {
	if t == nil || t.Slphr == nil {
		return
	}
	return *t.Slphr
}
func (t *AssetClassCommodityFertilizer3Choice) GetUrea() (out FertilizerCommodityUrea1) {
	if t == nil || t.Urea == nil {
		return
	}
	return *t.Urea
}
func (t *AssetClassCommodityFertilizer3Choice) GetUreaAndAmmnmNtrt() (out FertilizerCommodityUreaAndAmmoniumNitrate1) {
	if t == nil || t.UreaAndAmmnmNtrt == nil {
		return
	}
	return *t.UreaAndAmmnmNtrt
}
func (t *AssetClassCommodityFertilizer3Choice) GetOthr() (out FertilizerCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
}

type AssetClassCommodityFreight3Choice struct {
	Dry       *FreightCommodityDry2           `json:"dry,omitempty" xml:"Dry,omitempty"`
	Wet       *FreightCommodityWet2           `json:"wet,omitempty" xml:"Wet,omitempty"`
	CntnrShip *FreightCommodityContainerShip1 `json:"cntnrShip,omitempty" xml:"CntnrShip,omitempty"`
	Othr      *FreightCommodityOther1         `json:"othr,omitempty" xml:"Othr,omitempty"`
}

func (t *AssetClassCommodityFreight3Choice) GetDry() (out FreightCommodityDry2) {
	if t == nil || t.Dry == nil {
		return
	}
	return *t.Dry
}
func (t *AssetClassCommodityFreight3Choice) GetWet() (out FreightCommodityWet2) {
	if t == nil || t.Wet == nil {
		return
	}
	return *t.Wet
}
func (t *AssetClassCommodityFreight3Choice) GetCntnrShip() (out FreightCommodityContainerShip1) {
	if t == nil || t.CntnrShip == nil {
		return
	}
	return *t.CntnrShip
}
func (t *AssetClassCommodityFreight3Choice) GetOthr() (out FreightCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
}

type AssetClassCommodityIndustrialProduct1Choice struct {
	Cnstrctn *IndustrialProductCommodityConstruction1  `json:"cnstrctn,omitempty" xml:"Cnstrctn,omitempty"`
	Manfctg  *IndustrialProductCommodityManufacturing1 `json:"manfctg,omitempty" xml:"Manfctg,omitempty"`
}

func (t *AssetClassCommodityIndustrialProduct1Choice) GetCnstrctn() (out IndustrialProductCommodityConstruction1) {
	if t == nil || t.Cnstrctn == nil {
		return
	}
	return *t.Cnstrctn
}
func (t *AssetClassCommodityIndustrialProduct1Choice) GetManfctg() (out IndustrialProductCommodityManufacturing1) {
	if t == nil || t.Manfctg == nil {
		return
	}
	return *t.Manfctg
}

type AssetClassCommodityInflation1 struct {
	BasePdct AssetClassProductType12Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

func (t *AssetClassCommodityInflation1) GetBasePdct() (out AssetClassProductType12Code) {
	return t.BasePdct
}

type AssetClassCommodityMetal1Choice struct {
	NonPrcs *MetalCommodityNonPrecious1 `json:"nonPrcs,omitempty" xml:"NonPrcs,omitempty"`
	Prcs    *MetalCommodityPrecious1    `json:"prcs,omitempty" xml:"Prcs,omitempty"`
}

func (t *AssetClassCommodityMetal1Choice) GetNonPrcs() (out MetalCommodityNonPrecious1) {
	if t == nil || t.NonPrcs == nil {
		return
	}
	return *t.NonPrcs
}
func (t *AssetClassCommodityMetal1Choice) GetPrcs() (out MetalCommodityPrecious1) {
	if t == nil || t.Prcs == nil {
		return
	}
	return *t.Prcs
}

type AssetClassCommodityMultiCommodityExotic1 struct {
	BasePdct AssetClassProductType13Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

func (t *AssetClassCommodityMultiCommodityExotic1) GetBasePdct() (out AssetClassProductType13Code) {
	return t.BasePdct
}

type AssetClassCommodityOfficialEconomicStatistics1 struct {
	BasePdct AssetClassProductType14Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

func (t *AssetClassCommodityOfficialEconomicStatistics1) GetBasePdct() (out AssetClassProductType14Code) {
	return t.BasePdct
}

type AssetClassCommodityOther1 struct {
	BasePdct AssetClassProductType15Code `json:"basePdct,omitempty" xml:"BasePdct"`
}

func (t *AssetClassCommodityOther1) GetBasePdct() (out AssetClassProductType15Code) {
	return t.BasePdct
}

type AssetClassCommodityOtherC102Choice struct {
	Dlvrbl    *OtherC10CommodityDeliverable2    `json:"dlvrbl,omitempty" xml:"Dlvrbl,omitempty"`
	NonDlvrbl *OtherC10CommodityNonDeliverable2 `json:"nonDlvrbl,omitempty" xml:"NonDlvrbl,omitempty"`
}

func (t *AssetClassCommodityOtherC102Choice) GetDlvrbl() (out OtherC10CommodityDeliverable2) {
	if t == nil || t.Dlvrbl == nil {
		return
	}
	return *t.Dlvrbl
}
func (t *AssetClassCommodityOtherC102Choice) GetNonDlvrbl() (out OtherC10CommodityNonDeliverable2) {
	if t == nil || t.NonDlvrbl == nil {
		return
	}
	return *t.NonDlvrbl
}

type AssetClassCommodityPaper3Choice struct {
	CntnrBrd *PaperCommodityContainerBoard1 `json:"cntnrBrd,omitempty" xml:"CntnrBrd,omitempty"`
	Nwsprnt  *PaperCommodityNewsprint1      `json:"nwsprnt,omitempty" xml:"Nwsprnt,omitempty"`
	Pulp     *PaperCommodityPulp1           `json:"pulp,omitempty" xml:"Pulp,omitempty"`
	RcvrdPpr *PaperCommodityRecoveredPaper1 `json:"rcvrdPpr,omitempty" xml:"RcvrdPpr,omitempty"`
	Othr     *PaperCommodityRecoveredPaper2 `json:"othr,omitempty" xml:"Othr,omitempty"`
}

func (t *AssetClassCommodityPaper3Choice) GetCntnrBrd() (out PaperCommodityContainerBoard1) {
	if t == nil || t.CntnrBrd == nil {
		return
	}
	return *t.CntnrBrd
}
func (t *AssetClassCommodityPaper3Choice) GetNwsprnt() (out PaperCommodityNewsprint1) {
	if t == nil || t.Nwsprnt == nil {
		return
	}
	return *t.Nwsprnt
}
func (t *AssetClassCommodityPaper3Choice) GetPulp() (out PaperCommodityPulp1) {
	if t == nil || t.Pulp == nil {
		return
	}
	return *t.Pulp
}
func (t *AssetClassCommodityPaper3Choice) GetRcvrdPpr() (out PaperCommodityRecoveredPaper1) {
	if t == nil || t.RcvrdPpr == nil {
		return
	}
	return *t.RcvrdPpr
}
func (t *AssetClassCommodityPaper3Choice) GetOthr() (out PaperCommodityRecoveredPaper2) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
}

type AssetClassCommodityPolypropylene3Choice struct {
	Plstc *PolypropyleneCommodityPlastic1 `json:"plstc,omitempty" xml:"Plstc,omitempty"`
	Othr  *PolypropyleneCommodityOther1   `json:"othr,omitempty" xml:"Othr,omitempty"`
}

func (t *AssetClassCommodityPolypropylene3Choice) GetPlstc() (out PolypropyleneCommodityPlastic1) {
	if t == nil || t.Plstc == nil {
		return
	}
	return *t.Plstc
}
func (t *AssetClassCommodityPolypropylene3Choice) GetOthr() (out PolypropyleneCommodityOther1) {
	if t == nil || t.Othr == nil {
		return
	}
	return *t.Othr
}

// May be one of ALUM, ALUA, CBLT, COPR, IRON, MOLY, NASC, NICK, STEL, TINN, ZINC, OTHR, LEAD
type AssetClassDetailedSubProductType10Code string

func GetAssetClassDetailedSubProductType10CodePointer(in string) *AssetClassDetailedSubProductType10Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType10Code(in)
	return &out
}

// May be one of GOLD, OTHR, PLDM, PTNM, SLVR
type AssetClassDetailedSubProductType11Code string

func GetAssetClassDetailedSubProductType11CodePointer(in string) *AssetClassDetailedSubProductType11Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType11Code(in)
	return &out
}

// May be one of FWHT, SOYB, RPSD, OTHR, CORN, RICE
type AssetClassDetailedSubProductType1Code string

func GetAssetClassDetailedSubProductType1CodePointer(in string) *AssetClassDetailedSubProductType1Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType1Code(in)
	return &out
}

// May be one of LAMP, OTHR
type AssetClassDetailedSubProductType29Code string

func GetAssetClassDetailedSubProductType29CodePointer(in string) *AssetClassDetailedSubProductType29Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType29Code(in)
	return &out
}

// May be one of ROBU, CCOA, BRWN, WHSG, OTHR
type AssetClassDetailedSubProductType2Code string

func GetAssetClassDetailedSubProductType2CodePointer(in string) *AssetClassDetailedSubProductType2Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType2Code(in)
	return &out
}

// May be one of MWHT, OTHR
type AssetClassDetailedSubProductType30Code string

func GetAssetClassDetailedSubProductType30CodePointer(in string) *AssetClassDetailedSubProductType30Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType30Code(in)
	return &out
}

// May be one of GASP, LNGG, NCGG, TTFG, NBPG, OTHR
type AssetClassDetailedSubProductType31Code string

func GetAssetClassDetailedSubProductType31CodePointer(in string) *AssetClassDetailedSubProductType31Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType31Code(in)
	return &out
}

// May be one of BAKK, BDSL, BRNT, BRNX, CNDA, COND, DSEL, DUBA, ESPO, ETHA, FUEL, FOIL, GOIL, GSLN, HEAT, JTFL, KERO, LLSO, MARS, NAPH, NGLO, TAPI, WTIO, URAL, OTHR
type AssetClassDetailedSubProductType32Code string

func GetAssetClassDetailedSubProductType32CodePointer(in string) *AssetClassDetailedSubProductType32Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType32Code(in)
	return &out
}

// May be one of DBCR, OTHR
type AssetClassDetailedSubProductType33Code string

func GetAssetClassDetailedSubProductType33CodePointer(in string) *AssetClassDetailedSubProductType33Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType33Code(in)
	return &out
}

// May be one of TNKR, OTHR
type AssetClassDetailedSubProductType34Code string

func GetAssetClassDetailedSubProductType34CodePointer(in string) *AssetClassDetailedSubProductType34Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType34Code(in)
	return &out
}

// May be one of BSLD, FITR, PKLD, OFFP, OTHR
type AssetClassDetailedSubProductType5Code string

func GetAssetClassDetailedSubProductType5CodePointer(in string) *AssetClassDetailedSubProductType5Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType5Code(in)
	return &out
}

// May be one of CERE, ERUE, EUAE, EUAA, OTHR
type AssetClassDetailedSubProductType8Code string

func GetAssetClassDetailedSubProductType8CodePointer(in string) *AssetClassDetailedSubProductType8Code {
	if in == "" {
		return nil
	}
	out := AssetClassDetailedSubProductType8Code(in)
	return &out
}

// May be one of OTHC
type AssetClassProductType11Code string

func GetAssetClassProductType11CodePointer(in string) *AssetClassProductType11Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType11Code(in)
	return &out
}

// May be one of INFL
type AssetClassProductType12Code string

func GetAssetClassProductType12CodePointer(in string) *AssetClassProductType12Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType12Code(in)
	return &out
}

// May be one of MCEX
type AssetClassProductType13Code string

func GetAssetClassProductType13CodePointer(in string) *AssetClassProductType13Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType13Code(in)
	return &out
}

// May be one of OEST
type AssetClassProductType14Code string

func GetAssetClassProductType14CodePointer(in string) *AssetClassProductType14Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType14Code(in)
	return &out
}

// May be one of OTHR
type AssetClassProductType15Code string

func GetAssetClassProductType15CodePointer(in string) *AssetClassProductType15Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType15Code(in)
	return &out
}

// May be one of AGRI
type AssetClassProductType1Code string

func GetAssetClassProductType1CodePointer(in string) *AssetClassProductType1Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType1Code(in)
	return &out
}

// May be one of NRGY
type AssetClassProductType2Code string

func GetAssetClassProductType2CodePointer(in string) *AssetClassProductType2Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType2Code(in)
	return &out
}

// May be one of ENVR
type AssetClassProductType3Code string

func GetAssetClassProductType3CodePointer(in string) *AssetClassProductType3Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType3Code(in)
	return &out
}

// May be one of FRGT
type AssetClassProductType4Code string

func GetAssetClassProductType4CodePointer(in string) *AssetClassProductType4Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType4Code(in)
	return &out
}

// May be one of FRTL
type AssetClassProductType5Code string

func GetAssetClassProductType5CodePointer(in string) *AssetClassProductType5Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType5Code(in)
	return &out
}

// May be one of INDP
type AssetClassProductType6Code string

func GetAssetClassProductType6CodePointer(in string) *AssetClassProductType6Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType6Code(in)
	return &out
}

// May be one of METL
type AssetClassProductType7Code string

func GetAssetClassProductType7CodePointer(in string) *AssetClassProductType7Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType7Code(in)
	return &out
}

// May be one of PAPR
type AssetClassProductType8Code string

func GetAssetClassProductType8CodePointer(in string) *AssetClassProductType8Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType8Code(in)
	return &out
}

// May be one of POLY
type AssetClassProductType9Code string

func GetAssetClassProductType9CodePointer(in string) *AssetClassProductType9Code {
	if in == "" {
		return nil
	}
	out := AssetClassProductType9Code(in)
	return &out
}

// May be one of EMIS
type AssetClassSubProductType10Code string

func GetAssetClassSubProductType10CodePointer(in string) *AssetClassSubProductType10Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType10Code(in)
	return &out
}

// May be one of NPRM
type AssetClassSubProductType15Code string

func GetAssetClassSubProductType15CodePointer(in string) *AssetClassSubProductType15Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType15Code(in)
	return &out
}

// May be one of PRME
type AssetClassSubProductType16Code string

func GetAssetClassSubProductType16CodePointer(in string) *AssetClassSubProductType16Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType16Code(in)
	return &out
}

// May be one of PLST
type AssetClassSubProductType18Code string

func GetAssetClassSubProductType18CodePointer(in string) *AssetClassSubProductType18Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType18Code(in)
	return &out
}

// May be one of GROS
type AssetClassSubProductType1Code string

func GetAssetClassSubProductType1CodePointer(in string) *AssetClassSubProductType1Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType1Code(in)
	return &out
}

// May be one of DIRY
type AssetClassSubProductType20Code string

func GetAssetClassSubProductType20CodePointer(in string) *AssetClassSubProductType20Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType20Code(in)
	return &out
}

// May be one of FRST
type AssetClassSubProductType21Code string

func GetAssetClassSubProductType21CodePointer(in string) *AssetClassSubProductType21Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType21Code(in)
	return &out
}

// May be one of LSTK
type AssetClassSubProductType22Code string

func GetAssetClassSubProductType22CodePointer(in string) *AssetClassSubProductType22Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType22Code(in)
	return &out
}

// May be one of SEAF
type AssetClassSubProductType23Code string

func GetAssetClassSubProductType23CodePointer(in string) *AssetClassSubProductType23Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType23Code(in)
	return &out
}

// May be one of COAL
type AssetClassSubProductType24Code string

func GetAssetClassSubProductType24CodePointer(in string) *AssetClassSubProductType24Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType24Code(in)
	return &out
}

// May be one of DIST
type AssetClassSubProductType25Code string

func GetAssetClassSubProductType25CodePointer(in string) *AssetClassSubProductType25Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType25Code(in)
	return &out
}

// May be one of INRG
type AssetClassSubProductType26Code string

func GetAssetClassSubProductType26CodePointer(in string) *AssetClassSubProductType26Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType26Code(in)
	return &out
}

// May be one of LGHT
type AssetClassSubProductType27Code string

func GetAssetClassSubProductType27CodePointer(in string) *AssetClassSubProductType27Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType27Code(in)
	return &out
}

// May be one of RNNG
type AssetClassSubProductType28Code string

func GetAssetClassSubProductType28CodePointer(in string) *AssetClassSubProductType28Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType28Code(in)
	return &out
}

// May be one of CRBR
type AssetClassSubProductType29Code string

func GetAssetClassSubProductType29CodePointer(in string) *AssetClassSubProductType29Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType29Code(in)
	return &out
}

// May be one of SOFT
type AssetClassSubProductType2Code string

func GetAssetClassSubProductType2CodePointer(in string) *AssetClassSubProductType2Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType2Code(in)
	return &out
}

// May be one of WTHR
type AssetClassSubProductType30Code string

func GetAssetClassSubProductType30CodePointer(in string) *AssetClassSubProductType30Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType30Code(in)
	return &out
}

// May be one of DRYF
type AssetClassSubProductType31Code string

func GetAssetClassSubProductType31CodePointer(in string) *AssetClassSubProductType31Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType31Code(in)
	return &out
}

// May be one of WETF
type AssetClassSubProductType32Code string

func GetAssetClassSubProductType32CodePointer(in string) *AssetClassSubProductType32Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType32Code(in)
	return &out
}

// May be one of CSTR
type AssetClassSubProductType33Code string

func GetAssetClassSubProductType33CodePointer(in string) *AssetClassSubProductType33Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType33Code(in)
	return &out
}

// May be one of MFTG
type AssetClassSubProductType34Code string

func GetAssetClassSubProductType34CodePointer(in string) *AssetClassSubProductType34Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType34Code(in)
	return &out
}

// May be one of CBRD
type AssetClassSubProductType35Code string

func GetAssetClassSubProductType35CodePointer(in string) *AssetClassSubProductType35Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType35Code(in)
	return &out
}

// May be one of NSPT
type AssetClassSubProductType36Code string

func GetAssetClassSubProductType36CodePointer(in string) *AssetClassSubProductType36Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType36Code(in)
	return &out
}

// May be one of PULP
type AssetClassSubProductType37Code string

func GetAssetClassSubProductType37CodePointer(in string) *AssetClassSubProductType37Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType37Code(in)
	return &out
}

// May be one of RCVP
type AssetClassSubProductType38Code string

func GetAssetClassSubProductType38CodePointer(in string) *AssetClassSubProductType38Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType38Code(in)
	return &out
}

// May be one of AMMO
type AssetClassSubProductType39Code string

func GetAssetClassSubProductType39CodePointer(in string) *AssetClassSubProductType39Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType39Code(in)
	return &out
}

// May be one of OOLI
type AssetClassSubProductType3Code string

func GetAssetClassSubProductType3CodePointer(in string) *AssetClassSubProductType3Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType3Code(in)
	return &out
}

// May be one of DAPH
type AssetClassSubProductType40Code string

func GetAssetClassSubProductType40CodePointer(in string) *AssetClassSubProductType40Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType40Code(in)
	return &out
}

// May be one of PTSH
type AssetClassSubProductType41Code string

func GetAssetClassSubProductType41CodePointer(in string) *AssetClassSubProductType41Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType41Code(in)
	return &out
}

// May be one of SLPH
type AssetClassSubProductType42Code string

func GetAssetClassSubProductType42CodePointer(in string) *AssetClassSubProductType42Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType42Code(in)
	return &out
}

// May be one of UREA
type AssetClassSubProductType43Code string

func GetAssetClassSubProductType43CodePointer(in string) *AssetClassSubProductType43Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType43Code(in)
	return &out
}

// May be one of UAAN
type AssetClassSubProductType44Code string

func GetAssetClassSubProductType44CodePointer(in string) *AssetClassSubProductType44Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType44Code(in)
	return &out
}

// May be one of POTA
type AssetClassSubProductType45Code string

func GetAssetClassSubProductType45CodePointer(in string) *AssetClassSubProductType45Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType45Code(in)
	return &out
}

// May be one of CSHP
type AssetClassSubProductType46Code string

func GetAssetClassSubProductType46CodePointer(in string) *AssetClassSubProductType46Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType46Code(in)
	return &out
}

// May be one of DLVR
type AssetClassSubProductType47Code string

func GetAssetClassSubProductType47CodePointer(in string) *AssetClassSubProductType47Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType47Code(in)
	return &out
}

// May be one of NDLV
type AssetClassSubProductType48Code string

func GetAssetClassSubProductType48CodePointer(in string) *AssetClassSubProductType48Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType48Code(in)
	return &out
}

// May be one of OTHR
type AssetClassSubProductType49Code string

func GetAssetClassSubProductType49CodePointer(in string) *AssetClassSubProductType49Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType49Code(in)
	return &out
}

// May be one of GRIN
type AssetClassSubProductType5Code string

func GetAssetClassSubProductType5CodePointer(in string) *AssetClassSubProductType5Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType5Code(in)
	return &out
}

// May be one of ELEC
type AssetClassSubProductType6Code string

func GetAssetClassSubProductType6CodePointer(in string) *AssetClassSubProductType6Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType6Code(in)
	return &out
}

// May be one of NGAS
type AssetClassSubProductType7Code string

func GetAssetClassSubProductType7CodePointer(in string) *AssetClassSubProductType7Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType7Code(in)
	return &out
}

// May be one of OILP
type AssetClassSubProductType8Code string

func GetAssetClassSubProductType8CodePointer(in string) *AssetClassSubProductType8Code {
	if in == "" {
		return nil
	}
	out := AssetClassSubProductType8Code(in)
	return &out
}

// May be one of WIBO, TREA, TIBO, TLBO, SWAP, STBO, PRBO, PFAN, NIBO, MAAA, MOSP, LIBO, LIBI, JIBA, ISDA, GCFR, FUSW, EUCH, EUUS, EURI, EONS, EONA, CIBO, CDOR, BUBO, BBSW
type BenchmarkCurveName2Code string

func GetBenchmarkCurveName2CodePointer(in string) *BenchmarkCurveName2Code {
	if in == "" {
		return nil
	}
	out := BenchmarkCurveName2Code(in)
	return &out
}

type BenchmarkCurveName8Choice struct {
	Indx *BenchmarkCurveName2Code `json:"indx,omitempty" xml:"Indx,omitempty"`
	Nm   *Max350Text              `json:"nm,omitempty" xml:"Nm,omitempty"`
}

func (t *BenchmarkCurveName8Choice) GetIndx() (out BenchmarkCurveName2Code) {
	if t == nil || t.Indx == nil {
		return
	}
	return *t.Indx
}
func (t *BenchmarkCurveName8Choice) GetNm() (out Max350Text) {
	if t == nil || t.Nm == nil {
		return
	}
	return *t.Nm
}

type Branch2Choice struct {
	Id   *OrganisationIdentification9Choice `json:"id,omitempty" xml:"Id,omitempty"`
	Ctry *CountryCode                       `json:"ctry,omitempty" xml:"Ctry,omitempty"`
}

func (t *Branch2Choice) GetId() (out OrganisationIdentification9Choice) {
	if t == nil || t.Id == nil {
		return
	}
	return *t.Id
}
func (t *Branch2Choice) GetCtry() (out CountryCode) {
	if t == nil || t.Ctry == nil {
		return
	}
	return *t.Ctry
}

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

func GetCFIOct2015IdentifierPointer(in string) *CFIOct2015Identifier {
	if in == "" {
		return nil
	}
	out := CFIOct2015Identifier(in)
	return &out
}

type Cleared2Choice struct {
	Clrd    *ClearingPartyAndTime7 `json:"clrd,omitempty" xml:"Clrd,omitempty"`
	NonClrd *NoReasonCode          `json:"nonClrd,omitempty" xml:"NonClrd,omitempty"`
}

func (t *Cleared2Choice) GetClrd() (out ClearingPartyAndTime7) {
	if t == nil || t.Clrd == nil {
		return
	}
	return *t.Clrd
}
func (t *Cleared2Choice) GetNonClrd() (out NoReasonCode) {
	if t == nil || t.NonClrd == nil {
		return
	}
	return *t.NonClrd
}

type Cleared8Choice struct {
	Clrd    *ClearingPartyAndTime7 `json:"clrd,omitempty" xml:"Clrd,omitempty"`
	NonClrd *NoReasonCode          `json:"nonClrd,omitempty" xml:"NonClrd,omitempty"`
}

func (t *Cleared8Choice) GetClrd() (out ClearingPartyAndTime7) {
	if t == nil || t.Clrd == nil {
		return
	}
	return *t.Clrd
}
func (t *Cleared8Choice) GetNonClrd() (out NoReasonCode) {
	if t == nil || t.NonClrd == nil {
		return
	}
	return *t.NonClrd
}

type ClearingPartyAndTime7 struct {
	CCP        *LEIIdentifier `json:"ccp,omitempty" xml:"CCP,omitempty"`
	ClrDtTm    *ISODateTime   `json:"clrDtTm,omitempty" xml:"ClrDtTm,omitempty"`
	RptTrckgNb *Max52Text     `json:"rptTrckgNb,omitempty" xml:"RptTrckgNb,omitempty"`
	PrtflCd    *Max52Text     `json:"prtflCd,omitempty" xml:"PrtflCd,omitempty"`
}

func (t *ClearingPartyAndTime7) GetCCP() (out LEIIdentifier) {
	if t == nil || t.CCP == nil {
		return
	}
	return *t.CCP
}
func (t *ClearingPartyAndTime7) GetClrDtTm() (out ISODateTime) {
	if t == nil || t.ClrDtTm == nil {
		return
	}
	return *t.ClrDtTm
}
func (t *ClearingPartyAndTime7) GetRptTrckgNb() (out Max52Text) {
	if t == nil || t.RptTrckgNb == nil {
		return
	}
	return *t.RptTrckgNb
}
func (t *ClearingPartyAndTime7) GetPrtflCd() (out Max52Text) {
	if t == nil || t.PrtflCd == nil {
		return
	}
	return *t.PrtflCd
}

type Collateral15 struct {
	CollValDt         *ISODate                        `json:"collValDt,omitempty" xml:"CollValDt,omitempty"`
	AsstTp            []CollateralType7Choice         `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	NetXpsrCollstnInd *bool                           `json:"netXpsrCollstnInd,omitempty" xml:"NetXpsrCollstnInd,omitempty"`
	BsktIdr           *SecurityIdentification26Choice `json:"bsktIdr,omitempty" xml:"BsktIdr,omitempty"`
}

func (t *Collateral15) GetCollValDt() (out ISODate) {
	if t == nil || t.CollValDt == nil {
		return
	}
	return *t.CollValDt
}
func (t *Collateral15) GetAsstTp() (out []CollateralType7Choice) {
	return t.AsstTp
}
func (t *Collateral15) GetNetXpsrCollstnInd() (out bool) {
	if t == nil || t.NetXpsrCollstnInd == nil {
		return
	}
	return *t.NetXpsrCollstnInd
}
func (t *Collateral15) GetBsktIdr() (out SecurityIdentification26Choice) {
	if t == nil || t.BsktIdr == nil {
		return
	}
	return *t.BsktIdr
}

type Collateral27 struct {
	CollValDt *ISODate `json:"collValDt,omitempty" xml:"CollValDt,omitempty"`
}

func (t *Collateral27) GetCollValDt() (out ISODate) {
	if t == nil || t.CollValDt == nil {
		return
	}
	return *t.CollValDt
}

type CollateralData3 struct {
	AsstTp  []CollateralType7Choice         `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	BsktIdr *SecurityIdentification26Choice `json:"bsktIdr,omitempty" xml:"BsktIdr,omitempty"`
}

func (t *CollateralData3) GetAsstTp() (out []CollateralType7Choice) {
	return t.AsstTp
}
func (t *CollateralData3) GetBsktIdr() (out SecurityIdentification26Choice) {
	if t == nil || t.BsktIdr == nil {
		return
	}
	return *t.BsktIdr
}

type CollateralData7 struct {
	TxCollData TransactionCollateralData1Choice `json:"txCollData,omitempty" xml:"TxCollData"`
}

func (t *CollateralData7) GetTxCollData() (out TransactionCollateralData1Choice) {
	return t.TxCollData
}

// May be one of SICA, SIUR, TTCA
type CollateralDeliveryMethod1Code string

func GetCollateralDeliveryMethod1CodePointer(in string) *CollateralDeliveryMethod1Code {
	if in == "" {
		return nil
	}
	out := CollateralDeliveryMethod1Code(in)
	return &out
}

type CollateralFlag6 struct {
	HrcutOrMrgn string `json:"hrcutOrMrgn,omitempty" xml:"HrcutOrMrgn"`
}

func (t *CollateralFlag6) GetHrcutOrMrgn() (out string) {
	return t.HrcutOrMrgn
}

type CollateralFlag6Choice struct {
	Collsd   *CollaterisedData4 `json:"collsd,omitempty" xml:"Collsd,omitempty"`
	Uncollsd *NoReasonCode      `json:"uncollsd,omitempty" xml:"Uncollsd,omitempty"`
}

func (t *CollateralFlag6Choice) GetCollsd() (out CollaterisedData4) {
	if t == nil || t.Collsd == nil {
		return
	}
	return *t.Collsd
}
func (t *CollateralFlag6Choice) GetUncollsd() (out NoReasonCode) {
	if t == nil || t.Uncollsd == nil {
		return
	}
	return *t.Uncollsd
}

// May be one of INVG, NIVG, NOTR, NOAP
type CollateralQualityType1Code string

func GetCollateralQualityType1CodePointer(in string) *CollateralQualityType1Code {
	if in == "" {
		return nil
	}
	out := CollateralQualityType1Code(in)
	return &out
}

// May be one of GIVE, TAKE
type CollateralRole1Code string

func GetCollateralRole1CodePointer(in string) *CollateralRole1Code {
	if in == "" {
		return nil
	}
	out := CollateralRole1Code(in)
	return &out
}

type CollateralType7Choice struct {
	Scty   []HaircutPortfolioSecurityIdentification1 `json:"scty,omitempty" xml:"Scty,omitempty"`
	Csh    []ActiveOrHistoricCurrencyAndAmount       `json:"csh,omitempty" xml:"Csh,omitempty"`
	Cmmdty []Commodity3                              `json:"cmmdty,omitempty" xml:"Cmmdty,omitempty"`
}

func (t *CollateralType7Choice) GetScty() (out []HaircutPortfolioSecurityIdentification1) {
	return t.Scty
}
func (t *CollateralType7Choice) GetCsh() (out []ActiveOrHistoricCurrencyAndAmount) {
	return t.Csh
}
func (t *CollateralType7Choice) GetCmmdty() (out []Commodity3) {
	return t.Cmmdty
}

type CollaterisedData4 struct {
	CollValDt         *ISODate                        `json:"collValDt,omitempty" xml:"CollValDt,omitempty"`
	AsstTp            *CollateralType7Choice          `json:"asstTp,omitempty" xml:"AsstTp,omitempty"`
	NetXpsrCollstnInd *bool                           `json:"netXpsrCollstnInd,omitempty" xml:"NetXpsrCollstnInd,omitempty"`
	BsktIdr           *SecurityIdentification26Choice `json:"bsktIdr,omitempty" xml:"BsktIdr,omitempty"`
}

func (t *CollaterisedData4) GetCollValDt() (out ISODate) {
	if t == nil || t.CollValDt == nil {
		return
	}
	return *t.CollValDt
}
func (t *CollaterisedData4) GetAsstTp() (out CollateralType7Choice) {
	if t == nil || t.AsstTp == nil {
		return
	}
	return *t.AsstTp
}
func (t *CollaterisedData4) GetNetXpsrCollstnInd() (out bool) {
	if t == nil || t.NetXpsrCollstnInd == nil {
		return
	}
	return *t.NetXpsrCollstnInd
}
func (t *CollaterisedData4) GetBsktIdr() (out SecurityIdentification26Choice) {
	if t == nil || t.BsktIdr == nil {
		return
	}
	return *t.BsktIdr
}

type Commodity3 struct {
	Clssfctn *AssetClassCommodity5Choice        `json:"clssfctn,omitempty" xml:"Clssfctn,omitempty"`
	Qty      *Quantity13                        `json:"qty,omitempty" xml:"Qty,omitempty"`
	UnitPric *SecuritiesTransactionPrice2Choice `json:"unitPric,omitempty" xml:"UnitPric,omitempty"`
	MktVal   *string                            `json:"mktVal,omitempty" xml:"MktVal,omitempty"`
}

func (t *Commodity3) GetClssfctn() (out AssetClassCommodity5Choice) {
	if t == nil || t.Clssfctn == nil {
		return
	}
	return *t.Clssfctn
}
func (t *Commodity3) GetQty() (out Quantity13) {
	if t == nil || t.Qty == nil {
		return
	}
	return *t.Qty
}
func (t *Commodity3) GetUnitPric() (out SecuritiesTransactionPrice2Choice) {
	if t == nil || t.UnitPric == nil {
		return
	}
	return *t.UnitPric
}
func (t *Commodity3) GetMktVal() (out string) {
	if t == nil || t.MktVal == nil {
		return
	}
	return *t.MktVal
}

type ContractTerm2Choice struct {
	Opn *NoReasonCode       `json:"opn,omitempty" xml:"Opn,omitempty"`
	Fxd *FixedTermContract2 `json:"fxd,omitempty" xml:"Fxd,omitempty"`
}

func (t *ContractTerm2Choice) GetOpn() (out NoReasonCode) {
	if t == nil || t.Opn == nil {
		return
	}
	return *t.Opn
}
func (t *ContractTerm2Choice) GetFxd() (out FixedTermContract2) {
	if t == nil || t.Fxd == nil {
		return
	}
	return *t.Fxd
}

type CounterpartyData48 struct {
	RptgDtTm       ISODateTime                       `json:"rptgDtTm,omitempty" xml:"RptgDtTm"`
	RptSubmitgNtty OrganisationIdentification9Choice `json:"rptSubmitgNtty,omitempty" xml:"RptSubmitgNtty"`
	CtrPtyData     []CounterpartyData51              `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
}

func (t *CounterpartyData48) GetRptgDtTm() (out ISODateTime) {
	return t.RptgDtTm
}
func (t *CounterpartyData48) GetRptSubmitgNtty() (out OrganisationIdentification9Choice) {
	return t.RptSubmitgNtty
}
func (t *CounterpartyData48) GetCtrPtyData() (out []CounterpartyData51) {
	return t.CtrPtyData
}

type CounterpartyData49 struct {
	RptgDtTm       ISODateTime                       `json:"rptgDtTm,omitempty" xml:"RptgDtTm"`
	RptSubmitgNtty OrganisationIdentification9Choice `json:"rptSubmitgNtty,omitempty" xml:"RptSubmitgNtty"`
	CtrPtyData     []CounterpartyData50              `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
}

func (t *CounterpartyData49) GetRptgDtTm() (out ISODateTime) {
	return t.RptgDtTm
}
func (t *CounterpartyData49) GetRptSubmitgNtty() (out OrganisationIdentification9Choice) {
	return t.RptSubmitgNtty
}
func (t *CounterpartyData49) GetCtrPtyData() (out []CounterpartyData50) {
	return t.CtrPtyData
}

type CounterpartyData50 struct {
	RptgCtrPty CounterpartyIdentification1 `json:"rptgCtrPty,omitempty" xml:"RptgCtrPty"`
	OthrCtrPty CounterpartyIdentification2 `json:"othrCtrPty,omitempty" xml:"OthrCtrPty"`
}

func (t *CounterpartyData50) GetRptgCtrPty() (out CounterpartyIdentification1) {
	return t.RptgCtrPty
}
func (t *CounterpartyData50) GetOthrCtrPty() (out CounterpartyIdentification2) {
	return t.OthrCtrPty
}

type CounterpartyData51 struct {
	RptgCtrPty        CounterpartyIdentification1         `json:"rptgCtrPty,omitempty" xml:"RptgCtrPty"`
	OthrCtrPty        CounterpartyIdentification2         `json:"othrCtrPty,omitempty" xml:"OthrCtrPty"`
	NttyRspnsblForRpt *OrganisationIdentification9Choice  `json:"nttyRspnsblForRpt,omitempty" xml:"NttyRspnsblForRpt,omitempty"`
	TxSpcfcData       *TransactionCounterpartyData3Choice `json:"txSpcfcData,omitempty" xml:"TxSpcfcData,omitempty"`
}

func (t *CounterpartyData51) GetRptgCtrPty() (out CounterpartyIdentification1) {
	return t.RptgCtrPty
}
func (t *CounterpartyData51) GetOthrCtrPty() (out CounterpartyIdentification2) {
	return t.OthrCtrPty
}
func (t *CounterpartyData51) GetNttyRspnsblForRpt() (out OrganisationIdentification9Choice) {
	if t == nil || t.NttyRspnsblForRpt == nil {
		return
	}
	return *t.NttyRspnsblForRpt
}
func (t *CounterpartyData51) GetTxSpcfcData() (out TransactionCounterpartyData3Choice) {
	if t == nil || t.TxSpcfcData == nil {
		return
	}
	return *t.TxSpcfcData
}

type CounterpartyIdentification1 struct {
	Id    OrganisationIdentification9Choice `json:"id,omitempty" xml:"Id"`
	Ntr   *CounterpartyTradeNature3Choice   `json:"ntr,omitempty" xml:"Ntr,omitempty"`
	Brnch *Branch2Choice                    `json:"brnch,omitempty" xml:"Brnch,omitempty"`
	Sd    *CollateralRole1Code              `json:"sd,omitempty" xml:"Sd,omitempty"`
}

func (t *CounterpartyIdentification1) GetId() (out OrganisationIdentification9Choice) {
	return t.Id
}
func (t *CounterpartyIdentification1) GetNtr() (out CounterpartyTradeNature3Choice) {
	if t == nil || t.Ntr == nil {
		return
	}
	return *t.Ntr
}
func (t *CounterpartyIdentification1) GetBrnch() (out Branch2Choice) {
	if t == nil || t.Brnch == nil {
		return
	}
	return *t.Brnch
}
func (t *CounterpartyIdentification1) GetSd() (out CollateralRole1Code) {
	if t == nil || t.Sd == nil {
		return
	}
	return *t.Sd
}

type CounterpartyIdentification2 struct {
	Id     OrganisationIdentification9Choice `json:"id,omitempty" xml:"Id"`
	Brnch  *Branch2Choice                    `json:"brnch,omitempty" xml:"Brnch,omitempty"`
	CtryCd *CountryCode                      `json:"ctryCd,omitempty" xml:"CtryCd,omitempty"`
}

func (t *CounterpartyIdentification2) GetId() (out OrganisationIdentification9Choice) {
	return t.Id
}
func (t *CounterpartyIdentification2) GetBrnch() (out Branch2Choice) {
	if t == nil || t.Brnch == nil {
		return
	}
	return *t.Brnch
}
func (t *CounterpartyIdentification2) GetCtryCd() (out CountryCode) {
	if t == nil || t.CtryCd == nil {
		return
	}
	return *t.CtryCd
}

type CounterpartyTradeNature3Choice struct {
	FI  *FinancialPartyClassification1Choice `json:"fi,omitempty" xml:"FI,omitempty"`
	NFI []NACEDomainIdentifier               `json:"nfi,omitempty" xml:"NFI,omitempty"`
}

func (t *CounterpartyTradeNature3Choice) GetFI() (out FinancialPartyClassification1Choice) {
	if t == nil || t.FI == nil {
		return
	}
	return *t.FI
}
func (t *CounterpartyTradeNature3Choice) GetNFI() (out []NACEDomainIdentifier) {
	return t.NFI
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

func GetCountryCodePointer(in string) *CountryCode {
	if in == "" {
		return nil
	}
	out := CountryCode(in)
	return &out
}

type Document struct {
	SctiesFincgRptgTxRpt SecuritiesFinancingReportingTransactionReportV01 `json:"sctiesFincgRptgTxRpt,omitempty" xml:"SctiesFincgRptgTxRpt"`
}

func (t *Document) GetSctiesFincgRptgTxRpt() (out SecuritiesFinancingReportingTransactionReportV01) {
	return t.SctiesFincgRptgTxRpt
}

type EnergyCommodityCoal1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType24Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnergyCommodityCoal1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityCoal1) GetSubPdct() (out AssetClassSubProductType24Code) {
	return t.SubPdct
}

type EnergyCommodityDistillates1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType25Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnergyCommodityDistillates1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityDistillates1) GetSubPdct() (out AssetClassSubProductType25Code) {
	return t.SubPdct
}

type EnergyCommodityElectricity1 struct {
	BasePdct     AssetClassProductType2Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType6Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType5Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *EnergyCommodityElectricity1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityElectricity1) GetSubPdct() (out AssetClassSubProductType6Code) {
	return t.SubPdct
}
func (t *EnergyCommodityElectricity1) GetAddtlSubPdct() (out AssetClassDetailedSubProductType5Code) {
	return t.AddtlSubPdct
}

type EnergyCommodityInterEnergy1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType26Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnergyCommodityInterEnergy1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityInterEnergy1) GetSubPdct() (out AssetClassSubProductType26Code) {
	return t.SubPdct
}

type EnergyCommodityLightEnd1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType27Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnergyCommodityLightEnd1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityLightEnd1) GetSubPdct() (out AssetClassSubProductType27Code) {
	return t.SubPdct
}

type EnergyCommodityNaturalGas2 struct {
	BasePdct     AssetClassProductType2Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType7Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType31Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *EnergyCommodityNaturalGas2) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityNaturalGas2) GetSubPdct() (out AssetClassSubProductType7Code) {
	return t.SubPdct
}
func (t *EnergyCommodityNaturalGas2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType31Code) {
	return t.AddtlSubPdct
}

type EnergyCommodityOil2 struct {
	BasePdct     AssetClassProductType2Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType8Code          `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType32Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *EnergyCommodityOil2) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityOil2) GetSubPdct() (out AssetClassSubProductType8Code) {
	return t.SubPdct
}
func (t *EnergyCommodityOil2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType32Code) {
	return t.AddtlSubPdct
}

type EnergyCommodityOther1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnergyCommodityOther1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityOther1) GetSubPdct() (out AssetClassSubProductType49Code) {
	return t.SubPdct
}

type EnergyCommodityRenewableEnergy1 struct {
	BasePdct AssetClassProductType2Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType28Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnergyCommodityRenewableEnergy1) GetBasePdct() (out AssetClassProductType2Code) {
	return t.BasePdct
}
func (t *EnergyCommodityRenewableEnergy1) GetSubPdct() (out AssetClassSubProductType28Code) {
	return t.SubPdct
}

type EnvironmentCommodityOther1 struct {
	BasePdct AssetClassProductType3Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnvironmentCommodityOther1) GetBasePdct() (out AssetClassProductType3Code) {
	return t.BasePdct
}
func (t *EnvironmentCommodityOther1) GetSubPdct() (out AssetClassSubProductType49Code) {
	return t.SubPdct
}

type EnvironmentalCommodityCarbonRelated1 struct {
	BasePdct AssetClassProductType3Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType29Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnvironmentalCommodityCarbonRelated1) GetBasePdct() (out AssetClassProductType3Code) {
	return t.BasePdct
}
func (t *EnvironmentalCommodityCarbonRelated1) GetSubPdct() (out AssetClassSubProductType29Code) {
	return t.SubPdct
}

type EnvironmentalCommodityEmission2 struct {
	BasePdct     AssetClassProductType3Code            `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType10Code        `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType8Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *EnvironmentalCommodityEmission2) GetBasePdct() (out AssetClassProductType3Code) {
	return t.BasePdct
}
func (t *EnvironmentalCommodityEmission2) GetSubPdct() (out AssetClassSubProductType10Code) {
	return t.SubPdct
}
func (t *EnvironmentalCommodityEmission2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType8Code) {
	return t.AddtlSubPdct
}

type EnvironmentalCommodityWeather1 struct {
	BasePdct AssetClassProductType3Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType30Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *EnvironmentalCommodityWeather1) GetBasePdct() (out AssetClassProductType3Code) {
	return t.BasePdct
}
func (t *EnvironmentalCommodityWeather1) GetSubPdct() (out AssetClassSubProductType30Code) {
	return t.SubPdct
}

// May be one of SBSC, MGLD, REPO, SLEB
type ExposureType6Code string

func GetExposureType6CodePointer(in string) *ExposureType6Code {
	if in == "" {
		return nil
	}
	out := ExposureType6Code(in)
	return &out
}

// Must be at least 1 items long
type ExternalAgreementType1Code string

func GetExternalAgreementType1CodePointer(in string) *ExternalAgreementType1Code {
	if in == "" {
		return nil
	}
	out := ExternalAgreementType1Code(in)
	return &out
}

// Must be at least 1 items long
type ExternalSecuritiesLendingType1Code string

func GetExternalSecuritiesLendingType1CodePointer(in string) *ExternalSecuritiesLendingType1Code {
	if in == "" {
		return nil
	}
	out := ExternalSecuritiesLendingType1Code(in)
	return &out
}

type FertilizerCommodityAmmonia1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType39Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommodityAmmonia1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommodityAmmonia1) GetSubPdct() (out AssetClassSubProductType39Code) {
	return t.SubPdct
}

type FertilizerCommodityDiammoniumPhosphate1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType40Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommodityDiammoniumPhosphate1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommodityDiammoniumPhosphate1) GetSubPdct() (out AssetClassSubProductType40Code) {
	return t.SubPdct
}

type FertilizerCommodityOther1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommodityOther1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommodityOther1) GetSubPdct() (out AssetClassSubProductType49Code) {
	return t.SubPdct
}

type FertilizerCommodityPotash1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType41Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommodityPotash1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommodityPotash1) GetSubPdct() (out AssetClassSubProductType41Code) {
	return t.SubPdct
}

type FertilizerCommoditySulphur1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType42Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommoditySulphur1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommoditySulphur1) GetSubPdct() (out AssetClassSubProductType42Code) {
	return t.SubPdct
}

type FertilizerCommodityUrea1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType43Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommodityUrea1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommodityUrea1) GetSubPdct() (out AssetClassSubProductType43Code) {
	return t.SubPdct
}

type FertilizerCommodityUreaAndAmmoniumNitrate1 struct {
	BasePdct AssetClassProductType5Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType44Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FertilizerCommodityUreaAndAmmoniumNitrate1) GetBasePdct() (out AssetClassProductType5Code) {
	return t.BasePdct
}
func (t *FertilizerCommodityUreaAndAmmoniumNitrate1) GetSubPdct() (out AssetClassSubProductType44Code) {
	return t.SubPdct
}

type FinancialPartyClassification1Choice struct {
	Clssfctn           []FinancialPartySectorType2Code `json:"clssfctn,omitempty" xml:"Clssfctn,omitempty"`
	InvstmtFndClssfctn *FundType2Code                  `json:"invstmtFndClssfctn,omitempty" xml:"InvstmtFndClssfctn,omitempty"`
}

func (t *FinancialPartyClassification1Choice) GetClssfctn() (out []FinancialPartySectorType2Code) {
	return t.Clssfctn
}
func (t *FinancialPartyClassification1Choice) GetInvstmtFndClssfctn() (out FundType2Code) {
	if t == nil || t.InvstmtFndClssfctn == nil {
		return
	}
	return *t.InvstmtFndClssfctn
}

// May be one of AIFD, CSDS, CCPS, CDTI, INUN, ORPI, INVF, REIN, UCIT
type FinancialPartySectorType2Code string

func GetFinancialPartySectorType2CodePointer(in string) *FinancialPartySectorType2Code {
	if in == "" {
		return nil
	}
	out := FinancialPartySectorType2Code(in)
	return &out
}

type FixedRate2 struct {
	Rate       string                                 `json:"rate,omitempty" xml:"Rate"`
	DayCntBsis InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
}

func (t *FixedRate2) GetRate() (out string) {
	return t.Rate
}
func (t *FixedRate2) GetDayCntBsis() (out InterestComputationMethodFormat6Choice) {
	return t.DayCntBsis
}

type FixedRate7 struct {
	Rate       *string                                `json:"rate,omitempty" xml:"Rate,omitempty"`
	DayCntBsis InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
	MrgnLnAmt  *ActiveOrHistoricCurrencyAndAmount     `json:"mrgnLnAmt,omitempty" xml:"MrgnLnAmt,omitempty"`
}

func (t *FixedRate7) GetRate() (out string) {
	if t == nil || t.Rate == nil {
		return
	}
	return *t.Rate
}
func (t *FixedRate7) GetDayCntBsis() (out InterestComputationMethodFormat6Choice) {
	return t.DayCntBsis
}
func (t *FixedRate7) GetMrgnLnAmt() (out ActiveOrHistoricCurrencyAndAmount) {
	if t == nil || t.MrgnLnAmt == nil {
		return
	}
	return *t.MrgnLnAmt
}

type FixedTermContract2 struct {
	MtrtyDt     *ISODate                    `json:"mtrtyDt,omitempty" xml:"MtrtyDt,omitempty"`
	TermntnOptn *RepoTerminationOption2Code `json:"termntnOptn,omitempty" xml:"TermntnOptn,omitempty"`
}

func (t *FixedTermContract2) GetMtrtyDt() (out ISODate) {
	if t == nil || t.MtrtyDt == nil {
		return
	}
	return *t.MtrtyDt
}
func (t *FixedTermContract2) GetTermntnOptn() (out RepoTerminationOption2Code) {
	if t == nil || t.TermntnOptn == nil {
		return
	}
	return *t.TermntnOptn
}

type FloatingInterestRate10 struct {
	RefRate    BenchmarkCurveName8Choice `json:"refRate,omitempty" xml:"RefRate"`
	Term       InterestRateContractTerm2 `json:"term,omitempty" xml:"Term"`
	PmtFrqcy   InterestRateContractTerm2 `json:"pmtFrqcy,omitempty" xml:"PmtFrqcy"`
	RstFrqcy   InterestRateContractTerm2 `json:"rstFrqcy,omitempty" xml:"RstFrqcy"`
	BsisPtSprd string                    `json:"bsisPtSprd,omitempty" xml:"BsisPtSprd"`
}

func (t *FloatingInterestRate10) GetRefRate() (out BenchmarkCurveName8Choice) {
	return t.RefRate
}
func (t *FloatingInterestRate10) GetTerm() (out InterestRateContractTerm2) {
	return t.Term
}
func (t *FloatingInterestRate10) GetPmtFrqcy() (out InterestRateContractTerm2) {
	return t.PmtFrqcy
}
func (t *FloatingInterestRate10) GetRstFrqcy() (out InterestRateContractTerm2) {
	return t.RstFrqcy
}
func (t *FloatingInterestRate10) GetBsisPtSprd() (out string) {
	return t.BsisPtSprd
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

func (t *FloatingInterestRate12) GetRefRate() (out BenchmarkCurveName8Choice) {
	return t.RefRate
}
func (t *FloatingInterestRate12) GetTerm() (out InterestRateContractTerm2) {
	return t.Term
}
func (t *FloatingInterestRate12) GetPmtFrqcy() (out InterestRateContractTerm2) {
	return t.PmtFrqcy
}
func (t *FloatingInterestRate12) GetRstFrqcy() (out InterestRateContractTerm2) {
	return t.RstFrqcy
}
func (t *FloatingInterestRate12) GetBsisPtSprd() (out string) {
	return t.BsisPtSprd
}
func (t *FloatingInterestRate12) GetRateAdjstmnt() (out []RateAdjustment1) {
	return t.RateAdjstmnt
}
func (t *FloatingInterestRate12) GetDayCntBsis() (out InterestComputationMethodFormat6Choice) {
	return t.DayCntBsis
}

type FloatingInterestRate16 struct {
	RefRate    *BenchmarkCurveName8Choice             `json:"refRate,omitempty" xml:"RefRate,omitempty"`
	Term       *InterestRateContractTerm2             `json:"term,omitempty" xml:"Term,omitempty"`
	PmtFrqcy   *InterestRateContractTerm2             `json:"pmtFrqcy,omitempty" xml:"PmtFrqcy,omitempty"`
	RstFrqcy   *InterestRateContractTerm2             `json:"rstFrqcy,omitempty" xml:"RstFrqcy,omitempty"`
	BsisPtSprd *string                                `json:"bsisPtSprd,omitempty" xml:"BsisPtSprd,omitempty"`
	DayCntBsis InterestComputationMethodFormat6Choice `json:"dayCntBsis,omitempty" xml:"DayCntBsis"`
}

func (t *FloatingInterestRate16) GetRefRate() (out BenchmarkCurveName8Choice) {
	if t == nil || t.RefRate == nil {
		return
	}
	return *t.RefRate
}
func (t *FloatingInterestRate16) GetTerm() (out InterestRateContractTerm2) {
	if t == nil || t.Term == nil {
		return
	}
	return *t.Term
}
func (t *FloatingInterestRate16) GetPmtFrqcy() (out InterestRateContractTerm2) {
	if t == nil || t.PmtFrqcy == nil {
		return
	}
	return *t.PmtFrqcy
}
func (t *FloatingInterestRate16) GetRstFrqcy() (out InterestRateContractTerm2) {
	if t == nil || t.RstFrqcy == nil {
		return
	}
	return *t.RstFrqcy
}
func (t *FloatingInterestRate16) GetBsisPtSprd() (out string) {
	if t == nil || t.BsisPtSprd == nil {
		return
	}
	return *t.BsisPtSprd
}
func (t *FloatingInterestRate16) GetDayCntBsis() (out InterestComputationMethodFormat6Choice) {
	return t.DayCntBsis
}

type FreightCommodityContainerShip1 struct {
	BasePdct AssetClassProductType4Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType46Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FreightCommodityContainerShip1) GetBasePdct() (out AssetClassProductType4Code) {
	return t.BasePdct
}
func (t *FreightCommodityContainerShip1) GetSubPdct() (out AssetClassSubProductType46Code) {
	return t.SubPdct
}

type FreightCommodityDry2 struct {
	BasePdct     AssetClassProductType4Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType31Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType33Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *FreightCommodityDry2) GetBasePdct() (out AssetClassProductType4Code) {
	return t.BasePdct
}
func (t *FreightCommodityDry2) GetSubPdct() (out AssetClassSubProductType31Code) {
	return t.SubPdct
}
func (t *FreightCommodityDry2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType33Code) {
	return t.AddtlSubPdct
}

type FreightCommodityOther1 struct {
	BasePdct AssetClassProductType4Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *FreightCommodityOther1) GetBasePdct() (out AssetClassProductType4Code) {
	return t.BasePdct
}
func (t *FreightCommodityOther1) GetSubPdct() (out AssetClassSubProductType49Code) {
	return t.SubPdct
}

type FreightCommodityWet2 struct {
	BasePdct     AssetClassProductType4Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType32Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType34Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *FreightCommodityWet2) GetBasePdct() (out AssetClassProductType4Code) {
	return t.BasePdct
}
func (t *FreightCommodityWet2) GetSubPdct() (out AssetClassSubProductType32Code) {
	return t.SubPdct
}
func (t *FreightCommodityWet2) GetAddtlSubPdct() (out AssetClassDetailedSubProductType34Code) {
	return t.AddtlSubPdct
}

// May be one of ETFT, MMFT, OTHR, REIT
type FundType2Code string

func GetFundType2CodePointer(in string) *FundType2Code {
	if in == "" {
		return nil
	}
	out := FundType2Code(in)
	return &out
}

type HaircutPortfolioSecurityIdentification1 struct {
	PrtflHrcutOrMrgn *string   `json:"prtflHrcutOrMrgn,omitempty" xml:"PrtflHrcutOrMrgn,omitempty"`
	Id               Security3 `json:"id,omitempty" xml:"Id"`
}

func (t *HaircutPortfolioSecurityIdentification1) GetPrtflHrcutOrMrgn() (out string) {
	if t == nil || t.PrtflHrcutOrMrgn == nil {
		return
	}
	return *t.PrtflHrcutOrMrgn
}
func (t *HaircutPortfolioSecurityIdentification1) GetId() (out Security3) {
	return t.Id
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

func GetISINOct2015IdentifierPointer(in string) *ISINOct2015Identifier {
	if in == "" {
		return nil
	}
	out := ISINOct2015Identifier(in)
	return &out
}

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

func (t *IndustrialProductCommodityConstruction1) GetBasePdct() (out AssetClassProductType6Code) {
	return t.BasePdct
}
func (t *IndustrialProductCommodityConstruction1) GetSubPdct() (out AssetClassSubProductType33Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type IndustrialProductCommodityManufacturing1 struct {
	BasePdct AssetClassProductType6Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType34Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *IndustrialProductCommodityManufacturing1) GetBasePdct() (out AssetClassProductType6Code) {
	return t.BasePdct
}
func (t *IndustrialProductCommodityManufacturing1) GetSubPdct() (out AssetClassSubProductType34Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

// May be one of A001, A002, A003, A004, A005, A006, A007, A008, A009, A010, A011, A012, A013, A014
type InterestComputationMethod1Code string

func GetInterestComputationMethod1CodePointer(in string) *InterestComputationMethod1Code {
	if in == "" {
		return nil
	}
	out := InterestComputationMethod1Code(in)
	return &out
}

type InterestComputationMethodFormat6Choice struct {
	Cd    *InterestComputationMethod1Code `json:"cd,omitempty" xml:"Cd,omitempty"`
	Prtry *Max35Text                      `json:"prtry,omitempty" xml:"Prtry,omitempty"`
}

func (t *InterestComputationMethodFormat6Choice) GetCd() (out InterestComputationMethod1Code) {
	if t == nil || t.Cd == nil {
		return
	}
	return *t.Cd
}
func (t *InterestComputationMethodFormat6Choice) GetPrtry() (out Max35Text) {
	if t == nil || t.Prtry == nil {
		return
	}
	return *t.Prtry
}

type InterestRate14Choice struct {
	Fxd  *FixedRate2             `json:"fxd,omitempty" xml:"Fxd,omitempty"`
	Fltg *FloatingInterestRate12 `json:"fltg,omitempty" xml:"Fltg,omitempty"`
}

func (t *InterestRate14Choice) GetFxd() (out FixedRate2) {
	if t == nil || t.Fxd == nil {
		return
	}
	return *t.Fxd
}
func (t *InterestRate14Choice) GetFltg() (out FloatingInterestRate12) {
	if t == nil || t.Fltg == nil {
		return
	}
	return *t.Fltg
}

type InterestRate15Choice struct {
	Amt        *ActiveOrHistoricCurrencyAndAmount `json:"amt,omitempty" xml:"Amt,omitempty"`
	IntrstRate *InterestRate7Choice               `json:"intrstRate,omitempty" xml:"IntrstRate,omitempty"`
}

func (t *InterestRate15Choice) GetAmt() (out ActiveOrHistoricCurrencyAndAmount) {
	if t == nil || t.Amt == nil {
		return
	}
	return *t.Amt
}
func (t *InterestRate15Choice) GetIntrstRate() (out InterestRate7Choice) {
	if t == nil || t.IntrstRate == nil {
		return
	}
	return *t.IntrstRate
}

type InterestRate7Choice struct {
	Fxd  *FixedRate7             `json:"fxd,omitempty" xml:"Fxd,omitempty"`
	Fltg *FloatingInterestRate16 `json:"fltg,omitempty" xml:"Fltg,omitempty"`
}

func (t *InterestRate7Choice) GetFxd() (out FixedRate7) {
	if t == nil || t.Fxd == nil {
		return
	}
	return *t.Fxd
}
func (t *InterestRate7Choice) GetFltg() (out FloatingInterestRate16) {
	if t == nil || t.Fltg == nil {
		return
	}
	return *t.Fltg
}

type InterestRateContractTerm2 struct {
	Unit RateBasis1Code `json:"unit,omitempty" xml:"Unit"`
	Val  string         `json:"val,omitempty" xml:"Val"`
}

func (t *InterestRateContractTerm2) GetUnit() (out RateBasis1Code) {
	return t.Unit
}
func (t *InterestRateContractTerm2) GetVal() (out string) {
	return t.Val
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

func GetLEIIdentifierPointer(in string) *LEIIdentifier {
	if in == "" {
		return nil
	}
	out := LEIIdentifier(in)
	return &out
}

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

func (t *LoanData1) GetClrSts() (out Cleared8Choice) {
	if t == nil || t.ClrSts == nil {
		return
	}
	return *t.ClrSts
}
func (t *LoanData1) GetTradgVn() (out MICIdentifier) {
	if t == nil || t.TradgVn == nil {
		return
	}
	return *t.TradgVn
}
func (t *LoanData1) GetMstrAgrmt() (out MasterAgreement1) {
	if t == nil || t.MstrAgrmt == nil {
		return
	}
	return *t.MstrAgrmt
}
func (t *LoanData1) GetValDt() (out ISODate) {
	if t == nil || t.ValDt == nil {
		return
	}
	return *t.ValDt
}
func (t *LoanData1) GetMinNtcePrd() (out string) {
	if t == nil || t.MinNtcePrd == nil {
		return
	}
	return *t.MinNtcePrd
}
func (t *LoanData1) GetEarlstCallBckDt() (out ISODate) {
	if t == nil || t.EarlstCallBckDt == nil {
		return
	}
	return *t.EarlstCallBckDt
}
func (t *LoanData1) GetGnlColl() (out SpecialCollateral1Code) {
	if t == nil || t.GnlColl == nil {
		return
	}
	return *t.GnlColl
}
func (t *LoanData1) GetDlvryByVal() (out bool) {
	if t == nil || t.DlvryByVal == nil {
		return
	}
	return *t.DlvryByVal
}
func (t *LoanData1) GetCollDlvryMtd() (out CollateralDeliveryMethod1Code) {
	if t == nil || t.CollDlvryMtd == nil {
		return
	}
	return *t.CollDlvryMtd
}
func (t *LoanData1) GetTerm() (out []ContractTerm2Choice) {
	return t.Term
}
func (t *LoanData1) GetIntrstRate() (out InterestRate14Choice) {
	if t == nil || t.IntrstRate == nil {
		return
	}
	return *t.IntrstRate
}
func (t *LoanData1) GetPrncplAmt() (out PrincipalAmount1) {
	if t == nil || t.PrncplAmt == nil {
		return
	}
	return *t.PrncplAmt
}

type LoanData11 struct {
	UnqTradIdr Max52Text `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt      ISODate   `json:"evtDt,omitempty" xml:"EvtDt"`
	TermntnDt  ISODate   `json:"termntnDt,omitempty" xml:"TermntnDt"`
}

func (t *LoanData11) GetUnqTradIdr() (out Max52Text) {
	return t.UnqTradIdr
}
func (t *LoanData11) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *LoanData11) GetTermntnDt() (out ISODate) {
	return t.TermntnDt
}

type LoanData13 struct {
	UnqTradIdr Max52Text `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
}

func (t *LoanData13) GetUnqTradIdr() (out Max52Text) {
	return t.UnqTradIdr
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

func (t *LoanData2) GetClrSts() (out Cleared2Choice) {
	if t == nil || t.ClrSts == nil {
		return
	}
	return *t.ClrSts
}
func (t *LoanData2) GetTradgVn() (out MICIdentifier) {
	if t == nil || t.TradgVn == nil {
		return
	}
	return *t.TradgVn
}
func (t *LoanData2) GetMstrAgrmt() (out MasterAgreement1) {
	if t == nil || t.MstrAgrmt == nil {
		return
	}
	return *t.MstrAgrmt
}
func (t *LoanData2) GetValDt() (out ISODate) {
	if t == nil || t.ValDt == nil {
		return
	}
	return *t.ValDt
}
func (t *LoanData2) GetMtrtyDt() (out ISODate) {
	if t == nil || t.MtrtyDt == nil {
		return
	}
	return *t.MtrtyDt
}
func (t *LoanData2) GetGnlColl() (out SpecialCollateral1Code) {
	if t == nil || t.GnlColl == nil {
		return
	}
	return *t.GnlColl
}
func (t *LoanData2) GetPrncplAmt() (out PrincipalAmount1) {
	if t == nil || t.PrncplAmt == nil {
		return
	}
	return *t.PrncplAmt
}
func (t *LoanData2) GetUnitPric() (out SecuritiesTransactionPrice2Choice) {
	if t == nil || t.UnitPric == nil {
		return
	}
	return *t.UnitPric
}

type LoanData27 struct {
	PrncplAmt PrincipalAmount1 `json:"prncplAmt,omitempty" xml:"PrncplAmt"`
	MtrtyDt   ISODate          `json:"mtrtyDt,omitempty" xml:"MtrtyDt"`
}

func (t *LoanData27) GetPrncplAmt() (out PrincipalAmount1) {
	return t.PrncplAmt
}
func (t *LoanData27) GetMtrtyDt() (out ISODate) {
	return t.MtrtyDt
}

type LoanData29 struct {
	UnqTradIdr *Max52Text       `json:"unqTradIdr,omitempty" xml:"UnqTradIdr,omitempty"`
	MstrAgrmt  MasterAgreement1 `json:"mstrAgrmt,omitempty" xml:"MstrAgrmt"`
}

func (t *LoanData29) GetUnqTradIdr() (out Max52Text) {
	if t == nil || t.UnqTradIdr == nil {
		return
	}
	return *t.UnqTradIdr
}
func (t *LoanData29) GetMstrAgrmt() (out MasterAgreement1) {
	return t.MstrAgrmt
}

type LoanData32 struct {
	CollDlvryMtd     *CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd,omitempty"`
	OutsdngMrgnLnAmt *string                        `json:"outsdngMrgnLnAmt,omitempty" xml:"OutsdngMrgnLnAmt,omitempty"`
	ShrtMktValAmt    *string                        `json:"shrtMktValAmt,omitempty" xml:"ShrtMktValAmt,omitempty"`
	Ccy              *ActiveOrHistoricCurrencyCode  `json:"ccy,omitempty" xml:"Ccy,omitempty"`
	MrgnLnAttr       []InterestRate15Choice         `json:"mrgnLnAttr,omitempty" xml:"MrgnLnAttr,omitempty"`
}

func (t *LoanData32) GetCollDlvryMtd() (out CollateralDeliveryMethod1Code) {
	if t == nil || t.CollDlvryMtd == nil {
		return
	}
	return *t.CollDlvryMtd
}
func (t *LoanData32) GetOutsdngMrgnLnAmt() (out string) {
	if t == nil || t.OutsdngMrgnLnAmt == nil {
		return
	}
	return *t.OutsdngMrgnLnAmt
}
func (t *LoanData32) GetShrtMktValAmt() (out string) {
	if t == nil || t.ShrtMktValAmt == nil {
		return
	}
	return *t.ShrtMktValAmt
}
func (t *LoanData32) GetCcy() (out ActiveOrHistoricCurrencyCode) {
	if t == nil || t.Ccy == nil {
		return
	}
	return *t.Ccy
}
func (t *LoanData32) GetMrgnLnAttr() (out []InterestRate15Choice) {
	return t.MrgnLnAttr
}

type LoanData37 struct {
	DlvryByVal   bool                          `json:"dlvryByVal,omitempty" xml:"DlvryByVal"`
	CollDlvryMtd CollateralDeliveryMethod1Code `json:"collDlvryMtd,omitempty" xml:"CollDlvryMtd"`
	Term         []ContractTerm2Choice         `json:"term,omitempty" xml:"Term,omitempty"`
	IntrstRate   InterestRate14Choice          `json:"intrstRate,omitempty" xml:"IntrstRate"`
	PrncplAmt    PrincipalAmount1              `json:"prncplAmt,omitempty" xml:"PrncplAmt"`
}

func (t *LoanData37) GetDlvryByVal() (out bool) {
	return t.DlvryByVal
}
func (t *LoanData37) GetCollDlvryMtd() (out CollateralDeliveryMethod1Code) {
	return t.CollDlvryMtd
}
func (t *LoanData37) GetTerm() (out []ContractTerm2Choice) {
	return t.Term
}
func (t *LoanData37) GetIntrstRate() (out InterestRate14Choice) {
	return t.IntrstRate
}
func (t *LoanData37) GetPrncplAmt() (out PrincipalAmount1) {
	return t.PrncplAmt
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

func (t *LoanData39) GetClrSts() (out Cleared2Choice) {
	if t == nil || t.ClrSts == nil {
		return
	}
	return *t.ClrSts
}
func (t *LoanData39) GetTradgVn() (out MICIdentifier) {
	if t == nil || t.TradgVn == nil {
		return
	}
	return *t.TradgVn
}
func (t *LoanData39) GetMstrAgrmt() (out MasterAgreement1) {
	if t == nil || t.MstrAgrmt == nil {
		return
	}
	return *t.MstrAgrmt
}
func (t *LoanData39) GetValDt() (out ISODate) {
	if t == nil || t.ValDt == nil {
		return
	}
	return *t.ValDt
}
func (t *LoanData39) GetGnlColl() (out SpecialCollateral1Code) {
	if t == nil || t.GnlColl == nil {
		return
	}
	return *t.GnlColl
}
func (t *LoanData39) GetDlvryByVal() (out bool) {
	if t == nil || t.DlvryByVal == nil {
		return
	}
	return *t.DlvryByVal
}
func (t *LoanData39) GetCollDlvryMtd() (out CollateralDeliveryMethod1Code) {
	if t == nil || t.CollDlvryMtd == nil {
		return
	}
	return *t.CollDlvryMtd
}
func (t *LoanData39) GetTerm() (out []ContractTerm2Choice) {
	return t.Term
}
func (t *LoanData39) GetAsstTp() (out SecurityCommodity2Choice) {
	if t == nil || t.AsstTp == nil {
		return
	}
	return *t.AsstTp
}
func (t *LoanData39) GetLnVal() (out string) {
	if t == nil || t.LnVal == nil {
		return
	}
	return *t.LnVal
}
func (t *LoanData39) GetRbtRate() (out RebateRate1Choice) {
	if t == nil || t.RbtRate == nil {
		return
	}
	return *t.RbtRate
}
func (t *LoanData39) GetLndgFee() (out string) {
	if t == nil || t.LndgFee == nil {
		return
	}
	return *t.LndgFee
}
func (t *LoanData39) GetExclsvArrgmnt() (out bool) {
	if t == nil || t.ExclsvArrgmnt == nil {
		return
	}
	return *t.ExclsvArrgmnt
}

type LoanData4 struct {
	EvtDt    ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	TxLnData TransactionLoanData4Choice `json:"txLnData,omitempty" xml:"TxLnData"`
}

func (t *LoanData4) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *LoanData4) GetTxLnData() (out TransactionLoanData4Choice) {
	return t.TxLnData
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

func (t *LoanData40) GetUnqTradIdr() (out Max52Text) {
	return t.UnqTradIdr
}
func (t *LoanData40) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *LoanData40) GetCtrctTp() (out ExposureType6Code) {
	return t.CtrctTp
}
func (t *LoanData40) GetClrSts() (out Cleared8Choice) {
	return t.ClrSts
}
func (t *LoanData40) GetTradgVn() (out MICIdentifier) {
	return t.TradgVn
}
func (t *LoanData40) GetMstrAgrmt() (out MasterAgreement1) {
	if t == nil || t.MstrAgrmt == nil {
		return
	}
	return *t.MstrAgrmt
}
func (t *LoanData40) GetExctnDtTm() (out ISODateTime) {
	return t.ExctnDtTm
}
func (t *LoanData40) GetValDt() (out ISODate) {
	return t.ValDt
}
func (t *LoanData40) GetTermntnDt() (out ISODate) {
	return t.TermntnDt
}
func (t *LoanData40) GetGnlColl() (out SpecialCollateral1Code) {
	if t == nil || t.GnlColl == nil {
		return
	}
	return *t.GnlColl
}
func (t *LoanData40) GetUnitPric() (out SecuritiesTransactionPrice2Choice) {
	if t == nil || t.UnitPric == nil {
		return
	}
	return *t.UnitPric
}
func (t *LoanData40) GetTxSpcfcData() (out TransactionLoanData6Choice) {
	return t.TxSpcfcData
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

func (t *LoanData41) GetDlvryByVal() (out bool) {
	return t.DlvryByVal
}
func (t *LoanData41) GetCollDlvryMtd() (out CollateralDeliveryMethod1Code) {
	if t == nil || t.CollDlvryMtd == nil {
		return
	}
	return *t.CollDlvryMtd
}
func (t *LoanData41) GetTerm() (out []ContractTerm2Choice) {
	return t.Term
}
func (t *LoanData41) GetAsstTp() (out SecurityCommodity2Choice) {
	return t.AsstTp
}
func (t *LoanData41) GetRbtRate() (out RebateRate1Choice) {
	return t.RbtRate
}
func (t *LoanData41) GetLnVal() (out string) {
	return t.LnVal
}
func (t *LoanData41) GetLndgFee() (out string) {
	return t.LndgFee
}
func (t *LoanData41) GetExclsvArrgmnt() (out bool) {
	return t.ExclsvArrgmnt
}

type LoanData42 struct {
	UnqTradIdr  Max52Text                  `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt       ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp     *ExposureType6Code         `json:"ctrctTp,omitempty" xml:"CtrctTp,omitempty"`
	TxSpcfcData TransactionLoanData5Choice `json:"txSpcfcData,omitempty" xml:"TxSpcfcData"`
}

func (t *LoanData42) GetUnqTradIdr() (out Max52Text) {
	return t.UnqTradIdr
}
func (t *LoanData42) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *LoanData42) GetCtrctTp() (out ExposureType6Code) {
	if t == nil || t.CtrctTp == nil {
		return
	}
	return *t.CtrctTp
}
func (t *LoanData42) GetTxSpcfcData() (out TransactionLoanData5Choice) {
	return t.TxSpcfcData
}

type LoanData43 struct {
	UnqTradIdr  Max52Text                  `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt       ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp     *ExposureType6Code         `json:"ctrctTp,omitempty" xml:"CtrctTp,omitempty"`
	ExctnDtTm   *ISODateTime               `json:"exctnDtTm,omitempty" xml:"ExctnDtTm,omitempty"`
	TermntnDt   *ISODate                   `json:"termntnDt,omitempty" xml:"TermntnDt,omitempty"`
	TxSpcfcData TransactionLoanData5Choice `json:"txSpcfcData,omitempty" xml:"TxSpcfcData"`
}

func (t *LoanData43) GetUnqTradIdr() (out Max52Text) {
	return t.UnqTradIdr
}
func (t *LoanData43) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *LoanData43) GetCtrctTp() (out ExposureType6Code) {
	if t == nil || t.CtrctTp == nil {
		return
	}
	return *t.CtrctTp
}
func (t *LoanData43) GetExctnDtTm() (out ISODateTime) {
	if t == nil || t.ExctnDtTm == nil {
		return
	}
	return *t.ExctnDtTm
}
func (t *LoanData43) GetTermntnDt() (out ISODate) {
	if t == nil || t.TermntnDt == nil {
		return
	}
	return *t.TermntnDt
}
func (t *LoanData43) GetTxSpcfcData() (out TransactionLoanData5Choice) {
	return t.TxSpcfcData
}

type LoanData44 struct {
	UnqTradIdr Max52Text                  `json:"unqTradIdr,omitempty" xml:"UnqTradIdr"`
	EvtDt      ISODate                    `json:"evtDt,omitempty" xml:"EvtDt"`
	CtrctTp    ExposureType6Code          `json:"ctrctTp,omitempty" xml:"CtrctTp"`
	ExctnDtTm  ISODateTime                `json:"exctnDtTm,omitempty" xml:"ExctnDtTm"`
	TxLnData   TransactionLoanData5Choice `json:"txLnData,omitempty" xml:"TxLnData"`
}

func (t *LoanData44) GetUnqTradIdr() (out Max52Text) {
	return t.UnqTradIdr
}
func (t *LoanData44) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *LoanData44) GetCtrctTp() (out ExposureType6Code) {
	return t.CtrctTp
}
func (t *LoanData44) GetExctnDtTm() (out ISODateTime) {
	return t.ExctnDtTm
}
func (t *LoanData44) GetTxLnData() (out TransactionLoanData5Choice) {
	return t.TxLnData
}

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

func GetMICIdentifierPointer(in string) *MICIdentifier {
	if in == "" {
		return nil
	}
	out := MICIdentifier(in)
	return &out
}

type MasterAgreement1 struct {
	Tp                AgreementType1Choice `json:"tp,omitempty" xml:"Tp"`
	Vrsn              *ISORestrictedYear   `json:"vrsn,omitempty" xml:"Vrsn,omitempty"`
	OthrMstrAgrmtDtls *Max50Text           `json:"othrMstrAgrmtDtls,omitempty" xml:"OthrMstrAgrmtDtls,omitempty"`
}

func (t *MasterAgreement1) GetTp() (out AgreementType1Choice) {
	return t.Tp
}
func (t *MasterAgreement1) GetVrsn() (out ISORestrictedYear) {
	if t == nil || t.Vrsn == nil {
		return
	}
	return *t.Vrsn
}
func (t *MasterAgreement1) GetOthrMstrAgrmtDtls() (out Max50Text) {
	if t == nil || t.OthrMstrAgrmtDtls == nil {
		return
	}
	return *t.OthrMstrAgrmtDtls
}

// Must be at least 1 items long
type Max350Text string

func GetMax350TextPointer(in string) *Max350Text {
	if in == "" {
		return nil
	}
	out := Max350Text(in)
	return &out
}

// Must be at least 1 items long
type Max35Text string

func GetMax35TextPointer(in string) *Max35Text {
	if in == "" {
		return nil
	}
	out := Max35Text(in)
	return &out
}

// Must be at least 1 items long
type Max50Text string

func GetMax50TextPointer(in string) *Max50Text {
	if in == "" {
		return nil
	}
	out := Max50Text(in)
	return &out
}

// Must be at least 1 items long
type Max52Text string

func GetMax52TextPointer(in string) *Max52Text {
	if in == "" {
		return nil
	}
	out := Max52Text(in)
	return &out
}

type MetalCommodityNonPrecious1 struct {
	BasePdct     AssetClassProductType7Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType15Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType10Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *MetalCommodityNonPrecious1) GetBasePdct() (out AssetClassProductType7Code) {
	return t.BasePdct
}
func (t *MetalCommodityNonPrecious1) GetSubPdct() (out AssetClassSubProductType15Code) {
	return t.SubPdct
}
func (t *MetalCommodityNonPrecious1) GetAddtlSubPdct() (out AssetClassDetailedSubProductType10Code) {
	return t.AddtlSubPdct
}

type MetalCommodityPrecious1 struct {
	BasePdct     AssetClassProductType7Code             `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct      AssetClassSubProductType16Code         `json:"subPdct,omitempty" xml:"SubPdct"`
	AddtlSubPdct AssetClassDetailedSubProductType11Code `json:"addtlSubPdct,omitempty" xml:"AddtlSubPdct"`
}

func (t *MetalCommodityPrecious1) GetBasePdct() (out AssetClassProductType7Code) {
	return t.BasePdct
}
func (t *MetalCommodityPrecious1) GetSubPdct() (out AssetClassSubProductType16Code) {
	return t.SubPdct
}
func (t *MetalCommodityPrecious1) GetAddtlSubPdct() (out AssetClassDetailedSubProductType11Code) {
	return t.AddtlSubPdct
}

// May be one of PSTN, TCTN
type ModificationLevel1Code string

func GetModificationLevel1CodePointer(in string) *ModificationLevel1Code {
	if in == "" {
		return nil
	}
	out := ModificationLevel1Code(in)
	return &out
}

// Must match the pattern [A-U]{1,1}
type NACEDomainIdentifier string

func GetNACEDomainIdentifierPointer(in string) *NACEDomainIdentifier {
	if in == "" {
		return nil
	}
	out := NACEDomainIdentifier(in)
	return &out
}

// May be one of NORE
type NoReasonCode string

func GetNoReasonCodePointer(in string) *NoReasonCode {
	if in == "" {
		return nil
	}
	out := NoReasonCode(in)
	return &out
}

// May be one of NTAV
type NotAvailable1Code string

func GetNotAvailable1CodePointer(in string) *NotAvailable1Code {
	if in == "" {
		return nil
	}
	out := NotAvailable1Code(in)
	return &out
}

type OrganisationIdentification9Choice struct {
	LEI    *LEIIdentifier           `json:"lei,omitempty" xml:"LEI,omitempty"`
	ClntId *Max50Text               `json:"clntId,omitempty" xml:"ClntId,omitempty"`
	AnyBIC *AnyBICDec2014Identifier `json:"anyBIC,omitempty" xml:"AnyBIC,omitempty"`
}

func (t *OrganisationIdentification9Choice) GetLEI() (out LEIIdentifier) {
	if t == nil || t.LEI == nil {
		return
	}
	return *t.LEI
}
func (t *OrganisationIdentification9Choice) GetClntId() (out Max50Text) {
	if t == nil || t.ClntId == nil {
		return
	}
	return *t.ClntId
}
func (t *OrganisationIdentification9Choice) GetAnyBIC() (out AnyBICDec2014Identifier) {
	if t == nil || t.AnyBIC == nil {
		return
	}
	return *t.AnyBIC
}

type OtherC10CommodityDeliverable2 struct {
	BasePdct AssetClassProductType11Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType47Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *OtherC10CommodityDeliverable2) GetBasePdct() (out AssetClassProductType11Code) {
	return t.BasePdct
}
func (t *OtherC10CommodityDeliverable2) GetSubPdct() (out AssetClassSubProductType47Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type OtherC10CommodityNonDeliverable2 struct {
	BasePdct AssetClassProductType11Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType48Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *OtherC10CommodityNonDeliverable2) GetBasePdct() (out AssetClassProductType11Code) {
	return t.BasePdct
}
func (t *OtherC10CommodityNonDeliverable2) GetSubPdct() (out AssetClassSubProductType48Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PaperCommodityContainerBoard1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType35Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *PaperCommodityContainerBoard1) GetBasePdct() (out AssetClassProductType8Code) {
	return t.BasePdct
}
func (t *PaperCommodityContainerBoard1) GetSubPdct() (out AssetClassSubProductType35Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PaperCommodityNewsprint1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType36Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *PaperCommodityNewsprint1) GetBasePdct() (out AssetClassProductType8Code) {
	return t.BasePdct
}
func (t *PaperCommodityNewsprint1) GetSubPdct() (out AssetClassSubProductType36Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PaperCommodityPulp1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType37Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *PaperCommodityPulp1) GetBasePdct() (out AssetClassProductType8Code) {
	return t.BasePdct
}
func (t *PaperCommodityPulp1) GetSubPdct() (out AssetClassSubProductType37Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PaperCommodityRecoveredPaper1 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType38Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *PaperCommodityRecoveredPaper1) GetBasePdct() (out AssetClassProductType8Code) {
	return t.BasePdct
}
func (t *PaperCommodityRecoveredPaper1) GetSubPdct() (out AssetClassSubProductType38Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PaperCommodityRecoveredPaper2 struct {
	BasePdct AssetClassProductType8Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *PaperCommodityRecoveredPaper2) GetBasePdct() (out AssetClassProductType8Code) {
	return t.BasePdct
}
func (t *PaperCommodityRecoveredPaper2) GetSubPdct() (out AssetClassSubProductType49Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PolypropyleneCommodityOther1 struct {
	BasePdct AssetClassProductType9Code     `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  AssetClassSubProductType49Code `json:"subPdct,omitempty" xml:"SubPdct"`
}

func (t *PolypropyleneCommodityOther1) GetBasePdct() (out AssetClassProductType9Code) {
	return t.BasePdct
}
func (t *PolypropyleneCommodityOther1) GetSubPdct() (out AssetClassSubProductType49Code) {
	return t.SubPdct
}

type PolypropyleneCommodityPlastic1 struct {
	BasePdct AssetClassProductType9Code      `json:"basePdct,omitempty" xml:"BasePdct"`
	SubPdct  *AssetClassSubProductType18Code `json:"subPdct,omitempty" xml:"SubPdct,omitempty"`
}

func (t *PolypropyleneCommodityPlastic1) GetBasePdct() (out AssetClassProductType9Code) {
	return t.BasePdct
}
func (t *PolypropyleneCommodityPlastic1) GetSubPdct() (out AssetClassSubProductType18Code) {
	if t == nil || t.SubPdct == nil {
		return
	}
	return *t.SubPdct
}

type PrincipalAmount1 struct {
	ValDtAmt   string                       `json:"valDtAmt,omitempty" xml:"ValDtAmt"`
	MtrtyDtAmt string                       `json:"mtrtyDtAmt,omitempty" xml:"MtrtyDtAmt"`
	Ccy        ActiveOrHistoricCurrencyCode `json:"ccy,omitempty" xml:"Ccy"`
}

func (t *PrincipalAmount1) GetValDtAmt() (out string) {
	return t.ValDtAmt
}
func (t *PrincipalAmount1) GetMtrtyDtAmt() (out string) {
	return t.MtrtyDtAmt
}
func (t *PrincipalAmount1) GetCcy() (out ActiveOrHistoricCurrencyCode) {
	return t.Ccy
}

type Quantity13 struct {
	Val         string             `json:"val,omitempty" xml:"Val"`
	UnitOfMeasr UnitOfMeasure1Code `json:"unitOfMeasr,omitempty" xml:"UnitOfMeasr"`
}

func (t *Quantity13) GetVal() (out string) {
	return t.Val
}
func (t *Quantity13) GetUnitOfMeasr() (out UnitOfMeasure1Code) {
	return t.UnitOfMeasr
}

type QuantityNominalValue1Choice struct {
	Qty     *string                            `json:"qty,omitempty" xml:"Qty,omitempty"`
	NmnlVal *ActiveOrHistoricCurrencyAndAmount `json:"nmnlVal,omitempty" xml:"NmnlVal,omitempty"`
}

func (t *QuantityNominalValue1Choice) GetQty() (out string) {
	if t == nil || t.Qty == nil {
		return
	}
	return *t.Qty
}
func (t *QuantityNominalValue1Choice) GetNmnlVal() (out ActiveOrHistoricCurrencyAndAmount) {
	if t == nil || t.NmnlVal == nil {
		return
	}
	return *t.NmnlVal
}

type RateAdjustment1 struct {
	Rate       string  `json:"rate,omitempty" xml:"Rate"`
	AdjstmntDt ISODate `json:"adjstmntDt,omitempty" xml:"AdjstmntDt"`
}

func (t *RateAdjustment1) GetRate() (out string) {
	return t.Rate
}
func (t *RateAdjustment1) GetAdjstmntDt() (out ISODate) {
	return t.AdjstmntDt
}

// May be one of DAYS, MNTH, WEEK, YEAR
type RateBasis1Code string

func GetRateBasis1CodePointer(in string) *RateBasis1Code {
	if in == "" {
		return nil
	}
	out := RateBasis1Code(in)
	return &out
}

type RebateRate1Choice struct {
	Fxd  *string                 `json:"fxd,omitempty" xml:"Fxd,omitempty"`
	Fltg *FloatingInterestRate10 `json:"fltg,omitempty" xml:"Fltg,omitempty"`
}

func (t *RebateRate1Choice) GetFxd() (out string) {
	if t == nil || t.Fxd == nil {
		return
	}
	return *t.Fxd
}
func (t *RebateRate1Choice) GetFltg() (out FloatingInterestRate10) {
	if t == nil || t.Fltg == nil {
		return
	}
	return *t.Fltg
}

// May be one of EGRN, EGAE, ETSB, NOAP
type RepoTerminationOption2Code string

func GetRepoTerminationOption2CodePointer(in string) *RepoTerminationOption2Code {
	if in == "" {
		return nil
	}
	out := RepoTerminationOption2Code(in)
	return &out
}

type SecuritiesFinancingReportingTransactionReportV01 struct {
	TradData    []TradeTransactionReport6Choice `json:"tradData,omitempty" xml:"TradData"`
	SplmtryData []SupplementaryData1            `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *SecuritiesFinancingReportingTransactionReportV01) GetTradData() (out []TradeTransactionReport6Choice) {
	return t.TradData
}
func (t *SecuritiesFinancingReportingTransactionReportV01) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type SecuritiesLendingType3Choice struct {
	Cd    *ExternalSecuritiesLendingType1Code `json:"cd,omitempty" xml:"Cd,omitempty"`
	Prtry *Max35Text                          `json:"prtry,omitempty" xml:"Prtry,omitempty"`
}

func (t *SecuritiesLendingType3Choice) GetCd() (out ExternalSecuritiesLendingType1Code) {
	if t == nil || t.Cd == nil {
		return
	}
	return *t.Cd
}
func (t *SecuritiesLendingType3Choice) GetPrtry() (out Max35Text) {
	if t == nil || t.Prtry == nil {
		return
	}
	return *t.Prtry
}

type SecuritiesTransactionPrice2Choice struct {
	MntryVal *AmountAndDirection61 `json:"mntryVal,omitempty" xml:"MntryVal,omitempty"`
	Pctg     *string               `json:"pctg,omitempty" xml:"Pctg,omitempty"`
	Yld      *string               `json:"yld,omitempty" xml:"Yld,omitempty"`
	BsisPts  *string               `json:"bsisPts,omitempty" xml:"BsisPts,omitempty"`
}

func (t *SecuritiesTransactionPrice2Choice) GetMntryVal() (out AmountAndDirection61) {
	if t == nil || t.MntryVal == nil {
		return
	}
	return *t.MntryVal
}
func (t *SecuritiesTransactionPrice2Choice) GetPctg() (out string) {
	if t == nil || t.Pctg == nil {
		return
	}
	return *t.Pctg
}
func (t *SecuritiesTransactionPrice2Choice) GetYld() (out string) {
	if t == nil || t.Yld == nil {
		return
	}
	return *t.Yld
}
func (t *SecuritiesTransactionPrice2Choice) GetBsisPts() (out string) {
	if t == nil || t.BsisPts == nil {
		return
	}
	return *t.BsisPts
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

func (t *Security3) GetId() (out ISINOct2015Identifier) {
	if t == nil || t.Id == nil {
		return
	}
	return *t.Id
}
func (t *Security3) GetClssfctnTp() (out CFIOct2015Identifier) {
	if t == nil || t.ClssfctnTp == nil {
		return
	}
	return *t.ClssfctnTp
}
func (t *Security3) GetQtyOrNmnlVal() (out QuantityNominalValue1Choice) {
	if t == nil || t.QtyOrNmnlVal == nil {
		return
	}
	return *t.QtyOrNmnlVal
}
func (t *Security3) GetUnitPric() (out SecuritiesTransactionPrice2Choice) {
	if t == nil || t.UnitPric == nil {
		return
	}
	return *t.UnitPric
}
func (t *Security3) GetMktVal() (out string) {
	if t == nil || t.MktVal == nil {
		return
	}
	return *t.MktVal
}
func (t *Security3) GetHrcutOrMrgn() (out string) {
	if t == nil || t.HrcutOrMrgn == nil {
		return
	}
	return *t.HrcutOrMrgn
}
func (t *Security3) GetQlty() (out CollateralQualityType1Code) {
	if t == nil || t.Qlty == nil {
		return
	}
	return *t.Qlty
}
func (t *Security3) GetMtrty() (out ISODate) {
	if t == nil || t.Mtrty == nil {
		return
	}
	return *t.Mtrty
}
func (t *Security3) GetIssr() (out SecurityIssuer1) {
	if t == nil || t.Issr == nil {
		return
	}
	return *t.Issr
}
func (t *Security3) GetTp() (out SecuritiesLendingType3Choice) {
	if t == nil || t.Tp == nil {
		return
	}
	return *t.Tp
}
func (t *Security3) GetAvlblForCollReuse() (out bool) {
	if t == nil || t.AvlblForCollReuse == nil {
		return
	}
	return *t.AvlblForCollReuse
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

func (t *Security4) GetId() (out ISINOct2015Identifier) {
	if t == nil || t.Id == nil {
		return
	}
	return *t.Id
}
func (t *Security4) GetClssfctn() (out CFIOct2015Identifier) {
	if t == nil || t.Clssfctn == nil {
		return
	}
	return *t.Clssfctn
}
func (t *Security4) GetQtyOrNmnlVal() (out QuantityNominalValue1Choice) {
	if t == nil || t.QtyOrNmnlVal == nil {
		return
	}
	return *t.QtyOrNmnlVal
}
func (t *Security4) GetQlty() (out CollateralQualityType1Code) {
	if t == nil || t.Qlty == nil {
		return
	}
	return *t.Qlty
}
func (t *Security4) GetMtrty() (out ISODate) {
	if t == nil || t.Mtrty == nil {
		return
	}
	return *t.Mtrty
}
func (t *Security4) GetIssr() (out SecurityIssuer1) {
	if t == nil || t.Issr == nil {
		return
	}
	return *t.Issr
}
func (t *Security4) GetTp() (out []SecuritiesLendingType3Choice) {
	return t.Tp
}
func (t *Security4) GetUnitPric() (out SecuritiesTransactionPrice2Choice) {
	if t == nil || t.UnitPric == nil {
		return
	}
	return *t.UnitPric
}
func (t *Security4) GetExclsvArrgmnt() (out bool) {
	if t == nil || t.ExclsvArrgmnt == nil {
		return
	}
	return *t.ExclsvArrgmnt
}
func (t *Security4) GetMktVal() (out string) {
	if t == nil || t.MktVal == nil {
		return
	}
	return *t.MktVal
}

type SecurityCommodity2Choice struct {
	Scty   *Security4  `json:"scty,omitempty" xml:"Scty,omitempty"`
	Cmmdty *Commodity3 `json:"cmmdty,omitempty" xml:"Cmmdty,omitempty"`
}

func (t *SecurityCommodity2Choice) GetScty() (out Security4) {
	if t == nil || t.Scty == nil {
		return
	}
	return *t.Scty
}
func (t *SecurityCommodity2Choice) GetCmmdty() (out Commodity3) {
	if t == nil || t.Cmmdty == nil {
		return
	}
	return *t.Cmmdty
}

type SecurityIdentification26Choice struct {
	Id       *ISINOct2015Identifier `json:"id,omitempty" xml:"Id,omitempty"`
	NotAvlbl *NotAvailable1Code     `json:"notAvlbl,omitempty" xml:"NotAvlbl,omitempty"`
}

func (t *SecurityIdentification26Choice) GetId() (out ISINOct2015Identifier) {
	if t == nil || t.Id == nil {
		return
	}
	return *t.Id
}
func (t *SecurityIdentification26Choice) GetNotAvlbl() (out NotAvailable1Code) {
	if t == nil || t.NotAvlbl == nil {
		return
	}
	return *t.NotAvlbl
}

type SecurityIssuer1 struct {
	LEI          LEIIdentifier `json:"lei,omitempty" xml:"LEI"`
	JursdctnCtry CountryCode   `json:"jursdctnCtry,omitempty" xml:"JursdctnCtry"`
}

func (t *SecurityIssuer1) GetLEI() (out LEIIdentifier) {
	return t.LEI
}
func (t *SecurityIssuer1) GetJursdctnCtry() (out CountryCode) {
	return t.JursdctnCtry
}

type SettlementParties31Choice struct {
	CntrlSctiesDpstryPtcpt *LEIIdentifier `json:"cntrlSctiesDpstryPtcpt,omitempty" xml:"CntrlSctiesDpstryPtcpt,omitempty"`
	IndrctPtcpt            *LEIIdentifier `json:"indrctPtcpt,omitempty" xml:"IndrctPtcpt,omitempty"`
}

func (t *SettlementParties31Choice) GetCntrlSctiesDpstryPtcpt() (out LEIIdentifier) {
	if t == nil || t.CntrlSctiesDpstryPtcpt == nil {
		return
	}
	return *t.CntrlSctiesDpstryPtcpt
}
func (t *SettlementParties31Choice) GetIndrctPtcpt() (out LEIIdentifier) {
	if t == nil || t.IndrctPtcpt == nil {
		return
	}
	return *t.IndrctPtcpt
}

// May be one of GENE, SPEC
type SpecialCollateral1Code string

func GetSpecialCollateral1CodePointer(in string) *SpecialCollateral1Code {
	if in == "" {
		return nil
	}
	out := SpecialCollateral1Code(in)
	return &out
}

type SupplementaryData1 struct {
	PlcAndNm *Max350Text                `json:"plcAndNm,omitempty" xml:"PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `json:"envlp,omitempty" xml:"Envlp"`
}

func (t *SupplementaryData1) GetPlcAndNm() (out Max350Text) {
	if t == nil || t.PlcAndNm == nil {
		return
	}
	return *t.PlcAndNm
}
func (t *SupplementaryData1) GetEnvlp() (out SupplementaryDataEnvelope1) {
	return t.Envlp
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

func (t *SupplementaryDataEnvelope1) Get() (out string) {
	return t.Item
}

type TradeEarlyTermination2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData49   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      LoanData11           `json:"lnData,omitempty" xml:"LnData"`
	CollData    CollateralFlag6      `json:"collData,omitempty" xml:"CollData"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeEarlyTermination2) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeEarlyTermination2) GetCtrPtyData() (out CounterpartyData49) {
	return t.CtrPtyData
}
func (t *TradeEarlyTermination2) GetLnData() (out LoanData11) {
	return t.LnData
}
func (t *TradeEarlyTermination2) GetCollData() (out CollateralFlag6) {
	return t.CollData
}
func (t *TradeEarlyTermination2) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TradeError2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData49   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	UnqTxIdr    Max52Text            `json:"unqTxIdr,omitempty" xml:"UnqTxIdr"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeError2) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeError2) GetCtrPtyData() (out CounterpartyData49) {
	return t.CtrPtyData
}
func (t *TradeError2) GetUnqTxIdr() (out Max52Text) {
	return t.UnqTxIdr
}
func (t *TradeError2) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TradeNewTransaction6 struct {
	TechRcrdId  *Max35Text                        `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48                `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      LoanData44                        `json:"lnData,omitempty" xml:"LnData"`
	CollData    *TransactionCollateralData1Choice `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code            `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1              `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeNewTransaction6) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeNewTransaction6) GetCtrPtyData() (out CounterpartyData48) {
	return t.CtrPtyData
}
func (t *TradeNewTransaction6) GetLnData() (out LoanData44) {
	return t.LnData
}
func (t *TradeNewTransaction6) GetCollData() (out TransactionCollateralData1Choice) {
	if t == nil || t.CollData == nil {
		return
	}
	return *t.CollData
}
func (t *TradeNewTransaction6) GetLvlTp() (out ModificationLevel1Code) {
	return t.LvlTp
}
func (t *TradeNewTransaction6) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TradeTransactionCollateralUpdate2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData49   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      LoanData4            `json:"lnData,omitempty" xml:"LnData"`
	CollData    CollateralData7      `json:"collData,omitempty" xml:"CollData"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeTransactionCollateralUpdate2) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeTransactionCollateralUpdate2) GetCtrPtyData() (out CounterpartyData49) {
	return t.CtrPtyData
}
func (t *TradeTransactionCollateralUpdate2) GetLnData() (out LoanData4) {
	return t.LnData
}
func (t *TradeTransactionCollateralUpdate2) GetCollData() (out CollateralData7) {
	return t.CollData
}
func (t *TradeTransactionCollateralUpdate2) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TradeTransactionCorrection6 struct {
	TechRcrdId  *Max35Text             `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48     `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      *LoanData43            `json:"lnData,omitempty" xml:"LnData,omitempty"`
	CollData    *CollateralData7       `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1   `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeTransactionCorrection6) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeTransactionCorrection6) GetCtrPtyData() (out CounterpartyData48) {
	return t.CtrPtyData
}
func (t *TradeTransactionCorrection6) GetLnData() (out LoanData43) {
	if t == nil || t.LnData == nil {
		return
	}
	return *t.LnData
}
func (t *TradeTransactionCorrection6) GetCollData() (out CollateralData7) {
	if t == nil || t.CollData == nil {
		return
	}
	return *t.CollData
}
func (t *TradeTransactionCorrection6) GetLvlTp() (out ModificationLevel1Code) {
	return t.LvlTp
}
func (t *TradeTransactionCorrection6) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TradeTransactionModification7 struct {
	TechRcrdId  *Max35Text                        `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48                `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      *LoanData42                       `json:"lnData,omitempty" xml:"LnData,omitempty"`
	CollData    *TransactionCollateralData8Choice `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code            `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1              `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeTransactionModification7) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeTransactionModification7) GetCtrPtyData() (out CounterpartyData48) {
	return t.CtrPtyData
}
func (t *TradeTransactionModification7) GetLnData() (out LoanData42) {
	if t == nil || t.LnData == nil {
		return
	}
	return *t.LnData
}
func (t *TradeTransactionModification7) GetCollData() (out TransactionCollateralData8Choice) {
	if t == nil || t.CollData == nil {
		return
	}
	return *t.CollData
}
func (t *TradeTransactionModification7) GetLvlTp() (out ModificationLevel1Code) {
	return t.LvlTp
}
func (t *TradeTransactionModification7) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TradeTransactionPositionComponent2 struct {
	TechRcrdId  *Max35Text             `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48     `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	LnData      *LoanData40            `json:"lnData,omitempty" xml:"LnData,omitempty"`
	CollData    *CollateralData3       `json:"collData,omitempty" xml:"CollData,omitempty"`
	LvlTp       ModificationLevel1Code `json:"lvlTp,omitempty" xml:"LvlTp"`
	SplmtryData []SupplementaryData1   `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeTransactionPositionComponent2) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeTransactionPositionComponent2) GetCtrPtyData() (out CounterpartyData48) {
	return t.CtrPtyData
}
func (t *TradeTransactionPositionComponent2) GetLnData() (out LoanData40) {
	if t == nil || t.LnData == nil {
		return
	}
	return *t.LnData
}
func (t *TradeTransactionPositionComponent2) GetCollData() (out CollateralData3) {
	if t == nil || t.CollData == nil {
		return
	}
	return *t.CollData
}
func (t *TradeTransactionPositionComponent2) GetLvlTp() (out ModificationLevel1Code) {
	return t.LvlTp
}
func (t *TradeTransactionPositionComponent2) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
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

func (t *TradeTransactionReport6Choice) GetNew() (out TradeNewTransaction6) {
	if t == nil || t.New == nil {
		return
	}
	return *t.New
}
func (t *TradeTransactionReport6Choice) GetMod() (out TradeTransactionModification7) {
	if t == nil || t.Mod == nil {
		return
	}
	return *t.Mod
}
func (t *TradeTransactionReport6Choice) GetErr() (out TradeError2) {
	if t == nil || t.Err == nil {
		return
	}
	return *t.Err
}
func (t *TradeTransactionReport6Choice) GetEarlyTermntn() (out TradeEarlyTermination2) {
	if t == nil || t.EarlyTermntn == nil {
		return
	}
	return *t.EarlyTermntn
}
func (t *TradeTransactionReport6Choice) GetPosCmpnt() (out TradeTransactionPositionComponent2) {
	if t == nil || t.PosCmpnt == nil {
		return
	}
	return *t.PosCmpnt
}
func (t *TradeTransactionReport6Choice) GetCollUpd() (out TradeTransactionCollateralUpdate2) {
	if t == nil || t.CollUpd == nil {
		return
	}
	return *t.CollUpd
}
func (t *TradeTransactionReport6Choice) GetCrrctn() (out TradeTransactionCorrection6) {
	if t == nil || t.Crrctn == nil {
		return
	}
	return *t.Crrctn
}
func (t *TradeTransactionReport6Choice) GetValtnUpd() (out TradeValuationUpdate2) {
	if t == nil || t.ValtnUpd == nil {
		return
	}
	return *t.ValtnUpd
}

type TradeValuationUpdate2 struct {
	TechRcrdId  *Max35Text           `json:"techRcrdId,omitempty" xml:"TechRcrdId,omitempty"`
	CtrPtyData  CounterpartyData48   `json:"ctrPtyData,omitempty" xml:"CtrPtyData"`
	UnqTxIdr    Max52Text            `json:"unqTxIdr,omitempty" xml:"UnqTxIdr"`
	EvtDt       ISODate              `json:"evtDt,omitempty" xml:"EvtDt"`
	MktVal      string               `json:"mktVal,omitempty" xml:"MktVal"`
	SplmtryData []SupplementaryData1 `json:"splmtryData,omitempty" xml:"SplmtryData,omitempty"`
}

func (t *TradeValuationUpdate2) GetTechRcrdId() (out Max35Text) {
	if t == nil || t.TechRcrdId == nil {
		return
	}
	return *t.TechRcrdId
}
func (t *TradeValuationUpdate2) GetCtrPtyData() (out CounterpartyData48) {
	return t.CtrPtyData
}
func (t *TradeValuationUpdate2) GetUnqTxIdr() (out Max52Text) {
	return t.UnqTxIdr
}
func (t *TradeValuationUpdate2) GetEvtDt() (out ISODate) {
	return t.EvtDt
}
func (t *TradeValuationUpdate2) GetMktVal() (out string) {
	return t.MktVal
}
func (t *TradeValuationUpdate2) GetSplmtryData() (out []SupplementaryData1) {
	return t.SplmtryData
}

type TransactionCollateralData1Choice struct {
	RpTrad     *Collateral15          `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *Collateral15          `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *CollateralFlag6Choice `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
	MrgnLndg   *Security3             `json:"mrgnLndg,omitempty" xml:"MrgnLndg,omitempty"`
}

func (t *TransactionCollateralData1Choice) GetRpTrad() (out Collateral15) {
	if t == nil || t.RpTrad == nil {
		return
	}
	return *t.RpTrad
}
func (t *TransactionCollateralData1Choice) GetBuySellBck() (out Collateral15) {
	if t == nil || t.BuySellBck == nil {
		return
	}
	return *t.BuySellBck
}
func (t *TransactionCollateralData1Choice) GetSctiesLndg() (out CollateralFlag6Choice) {
	if t == nil || t.SctiesLndg == nil {
		return
	}
	return *t.SctiesLndg
}
func (t *TransactionCollateralData1Choice) GetMrgnLndg() (out Security3) {
	if t == nil || t.MrgnLndg == nil {
		return
	}
	return *t.MrgnLndg
}

type TransactionCollateralData8Choice struct {
	RpTrad     *Collateral27          `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *Collateral27          `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *CollateralFlag6Choice `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
}

func (t *TransactionCollateralData8Choice) GetRpTrad() (out Collateral27) {
	if t == nil || t.RpTrad == nil {
		return
	}
	return *t.RpTrad
}
func (t *TransactionCollateralData8Choice) GetBuySellBck() (out Collateral27) {
	if t == nil || t.BuySellBck == nil {
		return
	}
	return *t.BuySellBck
}
func (t *TransactionCollateralData8Choice) GetSctiesLndg() (out CollateralFlag6Choice) {
	if t == nil || t.SctiesLndg == nil {
		return
	}
	return *t.SctiesLndg
}

type TransactionCounterpartyData3Choice struct {
	RpTrad     *TransactionCounterpartyData7 `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *TransactionCounterpartyData7 `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *TransactionCounterpartyData7 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
}

func (t *TransactionCounterpartyData3Choice) GetRpTrad() (out TransactionCounterpartyData7) {
	if t == nil || t.RpTrad == nil {
		return
	}
	return *t.RpTrad
}
func (t *TransactionCounterpartyData3Choice) GetBuySellBck() (out TransactionCounterpartyData7) {
	if t == nil || t.BuySellBck == nil {
		return
	}
	return *t.BuySellBck
}
func (t *TransactionCounterpartyData3Choice) GetSctiesLndg() (out TransactionCounterpartyData7) {
	if t == nil || t.SctiesLndg == nil {
		return
	}
	return *t.SctiesLndg
}

type TransactionCounterpartyData7 struct {
	Bnfcry     *OrganisationIdentification9Choice `json:"bnfcry,omitempty" xml:"Bnfcry,omitempty"`
	TrptyAgt   *LEIIdentifier                     `json:"trptyAgt,omitempty" xml:"TrptyAgt,omitempty"`
	Brkr       *LEIIdentifier                     `json:"brkr,omitempty" xml:"Brkr,omitempty"`
	ClrMmb     *LEIIdentifier                     `json:"clrMmb,omitempty" xml:"ClrMmb,omitempty"`
	SttlmPties *SettlementParties31Choice         `json:"sttlmPties,omitempty" xml:"SttlmPties,omitempty"`
	AgtLndr    *LEIIdentifier                     `json:"agtLndr,omitempty" xml:"AgtLndr,omitempty"`
}

func (t *TransactionCounterpartyData7) GetBnfcry() (out OrganisationIdentification9Choice) {
	if t == nil || t.Bnfcry == nil {
		return
	}
	return *t.Bnfcry
}
func (t *TransactionCounterpartyData7) GetTrptyAgt() (out LEIIdentifier) {
	if t == nil || t.TrptyAgt == nil {
		return
	}
	return *t.TrptyAgt
}
func (t *TransactionCounterpartyData7) GetBrkr() (out LEIIdentifier) {
	if t == nil || t.Brkr == nil {
		return
	}
	return *t.Brkr
}
func (t *TransactionCounterpartyData7) GetClrMmb() (out LEIIdentifier) {
	if t == nil || t.ClrMmb == nil {
		return
	}
	return *t.ClrMmb
}
func (t *TransactionCounterpartyData7) GetSttlmPties() (out SettlementParties31Choice) {
	if t == nil || t.SttlmPties == nil {
		return
	}
	return *t.SttlmPties
}
func (t *TransactionCounterpartyData7) GetAgtLndr() (out LEIIdentifier) {
	if t == nil || t.AgtLndr == nil {
		return
	}
	return *t.AgtLndr
}

type TransactionLoanData4Choice struct {
	RpTrad     *LoanData29 `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *LoanData29 `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *LoanData29 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
	MrgnLndg   *LoanData13 `json:"mrgnLndg,omitempty" xml:"MrgnLndg,omitempty"`
}

func (t *TransactionLoanData4Choice) GetRpTrad() (out LoanData29) {
	if t == nil || t.RpTrad == nil {
		return
	}
	return *t.RpTrad
}
func (t *TransactionLoanData4Choice) GetBuySellBck() (out LoanData29) {
	if t == nil || t.BuySellBck == nil {
		return
	}
	return *t.BuySellBck
}
func (t *TransactionLoanData4Choice) GetSctiesLndg() (out LoanData29) {
	if t == nil || t.SctiesLndg == nil {
		return
	}
	return *t.SctiesLndg
}
func (t *TransactionLoanData4Choice) GetMrgnLndg() (out LoanData13) {
	if t == nil || t.MrgnLndg == nil {
		return
	}
	return *t.MrgnLndg
}

type TransactionLoanData5Choice struct {
	RpTrad     *LoanData1  `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *LoanData2  `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *LoanData39 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
	MrgnLndg   *LoanData32 `json:"mrgnLndg,omitempty" xml:"MrgnLndg,omitempty"`
}

func (t *TransactionLoanData5Choice) GetRpTrad() (out LoanData1) {
	if t == nil || t.RpTrad == nil {
		return
	}
	return *t.RpTrad
}
func (t *TransactionLoanData5Choice) GetBuySellBck() (out LoanData2) {
	if t == nil || t.BuySellBck == nil {
		return
	}
	return *t.BuySellBck
}
func (t *TransactionLoanData5Choice) GetSctiesLndg() (out LoanData39) {
	if t == nil || t.SctiesLndg == nil {
		return
	}
	return *t.SctiesLndg
}
func (t *TransactionLoanData5Choice) GetMrgnLndg() (out LoanData32) {
	if t == nil || t.MrgnLndg == nil {
		return
	}
	return *t.MrgnLndg
}

type TransactionLoanData6Choice struct {
	RpTrad     *LoanData37 `json:"rpTrad,omitempty" xml:"RpTrad,omitempty"`
	BuySellBck *LoanData27 `json:"buySellBck,omitempty" xml:"BuySellBck,omitempty"`
	SctiesLndg *LoanData41 `json:"sctiesLndg,omitempty" xml:"SctiesLndg,omitempty"`
}

func (t *TransactionLoanData6Choice) GetRpTrad() (out LoanData37) {
	if t == nil || t.RpTrad == nil {
		return
	}
	return *t.RpTrad
}
func (t *TransactionLoanData6Choice) GetBuySellBck() (out LoanData27) {
	if t == nil || t.BuySellBck == nil {
		return
	}
	return *t.BuySellBck
}
func (t *TransactionLoanData6Choice) GetSctiesLndg() (out LoanData41) {
	if t == nil || t.SctiesLndg == nil {
		return
	}
	return *t.SctiesLndg
}

// May be one of PIEC, TONS, FOOT, GBGA, USGA, GRAM, INCH, KILO, PUND, METR, CMET, MMET, LITR, CELI, MILI, GBOU, USOU, GBQA, USQA, GBPI, USPI, MILE, KMET, YARD, SQKI, HECT, ARES, SMET, SCMT, SMIL, SQMI, SQYA, SQFO, SQIN, ACRE
type UnitOfMeasure1Code string

func GetUnitOfMeasure1CodePointer(in string) *UnitOfMeasure1Code {
	if in == "" {
		return nil
	}
	out := UnitOfMeasure1Code(in)
	return &out
}

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
