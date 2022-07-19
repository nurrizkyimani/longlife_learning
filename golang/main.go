package main

// v3

import (
	"fmt"
	"log"

	"github.com/nurrizkyimani/longlife_learning/golang/greetings"
	"github.com/nurrizkyimani/longlife_learning/golang/utils"
)

func main() {

    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)

    n:= utils.TestingLen("Gladys")
    fmt.Println(n)

    _ , errHello := greetings.HelloV2("")

    _ , errHello2 := utils.HellosV3([]string{"Gladys", "Nur", "Imani", "Nurrizky"})
  
    if errHello != nil {
      fmt.Println("no File")
      log.Fatal(errHello)
    } else {
      fmt.Println("aman")
    }

    if errHello2 != nil {
      fmt.Println("no File")
      log.Fatal(errHello2)
    } else {
      fmt.Println("aman")
    }


  

}

// v2
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



// v2  
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

