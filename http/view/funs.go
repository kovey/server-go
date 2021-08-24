package view

import "fmt"

func Add(left int, right int) int {
	return left + right
}

func Sub(left int, right int) int {
	return left - right
}

func Mul(left int, right int) int {
	return left * right
}

func Div(left int, right int) int {
	return left / right
}

func MapGet(data map[string]interface{}, key string) (interface{}, error) {
	val, ok := data[key]
	if !ok {
		return nil, fmt.Errorf("key[%s] is not exists", key)
	}

	return val, nil
}
