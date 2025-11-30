// Example usage of the Translate API Go SDK
//
// Before running this example:
// 1. Get your API key at https://translate-api.com
// 2. Replace "YOUR_API_KEY_HERE" with your actual API key
//
// To run this example:
// 1. Open terminal in the example folder
// 2. Run: go run main.go
package main

import (
	"fmt"
	translateapi "github.com/Translate-api/translate-api-go-sdk"
)

// ⚠️ IMPORTANT: Replace this with your API key from translate-api.com
const APIKey = "YOUR_API_KEY_HERE"

func main() {
	fmt.Println("🌍 Translate API - Go SDK Example")
	fmt.Println()

	// Create the client
	client := translateapi.NewClient(APIKey)

	// Example 1: Translate to a single language
	fmt.Println("Example 1: Translate to Spanish")
	result1, err := client.Translate("Hello, how are you?", "es")
	if err != nil {
		handleError(err)
		return
	}
	fmt.Println("Input: Hello, how are you?")
	fmt.Printf("Spanish: %s\n\n", result1.Translations["es"])

	// Example 2: Translate to multiple languages
	fmt.Println("Example 2: Translate to multiple languages")
	result2, err := client.Translate("Good morning!", []string{"fr", "de", "it", "ja"})
	if err != nil {
		handleError(err)
		return
	}
	fmt.Println("Input: Good morning!")
	fmt.Printf("French: %s\n", result2.Translations["fr"])
	fmt.Printf("German: %s\n", result2.Translations["de"])
	fmt.Printf("Italian: %s\n", result2.Translations["it"])
	fmt.Printf("Japanese: %s\n\n", result2.Translations["ja"])

	fmt.Println("✅ All examples completed successfully!")
	fmt.Println()
	fmt.Println("📖 Documentation: https://translate-api.com/documentation")
}

func handleError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
	fmt.Println()
	fmt.Println("💡 Make sure you:")
	fmt.Println("   1. Have a valid API key from https://translate-api.com")
	fmt.Println("   2. Replaced YOUR_API_KEY_HERE with your actual key")
	fmt.Println("   3. Have enough character quota in your account")
}
