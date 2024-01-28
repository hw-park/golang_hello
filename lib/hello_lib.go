package lib

import "fmt"

func GetHelloString(id string) string {
	value := fmt.Sprintf("Hello, %v!", id)
	return value
}
