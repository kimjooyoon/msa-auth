package util

import (
	"reflect"
	"testing"
)

func Test_makeSnakeCase(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success:f()==''", args{}, ""},
		{"success:f('')==''", args{[]string{""}}, ""},
		{"success:f('test')=='test'", args{[]string{"test"}}, "test"},
		{"success:f('abc','def')=='abc_def'", args{[]string{"abc", "def"}}, "abc_def"},
		{"success:f('abc','def','test')=='abc_def_test'", args{[]string{"abc", "def", "test"}}, "abc_def_test"},
		{"success:f('','')=='_'", args{[]string{"", ""}}, "_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSnakeCase(tt.args.str...); got != tt.want {
				t.Errorf("makeSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeStringList(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"success:f('a')=={'a'}", args{[]string{"a"}}, []string{"a"}},
		{"success:f('a','b')=={'a','b'}", args{[]string{"a", "b"}}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStringList(tt.args.str...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeStringList() = %v, want %v", got, tt.want)
			}
		})
	}
}
