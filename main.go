package main

import (
	"bufio"        // for buffered IO operations (more efficient)
	"fmt"          // for formatted IO like printf
	"io"           // used to detect end of input
	"os"           // handling standard IO and file stats
	"strings"      // for string manipulation
	"unicode/utf8" // to handle ASCII
)

// building the speech balloon
// buildBalloon takes a slice of strings of max width maxWidth
// prepends/appends margins on first and last line, and at start/end of each line
// and returns a string with the contents of the balloon
func buildSpeechBalloon(lines []string, maxWidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "|", "<", ">"}

	horizontalBorder := " " + strings.Repeat("_", maxWidth+2)

	ret = append(ret, horizontalBorder)

	if count == 1 {
		s := fmt.Sprintf(`%s %s %s`, borders[3], lines[0], borders[4]) // < text >
		ret = append(ret, s)
	} else {
		// top diagonals
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1]) // / text \
		ret = append(ret, s)

		// left and right walls
		for i := 1; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[2]) // | text |
			ret = append(ret, s)
		}

		// bottom diagonals
		s = fmt.Sprintf(`%s %s %s`, borders[1], lines[count-1], borders[0]) // \ text /
		ret = append(ret, s)
	}

	ret = append(ret, horizontalBorder)
	return strings.Join(ret, "\n")
}

// tabsToSpaces converts all tabs found in the strings
// found in the `lines` slice to 4 spaces, to prevent misalignments in
// counting the runes
func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

// return length of the string with max length based on slice of strings given
func calculatemaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		length := utf8.RuneCountInString(l)
		if length > w {
			w = length
		}
	}

	return w
}

// append spaces to the end of each string to ensure that all strings are maxWidth long
func normaliseLines(lines []string, maxwdith int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwdith-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}

	return ret
}

func main() {
	// retrieve metadata about the standard input, and _ ignores error returned
	info, _ := os.Stdin.Stat()

	// get the file mode bits and check if the user ran the command without a pipe
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("This command is intended to work with pipes.")
		fmt.Println("Usage: [text string] | gocowsay")
		fmt.Println("Works with fortune too")
		return
	}

	// create a buffered reader to read inputs more efficiently
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine() // read one line at a time, _ takes the size of the return value
		if err != nil && err == io.EOF {  // checks for error and end of input
			break
		}

		lines = append(lines, string(line))
	}

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = tabsToSpaces(lines)
	maxWidth := calculatemaxWidth(lines)
	messages := normaliseLines(lines, maxWidth)
	speechWithBalloon := buildSpeechBalloon(messages, maxWidth)
	fmt.Println(speechWithBalloon)
	fmt.Println(cow)
}
