package utils

import "fmt"

func PrintError(err error) {

	if err != nil {
		fmt.Println("---------Error message Start---------")
		fmt.Println(err)
		fmt.Println("---------Error message End---------")
	}
}
