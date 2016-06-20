package main

import (
    "fmt"
    "flag"
    "strconv"
)

func main() {
    basesize := flag.Float64("base", 16, "base font size")
    flag.Parse()
    
    if flag.NArg() > 0 {
        if target, err := strconv.ParseFloat(flag.Arg(0), 64); err == nil {
            rem := target / *basesize
            if px, pxerr := strconv.Atoi(flag.Arg(0)); pxerr == nil {
                fmt.Println(createText(px, rem))
            }
        }
    } else {
        for px := 1; px <= 100; px++ {
            rem := float64(px) / *basesize
            fmt.Println(createText(px, rem))
        }
    }
}

func createText(px int, rem float64) string {
    text := strconv.Itoa(px) + "px = " + strconv.FormatFloat(rem, 'G', -1, 64) + "rem"
    return text
}