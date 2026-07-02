package parsinglogfiles

import "regexp"

// IsValidLine checks if the line starts with a valid log level.
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

// SplitLogLine splits the log line using any field separator like <*>, <~>, etc.
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords counts how many lines contain the case-insensitive word "password" 
// enclosed in double quotes. 
func CountQuotedPasswords(lines []string) int {
	// Using [^"]* instead of .* ensures we stay within the quotes on a single line
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText removes "end-of-line" followed by any digits.
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName prefixes lines containing "User [name]" with "[USER] [name] ".
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	results := make([]string, 0, len(lines))

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			// Change "[USER]" to "[USR]" to make the test pass
			line = "[USR] " + match[1] + " " + line
		}
		results = append(results, line)
	}
	return results
}