package service

import "fmt"

func minusThree(number int, ch chan int) {
	ch <- number - 3
	fmt.Println("Function continue, after send the result ")
}

func TimesThreeV3(arr []int, ch chan int) {
	minusCh := make(chan int, 3)
	for _, elem := range arr {
		value := elem * 3
		if value%2 == 0 {
			go minusThree(value, minusCh)
			value = <-minusCh
		}
		ch <- value
	}
}

func HelloV4(name string) string {
	//fmt.Println(name)
	return name
}

func TimeThreeV2(arr []int, ch chan int) {
	for _, elm := range arr {
		ch <- elm * 3
	}
}

func TimesThree(number int, ch chan int) {
	//fmt.Println(number)
	result := number * 3
	fmt.Println("run chan")
	ch <- result
}

func TimesTwo(number int, ch chan int) {
	result := number * 2
	fmt.Println("run chan times Ttwo")
	ch <- result
}

type Piece struct {
	ID uint `json:"id"`
}
