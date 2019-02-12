package raindrops

import (
    "strconv"
    "fmt"
    )

func RainSpeak (drop int) string {
    var result string
    if drop % 3 == 0 {
        result := fmt.Sprintf("Pling")
    } else if drop % 5 == 0 {
        result := fmt.Sprintf("Plong")
    } else if drop % 7 == 0 {
        result := fmt.Sprintf("Plang")
    } else {
        result := strconv.Itoa(drop)
    }
    return result
}

func main() {
    fmt.Println(RainSpeak(3))
}
