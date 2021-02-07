package main

func main() {
   // Declare a value of type int value 10
   count := 10
   
   // Display the value of and address of count
   println("count: \tValue Of[", count, "]\tAddr Of[", &count, "]")
   
   // Pass the value of the count
   increment(&count)
   
   println("count:\tValue of[", count, "]\tAddr of [",&count, "]")
}

func increment(inc *int) {
   *inc++
   println("inc:\tValue of[", inc, "]\tAddr Of[", &inc, "]")
}
