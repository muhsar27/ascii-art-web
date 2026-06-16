package functions

import "strings"

func PrintAsciiArt(inputString string, bannerLines []string) string {
	inputString = strings.ReplaceAll(inputString, "\r\n", "\n")
	stringWords := strings.Split(inputString, "\n")
	var result strings.Builder
	for idx, word := range stringWords {
		if word == "" {
			if idx != len(stringWords)-1 {
				result.WriteString("\n")
			}
			continue
		}
		for i := 1; i <= 9; i++ {
			for j := 0; j < len(word); j++ {
				numLine := (int(word[j])-32)*9 + i
				result.WriteString(bannerLines[numLine])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
