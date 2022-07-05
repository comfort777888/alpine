package ascii

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// OutputAscii func receives text to convert and font's name and returns converted result.
func OutputAscii(text string, font string) string {
	convertedString := stringCheck(text) // Receiving string that will be displayed at terminal.
	fontSymbols := convertFont(font)     // Creating array with ascii font.
	var result string
	for _, line := range convertedString { // Printing out string with font.
		a := fontCreator(line, fontSymbols)
		result += a
	}
	return result
}

// stringCheck func splits text by new lines and convert it to array
func stringCheck(stringToConvert string) []string {
	convertedString := strings.Split(stringToConvert, "\n")
	newlines := checkNewLine(convertedString) // checks if we received only newlines
	if len(newlines) > 0 {                    // checks if string consists only newlines
		convertedString = newlines
	}
	return convertedString
}

// ConvertFont checks if we can open file without any errors and returns error to POSTHandler.
func ConvertFont(font string) bool {
	_, err := os.ReadFile(font)
	// if err == nil {
	// 	return true
	// }
	return err == nil
}

func convertFont(font string) []string { // receiving font name
	data, err := os.ReadFile(font) // reading file with font which will be displayed through terminal.
	if err != nil {
		fmt.Println("error with font") // error cannot be as we already checked if file exist and if it have troubles to open it.
	}
	fontSymbols := strings.Split(string(data), "\n") // splits font's file by new lines
	return fontSymbols
}

// fontCreator cut ascii font and forms it althogether in line.
func fontCreator(s string, fontSymbols []string) string {
	result := ""
	if s == "" { // if we receive empty string - means we receive new line
		return "\n"
	}
	for l := 1; l <= 8; l++ { // 8 is the height of our font
		for _, i := range s { // we determine our symbol in fontSymbol
			result += fontSymbols[rune(i-32)*9+rune(l)] // as our font starts with 32nd rune and our font's file starts from the
		} // first line, the height of each symbols takes 8 line
		result += "\n" // and separate each line.
	}
	return result
}

// checkNewLine works if we receive a string only with newlines.
func checkNewLine(str []string) []string {
	count := 0
	var s []string
	for i := 0; i < len(str); i++ {
		if str[i] == "" {
			count++
		}
	}
	if count == len(str) {
		for i := 0; i < len(str)-1; i++ {
			s = append(s, "")
		}
	}
	return s
}

// Hashsum check if font's files were changed.
func HashSum(path string) bool {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return true
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Print(err)
		return true
	}
	hashSum := fmt.Sprintf("%x", h.Sum(nil))
	if path == "/standard.txt" {
		if hashSum != "ac85e83127e49ec42487f272d9b9db8b" {
			log.Println("Error: cannot read file. File was modified or empty for standard  font")
			return true
		}
	} else if path == "/shadow.txt" {
		if hashSum != "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			log.Println("Error: cannot read file. File was modified or empty for shadow font")
			return true
		}
	} else if path == "/thinkertoy.txt" {
		if hashSum != "8efd138877a4b281312f6dd1cbe84add" {
			log.Println("Error: cannot read file. File was modified or empty for thinkertoy font")
			return true
		}
	}
	return false
}
