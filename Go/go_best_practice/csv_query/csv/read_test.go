package csv_test

import (
	"testing"

	"github.com/AleksandrMac/GeekBrains/Go/go_best_practice/csv_query/csv"
	"github.com/stretchr/testify/assert"
)

var where string = "(continent='Asia' AND (date>'2020-04-14' AND date < '2020-04-20') OR (continent='Africa' AND '2020-04-14' != date))"

type splitData struct {
	Left, Right string
}

type getItemData struct {
	Operation, Values []string
}

func TestSplitReverse(t *testing.T) {
	left, right := where, ""
	want := []splitData{
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
		left, right = csv.SplitReverse(left)
		got := splitData{Left: left, Right: right}
		assert.Equal(t, val, got, "they should be equal")
	}
}

func TestGetItem(t *testing.T) {
	want := getItemData{
		Operation: []string{")", ")", "!=", "AND", "=", "(", "OR", ")", "<", "AND", ">", "(", "AND", "=", "("},
		Values:    []string{"'2020-04-14'", "'Africa'", "continent", "'2020-04-20'", "date", "'2020-04-14'", "date", "'Asia'", "continent"},
	}
	got := getItemData{}
	got.Operation, got.Values = csv.GetItem(where)
	assert.Equal(t, want, got, "func \"TestGetItem\": they should be equal")
}

func TestGetLex(t *testing.T) {
	want := []string{"(", "continent", "=", "'Asia'", "AND", "(", "date", ">", "'2020-04-14'", "AND", "date", "<", "'2020-04-20'", ")", "OR", "(", "continent", "=", "'Africa'", "AND", "'2020-04-14'", "!=", "date", ")", ")"}
	got := csv.GetLex(where)
	assert.Equal(t, want, got, "they should be equal")
}
func TestInfixToPostfix(t *testing.T) {
	want := []string{"continent", "'Asia'", "=", "date", "'2020-04-14'", ">", "date", "'2020-04-20'", "<", "AND", "continent", "'Africa'", "=", "'2020-04-14'", "date", "!=", "AND", "OR", "AND"}
	got := csv.InfixToPostfix(csv.GetLex(where))
	assert.Equal(t, want, got, "func \"TestGetItem\": they should be equal")
}
