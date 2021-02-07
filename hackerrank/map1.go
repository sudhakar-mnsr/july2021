package main

import "fmt"

type user struct {
   name string
   surname string
}

func main() {
users := make(map[string]user)

// Add key/value pairs to the map
users["Roy"] = user{"Rob", "Roy"}
users["Ford"] = user{"Henry", "Ford"}
users["Roy"] = user{"Mickey", "Mouse"}
users["Roy"] = user{"Michael", "Jackson"}

mouse := users["Mouse"]
fmt.Printf("%+v\n", mouse)

users["Mouse"] = user{"Jerry", "Mouse"}
fmt.Printf("%+v\n", users["Mouse"])

fmt.Println(len(users))
delete(users, "Roy")
fmt.Println(len(users))

fmt.Println("Goodbye.")
}

