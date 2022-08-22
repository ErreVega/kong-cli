package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"
)

func TablePrint[T any](obj []T, fields ...string) {

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 5, ' ', 0)
	header := ""
	for _, field := range fields {
		header += field + "\t"
	}
	fmt.Fprintln(w, header)

	for _, element := range obj {
		bodyElem := ""
		for _, field := range fields {
			value := getField(&element, field)
			bodyElem += value + "\t"
		}

		fmt.Fprintln(w, bodyElem)
	}
	w.Flush()
}

func getField[T any](v *T, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)

	var value string
	switch f.Kind() {
	case reflect.Int:
		value = strconv.FormatInt(f.Int(), 10)
	case reflect.String:
		value = f.String()
	case reflect.Bool:
		value = strconv.FormatBool(f.Bool())
	case reflect.Struct, reflect.Interface, reflect.Array, reflect.Slice:
		value = fmt.Sprintf("%v", f)
	case reflect.Invalid:
		value = "<field not found>"
	default:
		value = fmt.Sprintf("<unknown %v>", f)
	}
	return string(value)
}
