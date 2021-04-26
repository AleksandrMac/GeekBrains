package task

import (
	"fmt"
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

	if values == nil {
		return fmt.Errorf("\"values\" is empty")
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("\"in\" is not struct")
	}

	for key, v := range values {
		tmp := val.FieldByName(key)
		if tmp.IsZero() {
			err = fmt.Errorf("field \"%v\" is not correct\n%w", key, err)
			continue
		}
		newVal := reflect.ValueOf(v)
		if tmp.Type().Kind() != newVal.Type().Kind() {
			err = fmt.Errorf("type field \"%v\" is not correct, expected %v, current: %v\n%w", key, tmp.Type().Kind(), newVal.Type().Kind(), err)
			continue
		}
		tmp.Set(newVal)
	}
	return err
}
