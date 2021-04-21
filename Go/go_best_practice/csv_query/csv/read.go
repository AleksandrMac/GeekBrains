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

type operation string

const (
	less    operation = "<"
	great   operation = ">"
	lessEq  operation = "<="
	greatEq operation = ">="
	equal   operation = "="
	notEq   operation = "!="
	open    operation = "("
	close   operation = ")"
)

var (
	op  []operation
	val []string
)

func (d *Row) IsMatch(match string) bool {
	match = strings.TrimSpace(match)
	match = strings.ToUpper(match)
	return true
}

func Split(str string) (left, right string) {
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
