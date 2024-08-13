package main

import (
	"fmt"
	"os"
)

func main() {
	lexemeMap := map[string]rune{
		"LEFT_PAREN":  '(',
		"RIGHT_PAREN": ')',
		"LEFT_BRACE":  '{',
		"RIGHT_BRACE": '}',
		"COMMA":       ',',
		"DOT":         '.',
		"MINUS":       '-',
		"PLUS":        '+',
		"SEMICOLON":   ';',
		"STAR":        '*',
	}
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage
	//
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	stringContent := string(fileContents)

	for _, str := range stringContent {
		switch str {
		case lexemeMap["LEFT_PAREN"]:
			fmt.Printf("LEFT_PAREN %v null\n", string(lexemeMap["LEFT_PAREN"]))
		case lexemeMap["RIGHT_PAREN"]:
			fmt.Printf("RIGHT_PAREN %v null\n", string(lexemeMap["RIGHT_PAREN"]))
		case lexemeMap["LEFT_BRACE"]:
			fmt.Printf("LEFT_BRACE %v null\n", string(lexemeMap["LEFT_BRACE"]))
		case lexemeMap["RIGHT_BRACE"]:
			fmt.Printf("RIGHT_BRACE %v null\n", string(lexemeMap["RIGHT_BRACE"]))
		case lexemeMap["COMMA"]:
			fmt.Printf("COMMA %v null\n", string(lexemeMap["COMMA"]))
		case lexemeMap["DOT"]:
			fmt.Printf("DOT %v null\n", string(lexemeMap["DOT"]))
		case lexemeMap["MINUS"]:
			fmt.Printf("MINUS %v null\n", string(lexemeMap["MINUS"]))
		case lexemeMap["PLUS"]:
			fmt.Printf("PLUS %v null\n", string(lexemeMap["PLUS"]))
		case lexemeMap["SEMICOLON"]:
			fmt.Printf("SEMICOLON %v null\n", string(lexemeMap["SEMICOLON"]))
		case lexemeMap["STAR"]:
			fmt.Printf("STAR %v null\n", string(lexemeMap["STAR"]))
		}
	}
	fmt.Println("EOF  null")
}
