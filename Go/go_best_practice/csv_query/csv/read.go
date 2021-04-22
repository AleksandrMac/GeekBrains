// csv
package csv

import "strings"

type Head struct {
	Path   string
	Fields []string
}

type Body struct {
	*Head
	Rows []Row
}

type Row struct {
	*Head
	Values []string
}

// type operation string

// const (
// 	less    operation = "<"
// 	great   operation = ">"
// 	lessEq  operation = "<="
// 	greatEq operation = ">="
// 	equal   operation = "="
// 	notEq   operation = "!="
// 	open    operation = "("
// 	close   operation = ")"
// )

// var (
// 	op  []string
// 	val []string
// )

func (d *Row) IsMatch(match string) bool {
	match = strings.TrimSpace(match)
	match = strings.ToUpper(match)

	lexInfix := GetLex(match)
	InfixToPostfix(lexInfix)

	// for lop, lval := len(op), len(val); lop > 0 && lval > 0; {
	// 	// o := op[lop-1]
	// 	// leftExp, rightExp := "", ""

	// }
	return true
}

func InfixToPostfix(infix []string) (postfix []string) {
	stack := make([]string, 0, len(infix))
	for _, val := range infix {
		switch val {
		case "(":
			stack = append(stack, val)
		case ")":
			i := len(stack) - 1
			for ; stack[i] != "("; i-- {
				postfix = append(postfix, stack[i])
			}
			stack = stack[:i]
		case "+", "-", "!", "NOT", "/", "*", "DIV", "MOD", "AND", "OR", "<", ">", "<=", ">=", "<>", "!=", "=":
			//i := len(stack) - 1
			for i := len(stack) - 1; GetPriority(stack[i]) >= GetPriority(val) && i >= 0; i = len(stack) - 1 {
				postfix = append(postfix, stack[i])
				stack = stack[:i]
			}
			stack = append(stack, val)
		default:
			postfix = append(postfix, val)
		}
	}
	return postfix
}

func GetPriority(operator string) uint8 {
	operator = strings.ToUpper(operator)
	switch operator {
	case "(", ")":
		return 1
	case "+", "-", "!", "NOT":
		return 2
	case "/", "*", "DIV", "MOD", "AND":
		return 3
	case "OR":
		return 4
	case "<", ">", "<=", ">=", "<>", "!=", "=":
		return 5
	default:
		return 0
	}
}
func GetLex(str string) (lex []string) {
	for left, right := "", str; len(right) > 0; {
		left, right = Split(right)
		lex = append(lex, left)
	}
	return lex
}
func GetItem(str string) (op, val []string) {
	for left, right := str, ""; len(left) > 0; {
		left, right = Split(left)
		switch right[0] {
		case '<', '>', '=', '!', '(', ')':
			op = append(op, right)
		default:
			if right == "AND" || right == "OR" {
				op = append(op, right)
				continue
			}
			val = append(val, right)
		}
	}
	return op, val
}

func Split(str string) (left, right string) {
	inputBuffer := []byte(str)
	outputBuffer := make([]byte, 0, len(inputBuffer))
	spaceIgnoring := false
	for i, val := range inputBuffer {
		switch val {
		case ')', '(':
			if len(outputBuffer) > 0 {
				return string(outputBuffer), string(inputBuffer[i:])
			}
			return string(val), string(inputBuffer[i+1:])
		case '>', '<', '!', '=':
			if len(outputBuffer) > 0 {
				return string(outputBuffer), string(inputBuffer[i:])
			}
			outputBuffer = append(outputBuffer, val)
			switch inputBuffer[i+1] {
			case '=':
				outputBuffer = append(outputBuffer, inputBuffer[i+1])
				return string(outputBuffer), string(inputBuffer[i+2:])
			default:
				return string(outputBuffer), string(inputBuffer[i+1:])
			}
		default:
			switch val {
			case ' ':
				if !spaceIgnoring {
					if len(outputBuffer) > 0 {
						return string(outputBuffer), string(inputBuffer[i+1:])
					}
					continue
				}
				outputBuffer = append(outputBuffer, ' ')
			case []byte("'")[0]:
				if !spaceIgnoring {
					spaceIgnoring = true
					outputBuffer = append(outputBuffer, val)
					continue
				}
				outputBuffer = append(outputBuffer, val)
				return string(outputBuffer), string(inputBuffer[i+1:])
			default:
				outputBuffer = append(outputBuffer, val)
			}
		}
	}
	return string(outputBuffer), string(inputBuffer)
}

func SplitReverse(str string) (left, right string) {
	inputBuffer := []byte(str)
	outputBuffer := make([]byte, 0, len(inputBuffer))
	spaceIgnoring := false
	for i := len(str); i > 0; i-- {
		val := inputBuffer[i-1]
		switch val {
		case ')', '(':
			if len(outputBuffer) > 0 {
				i++
				return string(inputBuffer[:i-1]), string(reverse(outputBuffer))
			}
			return string(inputBuffer[:i-1]), string(val)
		case '>', '<', '=':
			outputBuffer = append(outputBuffer, val)
			val2 := inputBuffer[i-2]
			switch val2 {
			case '<', '>', '!':
				outputBuffer = append(outputBuffer, val2)
				return string(inputBuffer[:i-2]), string(reverse(outputBuffer))
			default:
				return string(inputBuffer[:i-1]), string(outputBuffer)
			}
		default:
			switch val {
			case ' ':
				if !spaceIgnoring {
					if len(outputBuffer) > 0 {
						return string(inputBuffer[:i]), string(reverse(outputBuffer))
					}
					continue
				}
				outputBuffer = append(outputBuffer, ' ')
			case []byte("'")[0]:
				if !spaceIgnoring {
					spaceIgnoring = true
					outputBuffer = append(outputBuffer, val)
					continue
				}
				outputBuffer = append(outputBuffer, val)
				return string(inputBuffer[:i-1]), string(reverse(outputBuffer))
			default:
				outputBuffer = append(outputBuffer, val)
			}
		}
	}
	return string(inputBuffer), string(reverse(outputBuffer))
}

func reverse(str []byte) []byte {
	newStr := make([]byte, 0, len(str))
	for i := len(str); i > 0; i-- {
		newStr = append(newStr, str[i-1])
	}
	return newStr
}

func GetFields(row, sep string) []string {
	if sep == "" {
		sep = ","
	}
	return strings.Split(row, sep)
}

func (h *Head) NewRow() *Row {
	return &Row{Head: h}
}

// var stack []string

// stack = append(stack, "world!") // Push
// stack = append(stack, "Hello ")

// for len(stack) > 0 {
//     n := len(stack) - 1 // Top element
//     fmt.Print(stack[n])

//     stack = stack[:n] // Pop
// }
