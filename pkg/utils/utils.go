package utils

import (
	"fmt"
	"strconv"
)

func ColoredText(text string, index int) string {
	coloredText := fmt.Sprintf(
		"%s%s%s%s%s",
		"\x1b[3", strconv.Itoa(index), "m", text, "\x1b[0m",
	)
	return coloredText
}

func Union(array1 []any, array2 []any) {

}
