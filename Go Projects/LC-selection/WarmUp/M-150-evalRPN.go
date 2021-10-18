package WarmUp

import "strconv"

// 逆波兰表达式
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch tokens[i] {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			stack = append(stack, res)
		default:
			// 非表达式则将 字符串转化为数字，追加到栈顶
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	return stack[0]
}
