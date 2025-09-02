package main

import (
	"daKit/modSystem"
	"fmt"
)

func main() {
	data1, err := modSystem.Detect()
	if err != nil {
		fmt.Println("detect system error: ", err)
	} else {
		fmt.Println(data1)
	}

	//serviceStart()
	//0925419460
}
