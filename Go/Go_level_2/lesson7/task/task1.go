package task

import (
	"fmt"
	"log"
	"reflect"
)

// Task1 - Написать функцию, которая принимает на вход структуру
// in (struct или кастомную struct) и
// values map[string]interface{} (key - название поля структуру, которому нужно присвоить value этой мапки).
// Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
// Функция может возвращать только ошибку error.
// Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).
func Task1(in interface{}, values map[string]interface{}) (err error) {
	if in == nil {
		return fmt.Errorf("\"in\" is nil")
	}

	val := reflect.ValueOf(&in)
	val = val.Elem()
	//num := val.NumField()
	//fmt.Println(num)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// if val.Kind() != reflect.Struct {
	// 	return fmt.Errorf("\"in\" is not struct")
	// }

	for key, v := range values {
		t := val.FieldByName(key)
		if t.IsZero() {
			return fmt.Errorf("field \"%v\" is not correct", key)
		}
		newVal := reflect.ValueOf(v)
		if t.Type().Kind() != newVal.Type().Kind() {
			return fmt.Errorf("type field \"%v\" is not correct, expected %v, curent: %v", key, t.Type().Kind().String(), newVal.Type().Kind().String())
		}
		fmt.Println(t)

		fmt.Println(newVal)
		tt := t.Elem()
		tt.Set(newVal)
		return
	}
	return fmt.Errorf("\"values\" is empty")
}

// PrintStruct -
func PrintStruct(in interface{}) {
	if in == nil {
		return
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		if typeField.Type.Kind() == reflect.Struct {
			log.Printf("nested field: %v", typeField.Name)
			PrintStruct(val.Field(i).Interface())
			continue
		}

		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			val.Field(i),
			typeField.Tag,
		)
	}
}
