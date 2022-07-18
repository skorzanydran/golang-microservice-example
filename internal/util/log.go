package util

import "fmt"

func Log(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
	}
}
