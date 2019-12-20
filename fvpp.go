package fvpp

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	Field = "field"
	Value = "value"
	Type  = "type"
)

// Pretty return the default format string,as: `[Name:'xiaosongfu'(string) Cars:'1'(int)]`
func Pretty(value interface{}) string {
	return CustomFormat(value, fmt.Sprintf("%s:'%s'(%s)", Field, Value, Type), " ", "[", "]")
}

// Println use fmt.Println function to print the Pretty() function's result
func Println(value interface{}) {
	fmt.Println(Pretty(value))
}

// CustomFormat is used for customize the format、 separator and wrapper with what you like
func CustomFormat(value interface{}, fieldFormat, fieldSeparator, wrapperBegin, wrapperEnd string) string {
	var tmp []string

	rType := reflect.TypeOf(value)
	rValue := reflect.ValueOf(value)

	n := rType.NumField()
	for i := 0; i < n; i++ {
		f := fmt.Sprintf("%v", rType.Field(i).Name)
		v := fmt.Sprintf("%v", rValue.Field(i).Interface())
		t := fmt.Sprintf("%v", rType.Field(i).Type)

		format := fieldFormat
		format = strings.ReplaceAll(format, Field, f)
		format = strings.ReplaceAll(format, Value, v)
		format = strings.ReplaceAll(format, Type, t)

		tmp = append(tmp, format)
	}
	return fmt.Sprintf("%s%s%s", wrapperBegin, strings.Join(tmp, fieldSeparator), wrapperEnd)
}
