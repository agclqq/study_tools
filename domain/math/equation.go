package math

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

type Problem struct {
	Stem    string   // 题目
	Options []string // 选项
	Answer  string   // 答案
}

// GenerateProblems 生成算述题目
// n:生成题目数量
// min:题目中计算的最小值
// max:题目中计算的最大值
// x:几个数的运算
// parenthesis: 是否带括号
// ops:运行符号
func GenerateProblems(n, min, max, x int, parenthesis bool, ops []string) []Problem {
	problems := make([]Problem, 0, n)
	seen := make(map[string]bool)

	for len(problems) < n {
		expr, answer, valid := generateOperation(min, max, x, parenthesis, ops)
		if !valid || seen[expr] {
			continue
		}
		expr = strings.ReplaceAll(expr, " ", "") + "="
		options := generateOptions(answer, min, max)
		problems = append(problems, Problem{
			Stem:    expr,
			Answer:  answer,
			Options: options,
		})
		seen[expr] = true
	}

	return problems
}

func generateOperation(min, max, x int, parenthesis bool, ops []string) (string, string, bool) {
	// 生成 x 个随机数和 x-1 个操作符
	numbers := make([]int, x)
	operators := make([]string, x-1)

	for i := 0; i < x; i++ {
		numbers[i] = rand.IntN(max-min+1) + min
	}
	for i := 0; i < x-1; i++ {
		operators[i] = ops[rand.IntN(len(ops))]
	}

	// 构造表达式
	expr := buildExpression(min, max, numbers, operators, parenthesis)

	// 转换为后缀表达式并计算结果
	rpn := toRPN(expr)
	result, valid := evaluateRPNWithRange(rpn, min, max)
	if !valid {
		return "", "", false
	}

	return expr, strconv.Itoa(result), true
}

func buildExpression(min, max int, numbers []int, operators []string, parenthesis bool) string {
	// 生成表达式
	var expr strings.Builder
	if parenthesis && len(numbers) > 2 {
		// 随机决定是否添加括号
		start := rand.IntN(len(numbers) - 1)
		end := rand.IntN(len(numbers)-start-1) + start + 1

		// 构造表达式，添加括号
		for i := 0; i < len(numbers); i++ {
			if i == start {
				expr.WriteString(" ( ")
			}
			expr.WriteString(strconv.Itoa(numbers[i]))
			if i == end {
				expr.WriteString(" ) ")
			}
			if i < len(operators) {
				expr.WriteString(" " + operators[i] + " ")
			}
		}

		// 如果加了括号不会改变优先级，则去掉括号
		original := expr.String()
		if !shouldKeepParentheses(min, max, original) {
			// 重新生成没有括号的表达式
			expr.Reset()
			for i := 0; i < len(numbers); i++ {
				expr.WriteString(strconv.Itoa(numbers[i]))
				if i < len(operators) {
					expr.WriteString(" " + operators[i] + " ")
				}
			}
		}

	} else {
		// 无括号情况
		for i := 0; i < len(numbers); i++ {
			expr.WriteString(strconv.Itoa(numbers[i]))
			if i < len(operators) {
				expr.WriteString(" " + operators[i] + " ")
			}
		}
	}

	return expr.String()
}

// 判断括号是否改变了优先级
func shouldKeepParentheses(min, max int, expr string) bool {
	// 移除括号并计算结果
	exprWithoutParentheses := removeParentheses(expr)

	// 计算带括号和不带括号的表达式结果
	rpn := toRPN(expr)
	originalResult, valid := evaluateRPNWithRange(rpn, min, max) // 可以适当扩展范围
	if !valid {
		return false
	}

	rpnWithout := toRPN(exprWithoutParentheses)
	resultWithout, valid := evaluateRPNWithRange(rpnWithout, min, max)
	if !valid {
		return false
	}

	// 如果加括号前后结果不同，保留括号，否则去掉括号
	return originalResult != resultWithout
}

// 移除表达式中的括号
func removeParentheses(expr string) string {
	// 移除括号
	expr = strings.ReplaceAll(expr, " ( ", "")
	expr = strings.ReplaceAll(expr, " ) ", "")
	return expr
}

func toRPN(expr string) []string {
	tokens := strings.Fields(expr)
	var output []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"(": 0,
	}

	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // 弹出左括号
		} else {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output
}

func evaluateRPNWithRange(rpn []string, min, max int) (int, bool) {
	var stack []int

	for _, token := range rpn {
		if isNumber(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, false
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 || a%b != 0 {
					return 0, false
				}
				result = a / b
			}

			if result < min || result > max {
				return 0, false
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, false
	}

	return stack[0], true
}

func generateOptions(answer string, min, max int) []string {
	options := make(map[string]struct{})
	options[answer] = struct{}{}

	for len(options) < 4 {
		incorrect := strconv.Itoa(rand.IntN(max-min+1) + min)
		options[incorrect] = struct{}{}
	}

	var result []string
	for opt := range options {
		result = append(result, opt)
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
