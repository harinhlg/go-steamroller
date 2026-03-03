package main

import (
    "fmt"
    "github.com/google/uuid"
    "github.com/harinhlg/go-steamroller"
)

func main() {
    fmt.Println("JFrog Go Test:", uuid.New())

    result := steamroller.ExampleFunction() // replace with actual exported function
    fmt.Println("Steamroller result:", result)
}
