package engine

import (
	"fmt"
	"os"
)

type Framebuffer struct {
	buffer []Symbol
}

func NewFramebuffer() *Framebuffer {
	return &Framebuffer{buffer: make([]Symbol, 0)}
}

func (fb *Framebuffer) Append(s *Symbol) {
	fb.buffer = append(fb.buffer, *s)
}

func (fb *Framebuffer) Clear() {
	fb.buffer = make([]Symbol, 0)
}

func (fb *Framebuffer) Flush(f *os.File) {
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		for j := 0; j < len(fb.buffer); j++ {
			for k := 0; k < SYMBOL_WIDTH; k++ {
				fmt.Fprintf(f, "%c", fb.buffer[j][i][k])
			}
			fmt.Fprintf(f, " ")
		}
		fmt.Fprintln(f)
	}

}
