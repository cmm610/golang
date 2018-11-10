package twelve

import "fmt"

var days = [...]string{"", "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var lyrics = [...]string{
	"",
	" a Partridge in a Pear Tree.",
	" two Turtle Doves, and",
	" three French Hens,",
	" four Calling Birds,",
	" five Gold Rings,",
	" six Geese-a-Laying,",
	" seven Swans-a-Swimming,",
	" eight Maids-a-Milking,",
	" nine Ladies Dancing,",
	" ten Lords-a-Leaping,",
	" eleven Pipers Piping,",
	" twelve Drummers Drumming,",
}
var verseStart = "On the %s day of Christmas my true love gave to me:"

func Song() string {
	song := ""
	for i := 1; i<len(lyrics); i++ {
		song += Verse(i) + "\n"
	}
	return song
}

func Verse(vn int) string {
	verse := fmt.Sprintf(verseStart, days[vn])
	for i:= vn; i>=1; i-- {
		verse += lyrics[i]
	}

	return verse
}
