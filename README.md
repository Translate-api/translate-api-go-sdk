# Translate API - Go SDK

Official Go SDK for [Translate API](https://translate-api.com).

## 🚀 Quick Start (For Beginners)

### Step 1: Get Your API Key
1. Go to [https://translate-api.com](https://translate-api.com)
2. Click "Login" or "Get Started"
3. Create an account (it's free to start!)
4. Go to your Dashboard
5. Click "Create API Key"
6. Copy your API key - you'll need it!

### Step 2: Install the SDK
In your Go project folder, run:

```bash
go get github.com/Translate-api/translate-api-go-sdk
```

### Step 3: Use It!

```go
package main

import (
    "fmt"
    translateapi "github.com/Translate-api/translate-api-go-sdk"
)

func main() {
    // Replace "your-api-key" with your actual API key from translate-api.com
    client := translateapi.NewClient("your-api-key")

    // Translate to one language
    result, err := client.Translate("Hello world", "es")
    if err != nil {
        panic(err)
    }
    fmt.Println(result.Translations["es"]) // Output: "Hola mundo"

    // Translate to multiple languages
    result, err = client.Translate("Hello world", []string{"es", "fr", "de"})
    if err != nil {
        panic(err)
    }
    fmt.Println(result.Translations)
    // Output: map[de:Hallo Welt es:Hola mundo fr:Bonjour le monde]
}
```

## 📖 Full API Reference

### NewClient(apiKey)
Create a new client with your API key.

### Translate(text, targetLanguage)
Translate text. targetLanguage can be a string or []string.

## 🌍 Supported Languages
Visit [translate-api.com/documentation](https://translate-api.com/documentation) for a full list.

## ❓ Need Help?
- Documentation: [translate-api.com/documentation](https://translate-api.com/documentation)
- Support: support@translate-api.com

## 📝 License
MIT License
