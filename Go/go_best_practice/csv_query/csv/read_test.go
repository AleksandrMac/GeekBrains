package csv_test

import (
	"testing"

	"github.com/AleksandrMac/GeekBrains/Go/go_best_practice/csv_query/csv"
	"github.com/stretchr/testify/assert"
)

var where string = "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND '2020-04-14' != date))"

type expesion struct {
	Left, Right string
}

func TestSplit(t *testing.T) {
	left, right := where, ""
	want := []expesion{
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND '2020-04-14' != date)", Right: ")"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND '2020-04-14' != date", Right: ")"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND '2020-04-14' != ", Right: "date"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND '2020-04-14' ", Right: "!="},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND ", Right: "'2020-04-14'"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' ", Right: "AND"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent=", Right: "'Africa'"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent", Right: "="},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (", Right: "continent"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR ", Right: "("},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') ", Right: "OR"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20'", Right: ")"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date < ", Right: "'2020-04-20'"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND date ", Right: "<"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' AND ", Right: "date"},
		{Left: "(continent='Asia' AND (date>'2020-04-14' ", Right: "AND"},
		{Left: "(continent='Asia' AND (date>", Right: "'2020-04-14'"},
		{Left: "(continent='Asia' AND (date", Right: ">"},
		{Left: "(continent='Asia' AND (", Right: "date"},
		{Left: "(continent='Asia' AND ", Right: "("},
		{Left: "(continent='Asia' ", Right: "AND"},
		{Left: "(continent=", Right: "'Asia'"},
		{Left: "(continent", Right: "="},
		{Left: "(", Right: "continent"},
		{Left: "", Right: "("},
	}
	for _, val := range want {
		left, right = csv.Split(left)
		got := expesion{Left: left, Right: right}
		assert.Equal(t, val, got, "they should be equal")
	}
}
