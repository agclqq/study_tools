package math

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestGenerateProblems(t *testing.T) {
	type args struct {
		n         int
		min       int
		max       int
		x         int
		p         bool
		operators []string
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{name: "n1", args: args{n: 50, min: 0, max: 10, x: 2, p: true, operators: []string{"+", "-", "*"}}, wantNum: 50},
		{name: "n1", args: args{n: 50, min: 0, max: 20, x: 2, p: true, operators: []string{"+", "-", "*"}}, wantNum: 50},
		{name: "n1", args: args{n: 50, min: 0, max: 50, x: 3, p: true, operators: []string{"+", "-", "*"}}, wantNum: 50},
		{name: "n2", args: args{n: 10, min: 0, max: 50, x: 3, p: false, operators: []string{"+", "-", "*"}}, wantNum: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateProblems(tt.args.n, tt.args.min, tt.args.max, tt.args.x, tt.args.p, tt.args.operators)
			if len(got) != tt.wantNum {
				t.Errorf("want num %d,but got %d", tt.wantNum, len(got))
				return
			}
			for _, v := range got {
				if cast.ToInt(v.Answer) > tt.args.max {
					t.Errorf("answer greater than max,%v", v)
					return
				}
				if cast.ToInt(v.Answer) < tt.args.min {
					t.Errorf("answer less than max,%v", v)
					return
				}
				fmt.Printf("\n题目：%s\n选项：%s\n答案：%s\n", v.Stem, v.Options, v.Answer)
			}
		})
	}
}
