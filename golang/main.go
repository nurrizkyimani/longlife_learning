package main


// v3

import (
    "fmt"
    "github.com/nurrizkyimani/longlife_learning/golang/greetings"
    "github.com/nurrizkyimani/longlife_learning/golang/utils"
    "log"
  
)

func main() {

    log.setPrefix("greetings: ")
    log.setFlags(0)

  
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)

    n:= utils.TestingLen("Gladys")
    fmt.Println(n)

    msgV2, errHello := greetings.HelloV2("")

  
    if errHello != nil {
      fmt.Println("no File")
      log.Fatal(err)
    } else {
      fmt.Prinln("aman")
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

