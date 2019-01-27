// +build ignore

package main

import (
	"log"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	i := widgets.NewInput()
	i.Text = []string{"Hello World!"}
	i.SetRect(0, 0, 25, 5)

	ui.Render(i)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch {
			case e.ID == "<C-c>":
				return
			case e.ID == "<Left>":
				i.MoveLeft()
				ui.Render(i)
			case e.ID == "<Right>":
				i.MoveRight()
				ui.Render(i)
			case e.ID == "<Up>":
				i.MoveUp()
				ui.Render(i)
			case e.ID == "<Down>":
				i.MoveDown()
				ui.Render(i)
			case ui.ContainsString(ui.PRINTABLE_STRINGS, e.ID):
				i.Keypress(e.ID)
				ui.Render(i)
			}
		}
	}
}
