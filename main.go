package main

import (
	"bufio"   // for buffered IO operations (more efficient)
	"fmt"     // for formatted IO like printf
	"io"      // used to detect end of input
	"os"      // handling standard IO and file stats
	"strings" // for string manipulation
	// "unicode/utf8" // to handle ASCII
)

// building the speech balloon
func buildSpeechBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "|", "<", ">"}

	horizontalBorder := " " + strings.Repeat("_", maxwidth+2)

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
			s = fmt.Sprintf(`%s %s %s`, borders[2], lines[0], borders[2]) // | text |
			ret = append(ret, s)
		}

		// bottom diagonals
		s = fmt.Sprintf(`%s %s %s`, borders[1], lines[0], borders[0]) // \ text /
		ret = append(ret, s)
	}

	ret = append(ret, horizontalBorder)
	return strings.Join(ret, "\n")
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

	speechBalloon := buildSpeechBalloon(lines, 15)
	fmt.Printf(speechBalloon)
}
