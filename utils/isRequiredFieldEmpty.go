package utils

import "fmt"

func IsRequiredFieldEmpty(field interface{}, fieldName string) bool {
	if field == nil || field == "" {
		fmt.Printf("%s is a required field\n", fieldName)
		return true
	}
	return false
}
