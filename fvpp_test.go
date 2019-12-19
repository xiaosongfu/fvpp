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

func TestString(t *testing.T) {
	tests := []struct {
		user
		want string
	}{
		{
			user: user{
				Name:    "kili",
				Married: false,
				Cars:    0,
				Deposit: 20000,
			},
			want: "[Name:'kili'(string) Married:'false'(bool) Cars:'0'(int) Deposit:'20000'(float64)]",
		},
		{
			user: user{
				Name:    "bili",
				Married: true,
				Cars:    2,
				Deposit: 200000,
			},
			want: "[Name:'bili'(string) Married:'true'(bool) Cars:'2'(int) Deposit:'200000'(float64)]",
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
	u := user{
		Name:    "bili",
		Married: true,
		Cars:    2,
		Deposit: 200000,
	}
	tests := []struct {
		fieldFormat    string
		fieldSeparator string
		wrapperBegin   string
		wrapperEnd     string
		want           string
	}{
		{
			fieldFormat:    fmt.Sprintf("%s-%s-%s", FlagField, FlagType, FlagValue),
			fieldSeparator: ";",
			wrapperBegin:   "<",
			wrapperEnd:     ">",
			want:           "",
		},
		{
			fieldFormat:    fmt.Sprintf("%s>%s>%s", FlagField, FlagValue, FlagType),
			fieldSeparator: "|",
			wrapperBegin:   "(",
			wrapperEnd:     ")",
			want:           "",
		},
		{
			fieldFormat:    fmt.Sprintf("%s#%s#%s", FlagType, FlagField, FlagValue),
			fieldSeparator: " ",
			wrapperBegin:   "{",
			wrapperEnd:     "}",
			want:           "",
		},
	}

	for _, tt := range tests {
		got := CustomFormat(u, tt.fieldFormat, tt.fieldSeparator, tt.wrapperBegin, tt.wrapperEnd)
		if tt.want != got {
			t.Errorf("test failed, want: %s, but got: %s", tt.want, got)
		}
	}
}
