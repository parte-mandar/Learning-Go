package main

import "fmt"
import "rsc.io/quote/v4"
import "math"
import "google.golang.org/appengine"

func main() {
    fmt.Println("Namaskaram! " + quote.Go())
    fmt.Println(math.Log(82.323)) // Standard math library
    fmt.Println(appengine.IsSecondGen()) // Some random (non-standard) package tested
}