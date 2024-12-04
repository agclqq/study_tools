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
	rs := math.GenerateProblems(10, 0, 20, 2, false, []string{"+", "-", "*"})
	rs = append(rs, math.GenerateProblems(10, 0, 50, 3, false, []string{"+", "-", "*"})...)
	rs = append(rs, math.GenerateProblems(10, 0, 50, 2, false, []string{"+", "-", "*"})...)
	rs = append(rs, math.GenerateProblems(20, 0, 100, 2, false, []string{"+", "-", "*"})...)
	rs = append(rs, math.GenerateProblems(10, 1000, 10000, 2, false, []string{"+", "-", "*"})...)

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
	//20以内加减乘法
	rs := math.GenerateProblems(10, 0, 20, 3, false, []string{"+", "-", "*"})
	rs = append(rs, math.GenerateProblems(30, 0, 50, 3, false, []string{"+", "-"})...)
	rs = append(rs, math.GenerateProblems(10, 0, 100, 3, false, []string{"+", "-"})...)

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
	//20以内加减乘法
	rs := math.GenerateProblems(10, 0, 100, 2, false, []string{"+", "-", "*"})
	rs = append(rs, math.GenerateProblems(20, 100, 1000, 2, false, []string{"+", "-", "*"})...)

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
