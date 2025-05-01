package ascii

import (
	"os"
	"strings"
)

// GenerateArt creates ASCII art based on the input text and selected banner style.
// It reads the appropriate banner file, creates a map of ASCII characters,
// and generates the ASCII art string.
func GenerateArt(inputUser string, banner string) (string, error) {
	// reads banner file
	fileContent, err := os.ReadFile("Banners/" + banner + ".txt")
	if err != nil {
		return "", err
	}
	
	splitFile := strings.Split(string(fileContent), "\r\n")
	
	// Create a map of ASCII characters from the banner file
	asciiMap := make(map[rune][]string)
	dec := 31
	for _, line := range splitFile {
		if line == "" {
			dec++
		} else {
			asciiMap[rune(dec)] = append(asciiMap[rune(dec)], line)
		}
	}

	// Split the input text into lines
	splittedInput := strings.Split(inputUser, "\r\n")
	// Generate Art
	var result string
	for _, line := range splittedInput {
		if line == "" {
			result += "\n"
			continue
		}

		ascii := [][]string{}
		for _, char := range line {
			ascii = append(ascii, asciiMap[char])
		}
		for i := 0; i < 8; i++ {
			for _, char := range ascii {
				result += char[i]
			}
			result += "\n"
		}
	}
	return result, nil
}
