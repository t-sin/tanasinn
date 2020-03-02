package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func prinHelp() {
	fmt.Println("tanasinn...")
	fmt.Println("usage: tanasin [OPTION] [FILE]")
}

type option struct {
	input * os.File
	isHelp bool
	th float32
}


func tanasinnize(th float32, r rune) string {
	s := []string{
		"∴",
		"∵",
		"∴∵",
		")",
		"(･)",
		"(･",
		"∴･",
		"∴",
        "･･",
        "･:",
        "..",
	}
	if th < rand.Float32() {
		return fmt.Sprintf("%s", s[int(rand.Uint32()) % len(s)])
	} else {
		return fmt.Sprintf("%c", r)
	}
}

func main() {
	var option = option{input: os.Stdin, th: 0.7, isHelp: false}
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 1; i < len(os.Args); i++ {

		if os.Args[i][0] == '-' {

			if len(os.Args[i]) < 2 {
				option.input = os.Stdin

			} else {
				switch os.Args[i][1] {
				case 'h':
					option.isHelp = true
				case 't':
					i++
					f, err := strconv.ParseFloat(os.Args[i], 32)
					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}
					option.th = float32(f)
				}
			}

		} else if option.input == nil {
			f, err := os.OpenFile(os.Args[i], os.O_RDONLY, 0x600)
			if err != nil {
				log.Fatal("cannot open `%s`\n", os.Args[i])
				os.Exit(1)
			}
			option.input = f
			defer f.Close()
		}
	}

	if option.input == nil {
		option.input = os.Stdin
	}

	reader := bufio.NewReader(option.input)
	for true {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", tanasinnize(option.th, r))
	}
}
