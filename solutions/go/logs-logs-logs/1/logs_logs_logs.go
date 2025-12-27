package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	var apps = map[string]int{
		"recommendation": find(log, '❗'),
		"search":         find(log, '🔍'),
		"weather":        find(log, '☀'),
	}
	lowestIndex := -1
	application := "default"

	for app, index := range apps {
		if index > -1 && (lowestIndex == -1 || index < lowestIndex) {
			//if index < lowestIndex {
			lowestIndex = index
			application = app
		}
	}
	return application
}

func find(log string, letter rune) int {
	for index, char := range log {
		if char == letter {
			return index
		}
	}
	return -1
}

// Replace replaces all occurrences of old with new, returning the modified log to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, char := range log {
		if char == oldRune {
			newLog = newLog + string(newRune)
		} else {
			newLog = newLog + string(char)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
