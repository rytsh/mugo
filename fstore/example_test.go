package fstore_test

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/rytsh/mugo/fstore"
)

func Example() {
	tpl := template.New("test").Funcs(fstore.FuncMap())

	output := &bytes.Buffer{}
	tplParsed, err := tpl.Parse(`{{ $v := codec.JsonDecode (codec.StringToByte .) }}{{ $v.data.name }}`)
	if err != nil {
		log.Fatal(err)
	}

	if err := tplParsed.Execute(output, `{"data": {"name": "Hatay"}}`); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
	// Output:
	// Hatay
}

func Example_sprig() {
	tpl := template.New("test").Funcs(fstore.FuncMap(
		fstore.WithSpecificGroups("sprig"),
	))

	output := &bytes.Buffer{}
	tplParsed, err := tpl.Parse(`{{b64dec "TWVyaGFiYQ=="}}`)
	if err != nil {
		log.Fatal(err)
	}

	if err := tplParsed.Execute(output, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
	// Output:
	// Merhaba
}

func Example_simple() {
	tpl := template.New("test").Funcs(fstore.FuncMap())

	output := &bytes.Buffer{}
	tplParsed, err := tpl.Parse(`{{nothing "nothing for nothing" true 42}}`)
	if err != nil {
		log.Fatal(err)
	}

	if err := tplParsed.Execute(output, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
	// Output:
	//
}

func Example_execTemplate() {
	tpl := template.New("test")
	tpl.Funcs(fstore.FuncMap(
		fstore.WithSpecificFuncs("execTemplate"),
		fstore.WithExecuteTemplate(tpl),
	))

	output := &bytes.Buffer{}
	tplParsed, err := tpl.Parse(`{{ define "ochtend" }}Dag!{{ end }}{{ execTemplate "ochtend" nil | printf }}`)
	if err != nil {
		log.Fatal(err)
	}

	if err := tplParsed.Execute(output, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
	// Output:
	// Dag!
}
