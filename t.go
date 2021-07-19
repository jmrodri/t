package main

import (
	"fmt"
	"io/ioutil"

	"github.com/lestrrat/go-libxml2/parser"
	"github.com/lestrrat/go-libxml2/xsd"
)

func main() {

	xsdsrc, err := ioutil.ReadFile("JUnit.xsd")
	if err != nil {
		panic(err)
	}

	schema, err := xsd.Parse(xsdsrc)
	if err != nil {
		panic(err)
	}

	// xmlfile, err := ioutil.ReadFile("xunit_results.xml")
	xmlfile, err := ioutil.ReadFile("sample.xml")
	if err != nil {
		panic(err)
	}

	daparser := parser.New()
	doc, err := daparser.Parse(xmlfile)
	if err != nil {
		panic(err)
	}

	defer schema.Free()
	if err := schema.Validate(doc); err != nil {
		// if err != nil {
		//     fmt.Println(err)
		// }

		for _, e := range err.(xsd.SchemaValidationError).Errors() {
			fmt.Println(e.Error())
			// fmt.Println(e)
		}
	}

}
