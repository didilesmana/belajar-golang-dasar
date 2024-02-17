package main

import (
	"fmt"
)

type validationErrr struct {
	Message string
}

func (v *validationErrr) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SavedData(id string, data any) error {
	if id == "" {
		return &validationErrr{"valid error"}
	}

	if id != "didi" {
		return &notFoundError{"not found"}
	}

	return nil
}

func main() {
	err := SavedData("eko", nil)
	if err != nil {
		if validationErrr, ok := err.(*validationErrr); ok {
			fmt.Println("valid error", validationErrr.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found", notFoundError.Error())
		} else {
			fmt.Println("unknown", err.Error())
		}
	} else {
		fmt.Println("susses")
	}
}
