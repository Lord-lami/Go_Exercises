package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~*\**=*-*)+>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`".*[pP][aA][sS][sS][wW][oO][rR][dD].*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (taggedLines []string) {
	taggedLines = lines
	re := regexp.MustCompile(`User\s+(\S+)\s`)
	for i, line := range taggedLines {
		results := re.FindStringSubmatch(line)
		if results != nil {
			taggedLines[i] = "[USR] " + results[1] + " " + line
		}
	}
	return
}
