package main

import (
	"fmt"
	"os"
)

func handleError(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err.Error())
	}
}

func handleErrorWithReturn(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err.Error())
		return
	}
}

func handleErrorWithExit(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err.Error())
		os.Exit(1)
	}
}
