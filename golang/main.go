package main

import "fmt"

func main() {
	fmt.Println("first trial")

}

// V3
//func main() {
//
//	log.SetPrefix("greetings: ")
//	log.SetFlags(0)
//
//	// Get a greeting message and print it.
//	message := greetings.Hello("Gladys")
//	fmt.Println(message)
//
//	n := utils.TestingLen("Gladys")
//	fmt.Println(n)
//
//	_, errHello := greetings.HelloV2("")
//
//	_, errHello2 := utils.HellosV3([]string{"Gladys", "Nur", "Imani", "Nurrizky"})
//
//	nameV4 := service.HelloV4("imani_V4")
//
//	fmt.Println(nameV4)
//
//	if errHello != nil {
//		fmt.Println("no File")
//		log.Fatal(errHello)
//	} else {
//		fmt.Println("aman")
//	}
//
//	if errHello2 != nil {
//		fmt.Println("no File")
//		log.Fatal(errHello2)
//	} else {
//		fmt.Println("aman")
//	}
//}

// === V1
// func threeTimes(number int, ch chan int){
//   result := number * 3
//   fmt.Println(result)
//   ch <- result
// }

// func main(){
//   fmt.Println("Executing goroutine")
//   ch := make(chan int)
//   go threeTimes(3, ch)
//   res := <- ch

//   fmt.Println("The result is:%v", res)
//   fmt.Println("Finishing")

// === V2
//   fmt.Println(utils.testingLen("imani"))
// }

// func threeTimes_v2(number int){
//   result := number * 3
//   fmt.Println(result)
// }

// func main(){
//   fmt.Println("Executing goroutine")
//   go threeTimes_v2(3)
//   // fmt.Println("The result is:%v", res)
//   fmt.Println("Finishing")
// }
