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
			// first,we assume the field is builtin type,so set it `rValue.Field(i).Interface()`
			// but if it is a struct,we need recursive each field to got it's value use current function
			fieldValue := rValue.Field(i).Interface()

			if rType.Field(i).Type.Kind() == reflect.Struct {
				fieldValue = CustomFormat(rValue.Field(i).Interface(), fieldFormat, fieldSeparator, wrapperBegin, wrapperEnd)
			}

			f := fmt.Sprintf("%v", rType.Field(i).Name)
			v := fmt.Sprintf("%v", fieldValue)
			t := fmt.Sprintf("%v", rType.Field(i).Type)

			format := fieldFormat
			format = strings.ReplaceAll(format, Field, f)
			format = strings.ReplaceAll(format, Value, v)
			format = strings.ReplaceAll(format, Type, t)

			tmp = append(tmp, format)
		}
		return assembleResult(strings.Join(tmp, fieldSeparator), wrapperBegin, wrapperEnd)
	} else { // builtin types
		//content :=
		return assembleResult(fmt.Sprintf("%v", value), wrapperBegin, wrapperEnd)
	}
}

// assembleResult assemble result
func assembleResult(content, wrapperBegin, wrapperEnd string) string {
	return fmt.Sprintf("%s%s%s", wrapperBegin, content, wrapperEnd)
}
