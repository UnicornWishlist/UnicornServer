package prettylog

import (
	"log/slog"

	colors "github.com/nikonov1101/colors.go"
)

type DefaultColorize struct{}

// Attrs implements Colorize.
func (d DefaultColorize) Attrs(s string) string {
	return colors.Gray(s)
}

// Level implements Colorize.
func (d DefaultColorize) Level(level slog.Level, s string) string {
	switch level {
	case slog.LevelDebug:
		return colors.Green(s)
	case slog.LevelInfo:
		return colors.Blue(s)
	case slog.LevelWarn:
		return colors.Yellow(s)
	case slog.LevelError:
		return colors.Red(s)
	default:
		return colors.White(s)
	}
}

// Message implements Colorize.
func (d DefaultColorize) Message(s string) string {
	return colors.White(s)
}

// Timestamp implements Colorize.
func (d DefaultColorize) Timestamp(s string) string {
	return colors.Cyan(s)
}

func NewDefaultColorize() Colorize {
	return &DefaultColorize{}
}
