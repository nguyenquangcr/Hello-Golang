package utils

import "fmt"

func IsRequiredFieldEmpty(field interface{}, fieldName string) bool {
	if field == nil || field == "" {
		fmt.Printf("%s is a required field\n", fieldName)
		return true
	}
	return false
}

func AreRequiredFieldsEmpty(fields ...interface{}) bool {
	for _, field := range fields {
		if field == nil || field == "" {
			fmt.Println("A required field is empty")
			return true
		}
	}
	return false
}
