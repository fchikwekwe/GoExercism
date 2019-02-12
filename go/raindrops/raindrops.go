package raindrops

import (
    "strconv"
    "fmt"
    )

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
