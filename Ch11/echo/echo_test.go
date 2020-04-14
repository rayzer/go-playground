package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_echo(t *testing.T) {
	type args struct {
		newline bool
		sep     string
		args    []string
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{"test1", args{true, "", []string{}}, "\n"},
		{"test2", args{false, "", []string{}}, ""},
		{"test3", args{true, "\t", []string{"one", "two", "three"}}, "one\ttwo\tthree\n"},
		{"test4", args{true, ",", []string{"a", "b", "c"}}, "a,b,c\n"},
		{"test5", args{false, ":", []string{"1", "2", "3"}}, "1:2:3"},
		{"testError", args{true, ",", []string{"a", "b", "c"}}, "a b c\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			descr := fmt.Sprintf("echo(%v, %q, %q", tt.args.newline, tt.args.sep, tt.args.args)
			out = new(bytes.Buffer)
			if err := echo(tt.args.newline, tt.args.sep, tt.args.args); err != nil {
				t.Errorf("%s failed:; %v", descr, err)
				return
			}
			got := out.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("%s = %q, want %q", descr, got, tt.want)
			}

		})
	}
}
