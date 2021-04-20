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

func (d *Row) IsMatch(match string) bool {
	return true
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
