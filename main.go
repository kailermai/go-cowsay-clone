package main

import (
	"bufio" // for buffered IO operations (more efficient)
	"fmt"   // for formatted IO like printf
	"io"    // used to detect end of input
	"os"    // handling standard IO and file stats
)

func main() {
	// retrieve metadata about the standard input, and _ ignores error returned
	info, _ := os.Stdin.Stat()

	// get the file mode bits and check if the user ran the command without a pipe
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("This command is intended to work with pipes.")
		fmt.Println("Usage: [text string] | gocowsay")
		return
	}

	// create a buffered reader to read inputs more efficiently
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF { // checks for error and end of input
			break
		}

		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
