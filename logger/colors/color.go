package colors

import (
	"fmt"
	"io"
)

type outputMode int

// DiscardNonColorEscSeq supports the divided color escape sequence.
// But non-color escape sequence is not output.
// Please use the OutputNonColorEscSeq If you want to output a non-color
// escape sequences such as ncurses. However, it does not support the divided
// color escape sequence.
const (
	_ outputMode = iota
	DiscardNonColorEscSeq
	OutputNonColorEscSeq
)

// NewColorWriter creates and initializes a new ansiColorWriter
// using io.Writer w as its initial contents.
// In the console of Windows, which change the foreground and background
// colors of the text by the escape sequence.
// In the console of other systems, which writes to w all text.
func NewColorWriter(w io.Writer) io.Writer {
	return NewModeColorWriter(w, DiscardNonColorEscSeq)
}

// NewModeColorWriter create and initializes a new ansiColorWriter
// by specifying the outputMode.
func NewModeColorWriter(w io.Writer, mode outputMode) io.Writer {
	if _, ok := w.(*colorWriter); !ok {
		return &colorWriter{
			w:    w,
			mode: mode,
		}
	}
	return w
}

func Bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

// Black returns a black string
func Black(message string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", message)
}

// White returns a white string
func White(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Cyan returns a cyan string
func Cyan(message string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", message)
}

// Blue returns a blue string
func Blue(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", message)
}

// Red returns a red string
func Red(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}

// Green returns a green string
func Green(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
}

// Yellow returns a yellow string
func Yellow(message string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", message)
}

// Gray returns a gray string
func Gray(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Magenta returns a magenta string
func Magenta(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", message)
}

// BlackBold returns a black Bold string
func BlackBold(message string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", Bold(message))
}

// WhiteBold returns a white Bold string
func WhiteBold(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", Bold(message))
}

// CyanBold returns a cyan Bold string
func CyanBold(message string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", Bold(message))
}

// BlueBold returns a blue Bold string
func BlueBold(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", Bold(message))
}

// RedBold returns a red Bold string
func RedBold(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", Bold(message))
}

// GreenBold returns a green Bold string
func GreenBold(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", Bold(message))
}

// YellowBold returns a yellow Bold string
func YellowBold(message string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", Bold(message))
}

// GrayBold returns a gray Bold string
func GrayBold(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", Bold(message))
}

// MagentaBold returns a magenta Bold string
func MagentaBold(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", Bold(message))
}
