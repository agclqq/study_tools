package command

import (
	"fmt"
	"math/rand/v2"

	"github.com/agclqq/prowjob"

	"github.com/agclqq/study_tools/domain/math"
)

type PrimaryThreeMath struct {
}

func (p *PrimaryThreeMath) GetCommand() string {
	return "command:primaryThreeMath"
}
func (p *PrimaryThreeMath) Usage() string {
	return ``
}
func (p *PrimaryThreeMath) Handle(ctx *prowjob.Context) {
	mentalMath()
	fmt.Println()
	detachedCalculation()
	fmt.Println()
	columnarCalculation()
}

// 口算题
func mentalMath() {
	fmt.Println("口算题")
	num := 30
	//20以内加减乘法
	rs := math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 81, Nums: []math.Num{{Min: 0, Max: 9}, {Min: 0, Max: 9}}, Parenthesis: false, Ops: []string{"+", "-", "*"}})
	//100以后加减法
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 100, Nums: []math.Num{{Min: 0, Max: 100}, {Min: 0, Max: 100}}, Parenthesis: false, Ops: []string{"+", "-"}})...)
	//1000以内2位数乘1位数
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 1000, Nums: []math.Num{{Min: 1, Max: 100}, {Min: 0, Max: 10}}, Parenthesis: false, Ops: []string{"*"}})...)

	rand.Shuffle(len(rs), func(i, j int) {
		rs[i], rs[j] = rs[j], rs[i]
	})
	for i, v := range rs {
		if i >= num {
			break
		}
		fmt.Printf("%s\t\t\t", v.Stem)
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
}

// 脱式计算
func detachedCalculation() {
	fmt.Println("脱式计算")
	num := 10
	//100以内三数加减法
	rs := math.GenerateProblems(&math.ProblemParam{N: 10, Min: 0, Max: 81, Nums: []math.Num{{Min: 0, Max: 100}, {Min: 0, Max: 100}, {Min: 0, Max: 100}}, Parenthesis: false, Ops: []string{"+", "-"}})
	//100以内三数加减乘法
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 200, Nums: []math.Num{{Min: 10, Max: 100}, {Min: 0, Max: 100}, {Min: 0, Max: 10}}, Parenthesis: false, Ops: []string{"+", "-"}})...)
	//10000以内三数加减乘法
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 1000, Nums: []math.Num{{Min: 10, Max: 100}, {Min: 100, Max: 1000}, {Min: 0, Max: 100}}, Parenthesis: false, Ops: []string{"+", "-"}})...)
	//万以内3位数乘1位数
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 10000, Nums: []math.Num{{Min: 100, Max: 999}, {Min: 0, Max: 10}}, Parenthesis: false, Ops: []string{"*"}})...)

	rand.Shuffle(len(rs), func(i, j int) {
		rs[i], rs[j] = rs[j], rs[i]
	})

	for i, v := range rs {
		if i >= num {
			break
		}
		fmt.Printf("%s\t\t\t", v.Stem)
		if (i+1)%5 == 0 {
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}
}

func columnarCalculation() {
	fmt.Println("竖式计算")
	num := 10

	//100以内三数加减法
	rs := math.GenerateProblems(&math.ProblemParam{N: 10, Min: 0, Max: 100, Nums: []math.Num{{Min: 0, Max: 100}, {Min: 0, Max: 100}, {Min: 0, Max: 100}}, Parenthesis: false, Ops: []string{"+", "-", "*"}})
	//1000以内三数加减乘法
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 1000, Nums: []math.Num{{Min: 10, Max: 100}, {Min: 0, Max: 100}, {Min: 0, Max: 10}}, Parenthesis: false, Ops: []string{"+", "-", "*"}})...)
	//10000以内三数加减乘法
	rs = append(rs, math.GenerateProblems(&math.ProblemParam{N: 30, Min: 0, Max: 10000, Nums: []math.Num{{Min: 10, Max: 100}, {Min: 100, Max: 1000}, {Min: 0, Max: 100}}, Parenthesis: false, Ops: []string{"+", "-"}})...)

	rand.Shuffle(len(rs), func(i, j int) {
		rs[i], rs[j] = rs[j], rs[i]
	})
	for i, v := range rs {
		if i >= num {
			break
		}
		fmt.Printf("%s\t\t", v.Stem)
		if (i+1)%5 == 0 {
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}
}
