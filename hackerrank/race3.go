package main

import ( 
   "fmt"
   "math/rand"
   "sync"
   "sync/atomic"
   "time"
)

var data []string

var rwMutex sync.RWMutex

var readCount int64

func init() {
   rand.Seed(time.Now().UnixNano())
}

func main() {
   var wg sync.WaitGroup
   wg.Add(1)
   
   go func() {
      for i := 0; i < 10; i++ {
         writer(i)
      }
      wg.Done()
   }()
   
   // Create eight reader goroutines
   for i := 0; i < 8; i++ {
      go func(id int) {
         for {
            reader(id)
         }
      }(i)
   }
   wg.Wait()
   fmt.Println("Program complete")
}

func writer(i int) {
   rwMutex.Lock()
   {
      rc := atomic.LoadInt64(&readCount)
      time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
      fmt.Printf("****> : Performing Write: RCount[%d]\n", rc)
      data = append(data, fmt.Sprintf("String: %d", i))
   }
   rwMutex.Unlock()
}

func reader(id int) {
   rwMutex.RLock()
   {
      rc := atomic.AddInt64(&readCount, 1)
      time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
      fmt.Printf("%d : Performing read: Length[%d] RCount[%d]\n", id, len(data), rc)
      atomic.AddInt64(&readCount, -1)
   }
   rwMutex.RUnlock()
}
