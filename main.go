package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func prinHelp() {
	fmt.Println("tanasinn...")
	fmt.Println("usage: tanasin [OPTION] [FILE]")
}

type option struct {
	input * os.File
	isHelp bool
	th []float64
}

type parseError struct {
	s string
	reason string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("cannot parse %s because %s", e.s, e.reason)
}

func parseThreasholds(th string) ([]float64, error) {
	var ths []float64

	if strings.Contains(th, ",") {
		list := strings.Split(th, ",")
		for i := 0; i < len(list); i++ {

			f, err := strconv.ParseFloat(list[i], 32)
			if err != nil {
				log.Fatal(err)
				return ths, &parseError{s: th, reason: "not float in the list"}
			}

			ths = append(ths, f)
		}

	} else {
		f, err := strconv.ParseFloat(th, 32)
		if err != nil {
			log.Fatal(err)
			return []float64{}, &parseError{s: th, reason: "not float"}
		}
		ths = []float64{f}
	}

	return ths, nil
}

func tanasinnize(th float64, r rune) string {
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
	if th < rand.Float64() {
		return fmt.Sprintf("%s", s[int(rand.Uint32()) % len(s)])
	} else {
		return fmt.Sprintf("%c", r)
	}
}

func main() {
	var option = option{input: os.Stdin, th: []float64{0.8}, isHelp: false}
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
					ths, err := parseThreasholds(os.Args[i])
					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}
					option.th = ths
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
	var text []rune = []rune{}
	for true {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		text = append(text, r)
	}

	for n := 0; n < len(text); n++ {
		p := float64(n) / float64(len(text))
		i := int(p * float64(len(option.th)))
		r := text[n]
		if r == '\n' {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%s", tanasinnize(option.th[i], r))
		}
	}
}
