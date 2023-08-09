package main

import (
	"reflect"
	"testing"
)

func Test_padParts(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Simple", args{[]string{"Larry", "Curly", "Moe"}}, []string{"Larry", "Curly", "Moe  "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padParts(tt.args.parts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("padParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
