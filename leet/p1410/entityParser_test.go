package main

import "testing"

func Test_entityParser(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"&amp; is an HTML entity but &ambassador; is not."}, "& is an HTML entity but &ambassador; is not."},
		{"test2", args{"and I quote: &quot;...&quot;"}, "and I quote: \"...\""},
		{"test3", args{"Stay home! Practice on Leetcode :)"}, "Stay home! Practice on Leetcode :)"},
		{"test4", args{"x &gt; y &amp;&amp; x &lt; y is always false"}, "x > y && x < y is always false"},
		{"test5", args{"leetcode.com&frasl;problemset&frasl;all"}, "leetcode.com/problemset/all"},
		{"test6", args{"&amp;gt;"}, "&gt;"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entityParser(tt.args.text); got != tt.want {
				t.Errorf("entityParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
