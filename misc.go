package termui

import (
	"strings"
)

var PRINTABLE_STRINGS = append(
	strings.Split(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,./<>?;:'\"[]\\{}|`~!@#$%^&*()-_=+",
		"",
	),
	"<Space>",
	"<Tab>",
	"<Enter>",
)

type Alignment uint

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)
