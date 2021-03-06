package format

import (
	"strings"
)

func (f *Format) Initialize() {
	f.Fg.Black = string(Black)
	f.Fg.Blue = string(Blue)
	f.Fg.Cyan = string(Cyan)
	f.Fg.Gray = string(Gray)
	f.Fg.Green = string(Green)
	f.Fg.LightBlue = string(LightBlue)
	f.Fg.LightCyan = string(LightCyan)
	f.Fg.LightGray = string(LightGray)
	f.Fg.LightGreen = string(LightGreen)
	f.Fg.LightMagenta = string(LightMagenta)
	f.Fg.LightRed = string(LightRed)
	f.Fg.LightYellow = string(LightYellow)
	f.Fg.Magenta = string(Magenta)
	f.Fg.NC = string(NC)
	f.Fg.Red = string(Red)
	f.Fg.White = string(White)
	f.Fg.Yellow = string(Yellow)

	f.Em.Blinking = string(Blinking)
	f.Em.Bold = string(Bold)
	f.Em.Inverted = string(Inverted)
	f.Em.Underlined = string(Underlined)
}

func FormatChar(opts ...string) string {
	return "\\[\\033[" + strings.Join(opts, ";") + "m\\]"
}

func Ellipse(text string, length int) string {
	if len(text) > length {
		return text[:(length-3)] + "..."
	}
	return text
}
