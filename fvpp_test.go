package fvpp

import (
	"fmt"
	"testing"
)

type user struct {
	Name    string
	Married bool
	Cars    int
	Deposit float64
}

var u = user{
	Name:    "bili",
	Married: true,
	Cars:    2,
	Deposit: 200000,
}

func TestPretty(t *testing.T) {
	tests := []struct {
		user
		want string
	}{
		{
			user: u,
			want: "[Name:'bili'(string) Married:'true'(bool) Cars:'2'(int) Deposit:'200000'(float64)]",
		},
		{
			user: user{
				Name:    "kili",
				Married: false,
				Cars:    0,
				Deposit: 20000,
			},
			want: "[Name:'kili'(string) Married:'false'(bool) Cars:'0'(int) Deposit:'20000'(float64)]",
		},
	}

	for _, tt := range tests {
		got := Pretty(tt.user)
		if tt.want != got {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}

func TestCustomFormat(t *testing.T) {
	t.Run("BuiltinType", testBuiltinType)
	t.Run("StructType", testStructType)
}

func testStructType(t *testing.T) {
	tests := []struct {
		Face string
		Eye  string
		User user
		want string
	}{
		{
			Face: "yellow",
			Eye:  "yellow",
			User: u,
			want: "",
		},
		{
			Face: "black",
			Eye:  "blue",
			User: u,
			want: "",
		},
	}

	for _, tt := range tests {
		got := Pretty(tt)
		if got != tt.want {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}

func testBuiltinType(t *testing.T) {
	tests := []struct {
		fieldFormat    string
		fieldSeparator string
		wrapperBegin   string
		wrapperEnd     string
		want           string
	}{
		{
			fieldFormat:    fmt.Sprintf("%s-%s-%s", Field, Type, Value),
			fieldSeparator: ";",
			wrapperBegin:   "<",
			wrapperEnd:     ">",
			want:           "<Name-string-bili;Married-bool-true;Cars-int-2;Deposit-float64-200000>",
		},
		{
			fieldFormat:    fmt.Sprintf("%s>%s>%s", Field, Value, Type),
			fieldSeparator: "|",
			wrapperBegin:   "(",
			wrapperEnd:     ")",
			want:           "(Name>bili>string|Married>true>bool|Cars>2>int|Deposit>200000>float64)",
		},
		{
			fieldFormat:    fmt.Sprintf("%s#%s#%s", Type, Field, Value),
			fieldSeparator: " ",
			wrapperBegin:   "{",
			wrapperEnd:     "}",
			want:           "{string#Name#bili bool#Married#true int#Cars#2 float64#Deposit#200000}",
		},
	}

	for _, tt := range tests {
		got := CustomFormat(u, tt.fieldFormat, tt.fieldSeparator, tt.wrapperBegin, tt.wrapperEnd)
		if tt.want != got {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}
