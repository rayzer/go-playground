package main

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test 1", args{[]int{1, 3}, 3}, []string{"Push", "Push", "Pop", "Push"}},
		{"test 2", args{[]int{1, 2, 3}, 3}, []string{"Push", "Push", "Push"}},
		{"test 3", args{[]int{1, 2}, 4}, []string{"Push", "Push"}},
		{"test 4", args{[]int{2, 3, 4}, 4}, []string{"Push", "Pop", "Push", "Push", "Push"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.target, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
