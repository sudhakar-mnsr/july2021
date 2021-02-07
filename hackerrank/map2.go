package main

import "fmt"

func main() {
// Create a map to track scores for players in a game
scores := make(map[string]int)

// Read the element at key "anna". It is absent so we get
// the zero-value for this maps value type.
score := scores["anna"]

fmt.Println("Score:", score)

// If we need to check for the presence of a key we use
// a 2 variable assignment. The 2nd variable is bool.

score, ok := scores["anna"]
fmt.Println("Score:", score, "Present:", ok)

scores["anna"]++
// Equivalent to 
if n, ok := scores["anna"]; ok {
   scores["anna"] = n + 1
} else {
   scores["anna"] = 1
}

score, ok = scores["anna"]
fmt.Println("Score:", score, "Present:", ok)

}
