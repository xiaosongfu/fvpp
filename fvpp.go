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
// if value's type is builtin type,return ...
// if value's type is struct,
func CustomFormat(value interface{}, fieldFormat, fieldSeparator, wrapperBegin, wrapperEnd string) string {
	rType := reflect.TypeOf(value)
	rValue := reflect.ValueOf(value)

	if rType.Kind() == reflect.Struct { // struct type
		var tmp []string
		n := rType.NumField()

		for i := 0; i < n; i++ {
			// if the field is builtin type,just set it `rValue.Field(i).Interface()`
			// but if it is a struct,we need recursive each field to got it's value use current function
			fieldValue := ""

			if rType.Field(i).Type.Kind() == reflect.Struct {
				fieldValue = CustomFormat(rValue.Field(i).Interface(), fieldFormat, fieldSeparator, wrapperBegin, wrapperEnd)
			} else {
				fieldValue = fmt.Sprintf("%v", rValue.Field(i).Interface())
			}

			f := fmt.Sprintf("%v", rType.Field(i).Name)
			v := fmt.Sprintf("%s", fieldValue)
			t := fmt.Sprintf("%v", rType.Field(i).Type)

			content := fieldFormat
			content = strings.ReplaceAll(content, Field, f)
			content = strings.ReplaceAll(content, Value, v)
			content = strings.ReplaceAll(content, Type, t)

			tmp = append(tmp, content)
		}
		return assembleResult(strings.Join(tmp, fieldSeparator), wrapperBegin, wrapperEnd)
	} else { // builtin types
		return "not a struct"
	}
}

// assembleResult assemble result
func assembleResult(content, wrapperBegin, wrapperEnd string) string {
	return fmt.Sprintf("%s%s%s", wrapperBegin, content, wrapperEnd)
}
