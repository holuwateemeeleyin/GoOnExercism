package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	type app struct {
		emoji rune
		name  string
		index int
	}

	apps := []app{
		{emoji: '❗', name: "recommendation", index: -1},
		{emoji: '⚠', name: "warning", index: -1},
		{emoji: 'ℹ', name: "info", index: -1},
		{emoji: '🔍', name: "search", index: -1},
		{emoji: '☀', name: "weather", index: -1},
	}

	for i := range apps {
		if idx := strings.IndexRune(log, apps[i].emoji); idx != -1 {
			apps[i].index = idx
		}
	}

	minIndex := len(log)
	result := "default"

	for _, a := range apps {
		if a.index != -1 && a.index < minIndex {
			minIndex = a.index
			result = a.name
		}
	}

	return result
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Map(func(r rune) rune {
		if r == oldRune {
			return newRune
		}
		return r
	}, log)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
