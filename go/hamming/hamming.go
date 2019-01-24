package hamming

import "fmt"

func Distance(a, b string) (int, error) {
    hammingDistance := 0
    if len(a) == len(b) {
        for index := range a {
            aIndex := a[index]
            for index := range b {
                if aIndex == b[index] {
                    fmt.Println("a value: %v, b value: %v", aIndex, b[index])
                    hammingDistance += 1
                }
            }
        }
    }
    return hammingDistance, fmt.Errorf("cannot compare; strings are not the same length")
}
