package main

import (
	"fmt"
	"os"
)

func prinHelp() {
	fmt.Println("tanasinn...")
	fmt.Println("usage: tanasin [OPTION] [FILE]")
}

type option struct {
	input * os.File
	isHelp bool
}

func main() {
	var option option

	for i := 1; i < len(os.Args); i++ {

		if os.Args[i][0] == '-' {

			if len(os.Args[i]) < 2 {
				option.input = os.Stdin

			} else {
				switch os.Args[i][1] {
				case 'h':
					option.isHelp = true
				}
			}

		} else if option.input == nil {
			f, err := os.OpenFile(os.Args[i], os.O_RDONLY, 0x600)
			if err != nil {
				fmt.Printf("cannot open `%s`\n", os.Args[i])
				os.Exit(1)
			}
			option.input = f
		}
	}

	fmt.Println(option)
}

