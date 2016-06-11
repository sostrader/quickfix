package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"text/template"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
)

var (
	tagTemplate   *template.Template
	fieldTemplate *template.Template
	enumTemplate  *template.Template
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fields [flags] <path to data dictionary> ... \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func genEnums(fieldTypes []*datadictionary.FieldType) error {
	writer := new(bytes.Buffer)

	if err := enumTemplate.Execute(writer, fieldTypes); err != nil {
		return err
	}

	return gen.WriteFile("enum/enums.go", writer.String())
}

func quickfixValueType(quickfixType string) (valueType string) {
	switch quickfixType {
	case "FIXString":
		valueType = "string"
	case "FIXBoolean":
		valueType = "bool"
	case "FIXInt":
		valueType = "int"
	case "FIXUTCTimestamp":
		valueType = "time.Time"
	case "FIXFloat":
		valueType = "float64"
	}

	return
}

func quickfixType(field *datadictionary.FieldType) (quickfixType string, err error) {
	switch field.Type {
	case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
		fallthrough
	case "MULTIPLECHARVALUE":
		fallthrough
	case "CHAR":
		fallthrough
	case "CURRENCY":
		fallthrough
	case "DATA":
		fallthrough
	case "MONTHYEAR":
		fallthrough
	case "LOCALMKTDATE":
		fallthrough
	case "TIME":
		fallthrough
	case "DATE":
		fallthrough
	case "EXCHANGE":
		fallthrough
	case "LANGUAGE":
		fallthrough
	case "XMLDATA":
		fallthrough
	case "COUNTRY":
		fallthrough
	case "UTCTIMEONLY":
		fallthrough
	case "UTCDATE":
		fallthrough
	case "UTCDATEONLY":
		fallthrough
	case "TZTIMEONLY":
		fallthrough
	case "TZTIMESTAMP":
		fallthrough
	case "STRING":
		quickfixType = "FIXString"

	case "BOOLEAN":
		quickfixType = "FIXBoolean"

	case "LENGTH":
		fallthrough
	case "DAYOFMONTH":
		fallthrough
	case "NUMINGROUP":
		fallthrough
	case "SEQNUM":
		fallthrough
	case "INT":
		quickfixType = "FIXInt"

	case "UTCTIMESTAMP":
		quickfixType = "FIXUTCTimestamp"

	case "QTY":
		fallthrough
	case "QUANTITY":
		fallthrough
	case "AMT":
		fallthrough
	case "PRICE":
		fallthrough
	case "PRICEOFFSET":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "FLOAT":
		quickfixType = "FIXFloat"

	default:
		err = fmt.Errorf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag())
	}

	return
}

func genFields(fieldTypes []*datadictionary.FieldType) error {
	writer := new(bytes.Buffer)

	if err := fieldTemplate.Execute(writer, fieldTypes); err != nil {
		return err
	}

	return gen.WriteFile("field/fields.go", writer.String())
}

func genTags(fieldTypes []*datadictionary.FieldType) error {
	writer := new(bytes.Buffer)

	if err := tagTemplate.Execute(writer, fieldTypes); err != nil {
		return err
	}

	return gen.WriteFile("tag/tag_numbers.go", writer.String())
}

func init() {
	tmplFuncs := template.FuncMap{
		"importRootPath":    gen.GetImportPathRoot,
		"quickfixType":      quickfixType,
		"quickfixValueType": quickfixValueType,
	}

	tagTemplate = template.Must(template.New("Tag").Parse(`
package tag
import("github.com/quickfixgo/quickfix")

const (
{{- range .}}
{{ .Name }} quickfix.Tag =  {{ .Tag }}
{{- end }}
)
	`))

	fieldTemplate = template.Must(template.New("Field").Funcs(tmplFuncs).Parse(`
package field
import(
	"github.com/quickfixgo/quickfix"
	"{{ importRootPath }}/tag"

	"time"
)

{{ range . }}
{{- $base_type := quickfixType . -}}
//{{ .Name }}Field is a {{ .Type }} field
type {{ .Name }}Field struct { quickfix.{{ $base_type }} }

//Tag returns tag.{{ .Name }} ({{ .Tag }})
func (f {{ .Name }}Field) Tag() quickfix.Tag { return tag.{{ .Name }} }

{{ if eq $base_type "FIXUTCTimestamp" }} 
//New{{ .Name }} returns a new {{ .Name }}Field initialized with val
func New{{ .Name }}(val time.Time) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXUTCTimestamp{ Time: val } } 
}

//New{{ .Name }}NoMillis returns a new {{ .Name }}Field initialized with val without millisecs
func New{{ .Name }}NoMillis(val time.Time) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXUTCTimestamp{ Time: val, NoMillis: true } } 
}

{{ else }}
//New{{ .Name }} returns a new {{ .Name }}Field initialized with val
func New{{ .Name }}(val {{ quickfixValueType $base_type }}) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.{{ $base_type }}(val) }
}
{{ end }}{{ end }}
`))

	enumTemplate = template.Must(template.New("Enum").Parse(`
package enum
{{ range $ft := . }}
{{ if $ft.Enums }} 
//Enum values for {{ $ft.Name }}
const(
{{ range $ft.Enums }}
{{ $ft.Name }}_{{ .Description }} = "{{ .Value }}"
{{- end }}
)
{{ end }}{{ end }}
	`))
}

//sort fieldtypes by name
type byFieldName []*datadictionary.FieldType

func (n byFieldName) Len() int           { return len(n) }
func (n byFieldName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byFieldName) Less(i, j int) bool { return n[i].Name() < n[j].Name() }

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	fieldTypeMap := make(map[string]*datadictionary.FieldType)

	for _, dataDict := range flag.Args() {
		spec, err := datadictionary.Parse(dataDict)

		if err != nil {
			fmt.Println(dataDict)
			panic(err)
		}

		for _, field := range spec.FieldTypeByTag {
			if oldField, ok := fieldTypeMap[field.Name()]; ok {
				//merge old enums with new
				if len(oldField.Enums) > 0 && field.Enums == nil {
					field.Enums = make(map[string]datadictionary.Enum)
				}

				for enumVal, enum := range oldField.Enums {
					if _, ok := field.Enums[enumVal]; !ok {
						//Verify an existing enum doesn't have the same description. Keep newer enum
						okToKeepEnum := true
						for _, newEnum := range field.Enums {
							if newEnum.Description == enum.Description {
								okToKeepEnum = false
								break
							}
						}

						if okToKeepEnum {
							field.Enums[enumVal] = enum
						}
					}
				}
			}

			fieldTypeMap[field.Name()] = field
		}
	}

	fieldTypes := make([]*datadictionary.FieldType, len(fieldTypeMap))
	i := 0
	for _, fieldType := range fieldTypeMap {
		fieldTypes[i] = fieldType
		i++
	}

	sort.Sort(byFieldName(fieldTypes))

	var h gen.ErrorHandler
	h.Handle(genTags(fieldTypes))
	h.Handle(genFields(fieldTypes))
	h.Handle(genEnums(fieldTypes))
	os.Exit(h.ReturnCode)
}
