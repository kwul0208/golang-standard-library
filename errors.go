package main

import "errors"

var (
	validationError = errors.New("Error Validation")
	notFounfError   = errors.New("Error Not Found")
)

func GetById(id string) error {
	if id == "" {
		return validationError
	} else if id == "ahmud" {
		return notFounfError
	}

	return nil
}

func main() {
	err := GetById("")

	// if err == validationError {
	// 	print(validationError.Error())
	// } else if err == notFounfError {
	// 	print("not found error")
	// } else {
	// 	print("ok")
	// }
	print(err.Error())
}
