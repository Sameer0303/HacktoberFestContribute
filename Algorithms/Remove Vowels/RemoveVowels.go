package main

import (
        "fmt"
        "regexp"
        "os"
)

func main() {
        re := regexp.MustCompile("a|e|i|o|u")
        removedVowels := re.ReplaceAllLiteralString(os.Args[1], "") 
        fmt.Println(removedVowels)
}
