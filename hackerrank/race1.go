package main

import (
   "fmt"
   "sync"
)

var counter int

func main() {
   const grs = 2
   // wg is used to manage concurrency
   var wg sync.WaitGroup
   wg.Add(grs)
   
   // Create two goroutines
   for g := 0; g < grs; g++ {
      go func() {
         for i := 0; i < 2; i++ {
            value := counter
            value++
            // fmt.Println(value)
            counter = value
         }
         wg.Done()
      }()
   }
   wg.Wait()
   fmt.Println("Final counter:", counter)
}
