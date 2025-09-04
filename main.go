package main

import (
	"daKit/modSystem"
	"daKit/modSystemService"
	"fmt"
)

func main() {
	data1, err := modSystem.Detect()
	if err != nil {
		fmt.Println("detect system error: ", err)
	} else {
		fmt.Println(data1)
	}

	modSystemService.SS_ServiceStart()
}
