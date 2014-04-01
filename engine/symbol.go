package engine

import (
	"bufio"
	"com/github/crackcell/asciitype/utils"
	"fmt"
	"os"
	"strings"
)

const (
	SYMBOL_HEIGHT = 6
	SYMBOL_WIDTH  = 9
)

type Symbol [SYMBOL_HEIGHT][SYMBOL_WIDTH]rune

type SymbolTable struct {
	table [26]Symbol
}

func (st *SymbolTable) Load(path string) {
	file, err := os.Open(path)
	utils.Check(err)

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

			for _, ch := range line {
				font = int(ch - 65)
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
					st.table[font][i][j] = r
				}
			}

			stat = 0
		}
	}

OUT:
}

func (st *SymbolTable) PrintSymbolTable() {
	for i := 0; i < 26; i++ {
		st.PrintSymbol(i)
	}
}

func (st *SymbolTable) PrintSymbol(index int) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", st.table[index][i][j])
		}
		fmt.Println()
	}
}

func (st *SymbolTable) GetSymbol(ch rune) *Symbol {
	i := ch - 65
	if i < 0 || i > 25 {
		panic("invalid char")
	}
	return &st.table[i]
}
