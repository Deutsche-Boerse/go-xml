package timetest

//go:generate go run ../../xsdgen/xsdgen.go -decimalsAsString -addGetMethods -marshalDatetimeWithTimezone -o timetest.go -vv -pkg timetest timetest.xsd
