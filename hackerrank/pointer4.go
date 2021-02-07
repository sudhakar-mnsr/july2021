package main

// Number of elements to grow each stack frame
// Run with 1 and then with 1024

// const size = 1
const size = 1024

// main is the entrypoint for the application
func main() {
   s := "Hello"
   stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
   println(c, s, *s)
   
   c++
   
   if c == 10 {
      return
   }
   
   stackCopy(s, c, a)
}
