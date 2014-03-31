package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

type symbol [6][9]rune

// width:  9 * 10 + 1 * 8
// height: 6 * 20
var framebuffer []symbol
var cursor [2]int = {0, 0}

var symbolTable [26]symbol

func showHelp() {
	fmt.Println("usage: asciifont -f [font file]")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadSymbolTable(path string) {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	reader := bufio.NewReader(file)

	stat := 0
	font := 0

	for {

		switch stat {

		case 0:
			line, err := reader.ReadString('\n')
			line = strings.Trim(line, "\n")
			if len(line) == 0 {
				goto OUT
			}

			if err != nil {
				// including EOF
				goto OUT
			}

			for _, r := range line {
				font = int(r - 65)
				if font < 0 || font > 25 {
					panic("invalid font number")
				}
			}

			stat = 1

		case 1:
			for i := 0; i < 6; i++ {
				line, err := reader.ReadString('\n')
				line = strings.Trim(line, "\n")
				if len(line) == 0 {
					continue
				}

				if err != nil {
					// including EOF
					goto OUT
				}

				for j, r := range line {
					if j > 8 {
						break
					}
					symbolTable[font][i][j] = r
				}
			}

			stat = 0
		}
	}

OUT:
}

func printSymbolTable() {
	for i := 0; i < 26; i++ {
		printSymbol(i)
	}
}

func printSymbol(index int) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", symbolTable[index][i][j])
		}
		fmt.Println()
	}
}

func readStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if err != nil {
			break
		}
	}
}

func convert(raw string) {
	for _, r := range raw {
		fmt.Printf("%c %d\n", r, r)
	}
}

func main() {
	h := flag.Bool("h", false, "show help messages")
	f := flag.String("f", "./fonts/tukasans.afont", "specify font file")
	flag.Parse()
	if *h {
		showHelp()
		return
	}

	loadSymbolTable(*f)

	//printSymbolTable()

	convert("ABC")

}
