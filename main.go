package main

import (
	"fmt"

	translator "github.com/firattamur/go-googletranslate/googletranslate"
)

func main() {

	text := "Merhaba Dünya!"

	googleTranslator, err := translator.NewDefaultGoogleTranslator()

	if err != nil {
		fmt.Printf("Error creating translator: %v", err)
		return
	}

	translatedText, err := googleTranslator.Translate(text, translator.TURKISH, translator.ENGLISH)

	if err != nil {
		fmt.Printf("Error translating text: %v", err)
		return
	}

	fmt.Printf("Original text   : %s\nTranslated text : %s\n", text, translatedText)

}
