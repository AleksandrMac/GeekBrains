package main

import (
	"fmt"

	"github.com/AleksandrMac/GeekBrains/Go/Go_level_2/lesson7/task"
)

func main() {
	st1 := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
		Slice       []int
		Object      struct {
			NestedField int
		}
	}{
		FieldString: "stroka",
		FieldInt:    107,
		Slice:       []int{112, 107, 207},
		Object:      struct{ NestedField int }{NestedField: 302},
	}

	v := map[string]interface{}{
		"FieldString": "NewString",
		"FieldInt":    "555",
		"Slice":       []int{1, 2, 3},
		"Object":      struct{ NestedField int }{NestedField: 777},
	}

	err := task.Task1(&st1, v)
	fmt.Println(err)
	fmt.Println(task.Task2("./task/task1.go", "Println"))
}
