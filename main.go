package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/morgulbrut/color256"
	tm "github.com/morgulbrut/goterm"
)

func showPage(p, w int, h int) {
	if p == 0 {
		t := fmt.Sprintf(color256.HiPink(color256.Bold("Current Time: %s", time.Now().Format(time.RFC1123))))
		x := (w - len(t)) / 2
		y := h / 5

		tm.MoveCursor(x, y)
		tm.Println(t)

		tm.MoveCursor(x+3, y+4)
		tm.Println("Current window height:", h)
		tm.MoveCursor(x+3, y+5)
		tm.Println("Current window width:", w)
	} else {
		t := fmt.Sprintf(color256.HiPink(color256.Italic("Hello world")))
		x := (w - len(t)) / 2
		y := h / 5

		tm.MoveCursor(x, y)
		tm.Println(t)
		tm.MoveCursor(w, h)
	}
}

func main() {
	tm.Clear() // Clear current screen

	p := 0
	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.Clear()

		w := tm.Width()
		h := tm.Height()
		showPage(p, w, h)
		tm.Flush() // Call it every time at the end of rendering
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if char == 'q' {
			tm.Clear()
			tm.MoveCursor(1, 1)
			tm.Flush()
			os.Exit(0)
		}
		if key == keyboard.KeyArrowLeft {
			p = 0
		} else {
			p = 1
		}
		time.Sleep(100 * time.Millisecond)
	}
}
