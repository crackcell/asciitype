package main

import (
	"bufio"
	"com/github/crackcell/asciitype/engine"
	"flag"
	"fmt"
	"os"
	"strings"
)

func showHelp() {
	fmt.Println("usage: asciifont -f [font file]")
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

func convert(raw string, fb *engine.Framebuffer, st *engine.SymbolTable) {
	fb.Clear()
	raw = strings.ToUpper(raw)
	for _, r := range raw {
		fb.Append(st.GetSymbol(r))
	}

	fb.Flush(os.Stdout)
}

func main() {
	h := flag.Bool("h", false, "show help messages")
	f := flag.String("f", "./fonts/tukasans.afont", "specify font file")
	flag.Parse()
	if *h {
		showHelp()
		return
	}

	st := new(engine.SymbolTable)
	st.Load(*f)

	fb := engine.NewFramebuffer()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		if len(line) == 0 {
			continue
		}

		if err != nil {
			// including EOF
			break
		}

		convert(line, fb, st)
	}

}
