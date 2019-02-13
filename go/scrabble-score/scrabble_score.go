package scrabble

import (
    "fmt"
    "strings"
)

// Score takes in a word that would be played in Scrabble and returns the score as an int
func Score (word string) (int) {
    // Increment this value to get final score
    score := 0

    // Classify all letters according to their score value
    oneScore := "AEIOULNRST"
    twoScore := "DG"
    threeScore := "BCMP"
    fourScore := "FHVWY"
    fiveScore := "K"
    eightScore := "JX"
    tenScore := "QZ"

    // Check for empty string
    if word == "" {
        return score
    }
    // Convert all letters in word to uppercase
    uWord := strings.ToUpper(word)

    // Compare the different letters to their scored counterparts
    for i := range uWord {
        for j := range oneScore {
            if uWord[i] == oneScore[j] {
                score += 1
                fmt.Println(i, j, score)
            }
        }
        for k := range twoScore {
            if uWord[i] == twoScore[k] {
                score += 2
                fmt.Println(score)
            }
        }
        for l := range threeScore {
            if uWord[i] == threeScore[l] {
                score += 3
                fmt.Println(score)
            }
        }
        for m := range fourScore {
            if uWord[i] == fourScore[m] {
                score += 4
                fmt.Println(score)
            }
        }
        for n := range fiveScore {
            if uWord[i] == fiveScore[n] {
                score += 5
                fmt.Println(score)
            }
        }
        for o := range eightScore {
            if uWord[i] == eightScore[o] {
                score += 8
                fmt.Println(score)
            }
        }
        for p := range tenScore {
            if uWord[i] == tenScore[p] {
                score += 10
                fmt.Println(score)
            }
        }
    }
    return score
}
