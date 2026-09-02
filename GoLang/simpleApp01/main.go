package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello from Go!")
    fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC3339))

    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, n := range numbers {
        sum += n
    }

    fmt.Printf("Sum of %v = %d\n", numbers, sum)
}
