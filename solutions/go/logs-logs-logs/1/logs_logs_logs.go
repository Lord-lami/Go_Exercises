package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case 0x2600:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	mutableLog := []rune(log)
	for i, char := range mutableLog {
		if char == oldRune {
			mutableLog[i] = newRune
		}
	}
	return string(mutableLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	charCount := 0
	for _, char := range log {
		char = char
		charCount++
	}
	return charCount <= limit
}
