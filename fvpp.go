package fvpp

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	FlagField = "field"
	FlagValue = "value"
	FlagType  = "type"
)

// Pretty return the default format string,as: `[Name:'xiaosongfu'(string) Cars:'1'(int)]`
func Pretty(value interface{}) string {
	return CustomFormat(value, fmt.Sprintf("%s:'%s'(%s)", FlagField, FlagValue, FlagType), " ", "[", "]")
}

// Println use fmt.Println function to print the Pretty() function's result
func Println(value interface{}) {
	fmt.Println(Pretty(value))
}

// CustomFormat is used for customize the format、 separator and wrapper with what you like
func CustomFormat(value interface{}, fieldFormat, fieldSeparator, wrapperBegin, wrapperEnd string) string {
	format := strings.ReplaceAll(fieldFormat, FlagField, "%v")
	format = strings.ReplaceAll(format, FlagValue, "%v")
	format = strings.ReplaceAll(format, FlagType, "%v")

	var tmp []string

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	n := t.NumField()
	for i := 0; i < n; i++ {
		tmp = append(tmp, fmt.Sprintf(format, t.Field(i).Name, v.Field(i).Interface(), t.Field(i).Type))
	}
	return fmt.Sprintf("%s%s%s", wrapperBegin, strings.Join(tmp, fieldSeparator), wrapperEnd)
}
