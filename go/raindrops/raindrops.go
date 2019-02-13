package raindrops

import (
    "strconv"
    "fmt"
    )

// Convert is a FizzBuzz-like function that checks to see if the value is divisible
// by 3 (return 'Pling'), by 5 (returns 'Plang') or by 7 (returns 'Plong')
func Convert (drop int) (string) {
    results := ""
    if drop % 3 == 0 {
        results += "Pling"
    }
    if drop % 5 == 0 {
        results += "Plang"
    }
    if drop % 7 == 0 {
        results += "Plong"
    }

    if len(results) == 0 {
        results += strconv.Itoa(drop)
    }
    return results
}

func main() {
    fmt.Println(Convert(3))
}
