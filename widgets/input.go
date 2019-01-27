package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

type Input struct {
	Block
	TextWrap    bool
	Text        []string
	TextStyle   Style
	CursorStyle Style
	CursorPoint image.Point
	Password    bool
}

func NewInput() *Input {
	return &Input{
		Block:       *NewBlock(),
		TextWrap:    true,
		Text:        []string{},
		TextStyle:   Theme.Input.Text,
		CursorStyle: Theme.Input.Cursor,
	}
}

func (self *Input) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	var text []string
	// if self.TextWrap {
	// text = strings.Split(wordwrap.WrapString(self.Text, uint(self.Inner.Dx())), "\n")
	// } else {
	text = self.Text
	// }

	for i, line := range text {
		for j, r := range line {
			buf.SetCell(
				NewCell(r, self.TextStyle),
				image.Pt(self.Inner.Min.X+j, self.Inner.Min.Y+i),
			)
		}
	}

	buf.SetCell(
		NewCell(rune(self.Text[self.CursorPoint.Y][self.CursorPoint.X]), self.CursorStyle),
		self.CursorPoint.Add(image.Pt(self.Inner.Min.X, self.Inner.Min.Y)),
	)
}

func (self *Input) Backspace() {
}

func (self *Input) Keypress(key string) {
	switch key {
	case "<Enter>":
		self.CursorPoint.X = 0
		self.CursorPoint.Y++
	case "<Tab>":
		self.Text[self.CursorPoint.Y] += "\t"
	case "<Space>":
		self.Text[self.CursorPoint.Y] += " "
	default:
		self.Text[self.CursorPoint.Y] += key
	}
}

func (self *Input) MoveLeft() {
	if self.CursorPoint.X > 0 {
		self.CursorPoint.X--
	}
}

func (self *Input) MoveRight() {
	if self.CursorPoint.X < self.Inner.Dx()-1 && self.CursorPoint.X < len(self.Text[self.CursorPoint.Y]) {
		self.CursorPoint.X++
	}
}

func (self *Input) MoveUp() {
	if self.CursorPoint.Y > 0 {
		self.CursorPoint.Y--
	}
}

func (self *Input) MoveDown() {
	if self.CursorPoint.Y < self.Inner.Dy()-1 {
		self.CursorPoint.Y++
	}
}
