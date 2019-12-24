package fvpp

import (
	"testing"
)

type car struct {
	Brand    string
	Wheel    int
	Price    float64
	Scrapped bool
}

type user struct {
	Name    string
	Married bool
	Cars    int
	Deposit float64
	Car     car
}

type people struct {
	Face string
	Eye  string
	User user
}

var u = user{
	Name:    "bili",
	Married: true,
	Cars:    2,
	Deposit: 200000,
	Car:     c,
}

var c = car{
	Brand:    "BMW",
	Wheel:    4,
	Price:    280000,
	Scrapped: false,
}

var p = people{
	Face: "black",
	Eye:  "blue",
	User: u,
}

func TestPretty(t *testing.T) {
	t.Run("BuiltinType", testBuiltinType)
	t.Run("StructType", testStructType)
	t.Run("NestedStructType", testNestedStructType)
	t.Run("MultilayerNestedStructType", testMultilayerNestedStructType)
}

func testMultilayerNestedStructType(t *testing.T) {
	tests := []struct {
		people people
		want   string
	}{
		{
			people: p,
			want:   "[Face:'black'(string) Eye:'blue'(string) User:'[Name:'bili'(string) Married:'true'(bool) Cars:'2'(int) Deposit:'200000'(float64) Car:'[Brand:'BMW'(string) Wheel:'4'(int) Price:'280000'(float64) Scrapped:'false'(bool)]'(fvpp.car)]'(fvpp.user)]",
		},
		{
			people: people{
				Face: "yellow",
				Eye:  "yellow",
				User: user{
					Name:    "vik",
					Married: false,
					Cars:    0,
					Deposit: 10000,
					Car: car{
						Brand:    "XiaoNiu",
						Wheel:    2,
						Price:    6000,
						Scrapped: false,
					},
				},
			},
			want: "[Face:'yellow'(string) Eye:'yellow'(string) User:'[Name:'vik'(string) Married:'false'(bool) Cars:'0'(int) Deposit:'10000'(float64) Car:'[Brand:'XiaoNiu'(string) Wheel:'2'(int) Price:'6000'(float64) Scrapped:'false'(bool)]'(fvpp.car)]'(fvpp.user)]",
		},
	}

	for _, tt := range tests {
		got := Pretty(tt.people)
		if got != tt.want {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}

func testNestedStructType(t *testing.T) {
	tests := []struct {
		user user
		want string
	}{
		{
			user: u,
			want: "[Name:'bili'(string) Married:'true'(bool) Cars:'2'(int) Deposit:'200000'(float64) Car:'[Brand:'BMW'(string) Wheel:'4'(int) Price:'280000'(float64) Scrapped:'false'(bool)]'(fvpp.car)]",
		},
		{
			user: user{
				Name:    "bill",
				Married: true,
				Cars:    100,
				Deposit: 100000000,
				Car: car{
					Brand:    "Masaladi",
					Wheel:    6,
					Price:    10000000,
					Scrapped: false,
				},
			},
			want: "[Name:'bill'(string) Married:'true'(bool) Cars:'100'(int) Deposit:'1e+08'(float64) Car:'[Brand:'Masaladi'(string) Wheel:'6'(int) Price:'1e+07'(float64) Scrapped:'false'(bool)]'(fvpp.car)]",
		},
	}

	for _, tt := range tests {
		got := Pretty(tt.user)
		if got != tt.want {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}

func testStructType(t *testing.T) {
	tests := []struct {
		car
		want string
	}{
		{
			car:  c,
			want: "[Brand:'BMW'(string) Wheel:'4'(int) Price:'280000'(float64) Scrapped:'false'(bool)]",
		},
		{
			car: car{
				Brand:    "Audi",
				Wheel:    2,
				Price:    100000,
				Scrapped: false,
			},
			want: "[Brand:'Audi'(string) Wheel:'2'(int) Price:'100000'(float64) Scrapped:'false'(bool)]",
		},
	}

	for _, tt := range tests {
		got := Pretty(tt.car)
		if tt.want != got {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}

func testBuiltinType(t *testing.T) {
	intValue := 2
	strValue := 3
	mapValue := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	intWant := "not a struct"
	strWant := "not a struct"
	mapWant := "not a struct"

	got := ""
	if got = Pretty(intValue); got != intWant {
		t.Errorf("test failed, want: %s, but got: %s", intWant, got)
	}
	if got = Pretty(strValue); got != strWant {
		t.Errorf("test failed, want: %s, but got: %s", strWant, got)
	}
	if got = Pretty(mapValue); got != mapWant {
		t.Errorf("test failed, want: %s, but got: %s", mapWant, got)
	}
}

//func TestCustomFormat(t *testing.T) {
//}
//
//func testStructTypeCustomFormat(t *testing.T) {
//	tests := []struct {
//		fieldFormat    string
//		fieldSeparator string
//		wrapperBegin   string
//		wrapperEnd     string
//		want           string
//	}{
//		{
//			fieldFormat:    fmt.Sprintf("%s-%s-%s", Field, Type, Value),
//			fieldSeparator: ";",
//			wrapperBegin:   "<",
//			wrapperEnd:     ">",
//			want:           "",
//		},
//		{
//			fieldFormat:    fmt.Sprintf("%s>%s>%s", Field, Value, Type),
//			fieldSeparator: "|",
//			wrapperBegin:   "(",
//			wrapperEnd:     ")",
//			want:           "",
//		},
//		{
//			fieldFormat:    fmt.Sprintf("%s#%s#%s", Type, Field, Value),
//			fieldSeparator: " ",
//			wrapperBegin:   "{",
//			wrapperEnd:     "}",
//			want:           "",
//		},
//	}
//
//	for _, tt := range tests {
//		got := CustomFormat(c, tt.fieldFormat, tt.fieldSeparator, tt.wrapperBegin, tt.wrapperEnd)
//		if tt.want != got {
//			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
//		}
//	}
//}
