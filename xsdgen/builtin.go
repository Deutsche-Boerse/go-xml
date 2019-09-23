package xsdgen

import (
	"go/ast"

	"aqwari.net/xml/xsd"
)

const timeXSD = "timeXSD"
const timeXSDCreator = "makeTimeXSDPointer"
const datetimeWithTZ = "2006-01-02T15:04:05.999999999Z07:00"


var timeTypes = map[xsd.Builtin]string{
	xsd.Date:       "2006-01-02",
	xsd.DateTime:   "2006-01-02T15:04:05.999999999",
	xsd.GDay:       "---02",
	xsd.GMonth:     "--01",
	xsd.GMonthDay:  "--01-02",
	xsd.GYear:      "2006",
	xsd.GYearMonth: "2006-01",
	xsd.Time:       "15:04:05.999999999",
}

func isTimeTypeBuiltin(t xsd.Type) bool {
	b, ok := t.(xsd.Builtin)
	if !ok {
		return false
	}
	_, ok = timeTypes[b]
	return ok
}

func builtinExpr(cfg *Config, b xsd.Builtin) ast.Expr {
	if int(b) > len(builtinTbl) || b < 0 {
		return nil
	}

	if !cfg.decimalsAsString {
		// default behaviour
		return builtinTbl[b]
	}

	switch b {
	case xsd.Float, xsd.Double, xsd.Decimal:
		return &ast.Ident{Name: "string"}
	default:
		return builtinTbl[b]
	}
}

// Returns true if t is an xsd.Builtin that is not trivially mapped to a
// builtin Go type; it requires additional marshal/unmarshal methods.
func nonTrivialBuiltin(t xsd.Type) bool {
	b, ok := t.(xsd.Builtin)
	if !ok {
		return false
	}
	switch b {
	case xsd.Base64Binary, xsd.HexBinary,
		xsd.Date, xsd.Time, xsd.DateTime,
		xsd.GDay, xsd.GMonth, xsd.GMonthDay, xsd.GYear, xsd.GYearMonth:
		return true
	}
	return false
}

// The 45 built-in types of the XSD schema
var builtinTbl = []ast.Expr{
	xsd.AnyType:      &ast.Ident{Name: "string"},
	xsd.ENTITIES:     &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
	xsd.ENTITY:       &ast.Ident{Name: "string"},
	xsd.ID:           &ast.Ident{Name: "string"},
	xsd.IDREF:        &ast.Ident{Name: "string"},
	xsd.IDREFS:       &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
	xsd.NCName:       &ast.Ident{Name: "string"},
	xsd.NMTOKEN:      &ast.Ident{Name: "string"},
	xsd.NMTOKENS:     &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
	xsd.NOTATION:     &ast.ArrayType{Elt: &ast.Ident{Name: "string"}},
	xsd.Name:         &ast.Ident{Name: "string"},
	xsd.QName:        &ast.Ident{Name: "xml.Name"},
	xsd.AnyURI:       &ast.Ident{Name: "string"},
	xsd.Base64Binary: &ast.ArrayType{Elt: &ast.Ident{Name: "byte"}},
	xsd.Boolean:      &ast.Ident{Name: "bool"},
	xsd.Byte:         &ast.Ident{Name: "byte"},
	xsd.Date:         &ast.Ident{Name: "XSDDate"},
	xsd.DateTime:     &ast.Ident{Name: "XSDDateTime"},
	xsd.Decimal:      &ast.Ident{Name: "float64"},
	xsd.Double:       &ast.Ident{Name: "float64"},
	// the "duration" built-in is especially broken, so we
	// don't parse it at all.
	xsd.Duration:           &ast.Ident{Name: "string"},
	xsd.Float:              &ast.Ident{Name: "float64"},
	xsd.GDay:               &ast.Ident{Name: "XSDGDay"},
	xsd.GMonth:             &ast.Ident{Name: "XSDGMonth"},
	xsd.GMonthDay:          &ast.Ident{Name: "XSDGMonthDay"},
	xsd.GYear:              &ast.Ident{Name: "XSDGYear"},
	xsd.GYearMonth:         &ast.Ident{Name: "XSDGYearMonth"},
	xsd.HexBinary:          &ast.ArrayType{Elt: &ast.Ident{Name: "byte"}},
	xsd.Int:                &ast.Ident{Name: "int"},
	xsd.Integer:            &ast.Ident{Name: "int"},
	xsd.Language:           &ast.Ident{Name: "string"},
	xsd.Long:               &ast.Ident{Name: "int64"},
	xsd.NegativeInteger:    &ast.Ident{Name: "int"},
	xsd.NonNegativeInteger: &ast.Ident{Name: "int"},
	xsd.NormalizedString:   &ast.Ident{Name: "string"},
	xsd.NonPositiveInteger: &ast.Ident{Name: "int"},
	xsd.PositiveInteger:    &ast.Ident{Name: "int"},
	xsd.Short:              &ast.Ident{Name: "int"},
	xsd.String:             &ast.Ident{Name: "string"},
	xsd.Time:               &ast.Ident{Name: "XSDTime"},
	xsd.Token:              &ast.Ident{Name: "string"},
	xsd.UnsignedByte:       &ast.Ident{Name: "byte"},
	xsd.UnsignedInt:        &ast.Ident{Name: "uint"},
	xsd.UnsignedLong:       &ast.Ident{Name: "uint64"},
	xsd.UnsignedShort:      &ast.Ident{Name: "uint"},
}
