package twelve

import (
	"fmt"
)

// Verse takes in the day of Christmas and returns the lyrics for that verse
func Verse(day int) string {
	dayOfXmas := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}
	gaveToMe := []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, and ",
		"three French Hens, ",
		"four Calling Birds, ",
		"five Gold Rings, ",
		"six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ",
		"eight Maids-a-Milking, ",
		"nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ",
		"eleven Pipers Piping, ",
		"twelve Drummers Drumming, ",
	}

	gifts := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", dayOfXmas[day-1])

	for num := day - 1; num >= 0; num-- {
		fmt.Println("index, day, gifts", num, day, gifts)
		gifts += gaveToMe[num]
	}
	return gifts
}

// Song returns the lyrics to "Twelve Days of Christmas"
func Song() string {
	songLyrics := ""

	for day := 1; day <= 12; day++ {
		songLyrics += Verse(day) + "\n"
	}
	return songLyrics
}
