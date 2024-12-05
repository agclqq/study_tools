package math

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestGenerateProblems(t *testing.T) {
	type args struct {
		pp *ProblemParam
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{name: "n1", args: args{&ProblemParam{N: 10, Min: 0, Max: 1000, Nums: []Num{{Min: 0, Max: 100}, {Min: 1, Max: 10}}, Parenthesis: false, Ops: []string{"*"}}}, wantNum: 10},
		{name: "n2", args: args{&ProblemParam{N: 10, Min: 0, Max: 500, Nums: []Num{{Min: 0, Max: 100}, {Min: 10, Max: 50}, {Min: 0, Max: 10}}, Parenthesis: true, Ops: []string{"+", "-", "*"}}}, wantNum: 10},
		{name: "n3", args: args{&ProblemParam{N: 10, Min: 0, Max: 300, Nums: []Num{{Min: 0, Max: 100}, {Min: 10, Max: 100}, {Min: 0, Max: 100}}, Parenthesis: true, Ops: []string{"+", "-"}}}, wantNum: 10},
		{name: "n4", args: args{&ProblemParam{N: 10, Min: 0, Max: 100, Nums: []Num{{Min: 0, Max: 1000}, {Min: 10, Max: 1000}}, Parenthesis: false, Ops: []string{"+", "-"}}}, wantNum: 10},
		{name: "n2", args: args{&ProblemParam{N: 10, Min: 0, Max: 100, Nums: []Num{{Min: 0, Max: 50}, {Min: 10, Max: 50}}, Parenthesis: false, Ops: []string{"+", "-"}}}, wantNum: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateProblems(tt.args.pp)
			if len(got) != tt.wantNum {
				t.Errorf("want num %d,but got %d", tt.wantNum, len(got))
				return
			}
			for _, v := range got {
				if cast.ToInt(v.Answer) > tt.args.pp.Max {
					t.Errorf("answer greater than max,%v", v)
					return
				}
				if cast.ToInt(v.Answer) < tt.args.pp.Min {
					t.Errorf("answer less than max,%v", v)
					return
				}
				fmt.Printf("\n题目：%s\n选项：%s\n答案：%s\n", v.Stem, v.Options, v.Answer)
			}
		})
	}
}
