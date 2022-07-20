package service

import "fmt"

func HelloV4(name string) string {
	//fmt.Println(name)
	return name
}

func timesThree(number int) {
	fmt.Println(number)
}

type Piece struct {
	ID uint `json:"id"`
}
