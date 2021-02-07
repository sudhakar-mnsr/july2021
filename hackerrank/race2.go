package main

import (
   "fmt"
   "sync"
   "sync/atomic"
)

var counter int64

func main() {
   // Number of goroutines to use
   const grs = 2
   
   // wg is used to manage concurrency
   var wg sync.WaitGroup
   wg.Add(grs)
   
   // create two goroutines
   for g := 0; g < grs; g++ {
      go func() {
         for i := 0; i < 2; i++ {
            atomic.AddInt64(&counter, 1)
         }
         wg.Done()
      }()
   }
   
   wg.Wait()
   
   fmt.Println("Final Counter:", counter)
}
