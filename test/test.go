package main

import (
	"fmt"
	"os"
	"path/filepath"
	"wordcount"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	wc := make(wordcount.WordCount)
	// for _, filename := range os.Args[1:] {
	//  wordcount.UpdateFreq(filename)
	// }
	wc.WordFreqCounter(os.Args[1:])

	wc.SortReport()
}
