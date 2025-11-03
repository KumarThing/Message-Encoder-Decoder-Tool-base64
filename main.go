package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var lastEncoded string // ✅ store the last encoded message for saving

	fmt.Println("\nWelcome to Message Encoder & Decoder Tool")
	fmt.Println(`
--------------------------------------------------------
1. encode message
2. decode message
3. save encode message
4. load and decode file
5. exit
--------------------------------------------------------
`)
	fmt.Print("Enter command: ")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "encode message" {
			fmt.Print("Enter your message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			encoded := base64.StdEncoding.EncodeToString([]byte(message))
			lastEncoded = encoded // ✅ store the encoded message
			fmt.Println("✅ Encoded message:")
			fmt.Println(encoded)

		} else if input == "decode message" {
			fmt.Print("Enter your encoded message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			decodedBytes, err := base64.StdEncoding.DecodeString(message)
			if err != nil {
				fmt.Println("❌ Error decoding message:", err)
				continue
			}
			fmt.Println("✅ Decoded message:")
			fmt.Println(string(decodedBytes))

		} else if input == "save encode message" {
			if len(lastEncoded) == 0 {
				fmt.Println("❌ No encoded message to save. Please encode a message first.")
				continue
			}

			file, err := os.Create("encoded_message.txt")
			if err != nil {
				fmt.Println("❌ Error creating file:", err)
				continue
			}
			defer file.Close()

			_, err = file.WriteString(lastEncoded)
			if err != nil {
				fmt.Println("❌ Error writing to file:", err)
				continue
			}

			fmt.Println("✅ Encoded message saved to encoded_message.txt")

		} else if input == "load and decode file" {
			data, err := os.ReadFile("encoded_message.txt")
			if err != nil {
				fmt.Println("❌ Error reading file:", err)
				continue
			}

			decodedBytes, err := base64.StdEncoding.DecodeString(string(data))
			if err != nil {
				fmt.Println("❌ Error decoding file content:", err)
				continue
			}

			fmt.Println("✅ Decoded message from file:")
			fmt.Println(string(decodedBytes))

		} else if input == "exit" {
			fmt.Println("👋 Goodbye!")
			break

		} else {
			fmt.Println("⚠️ Unknown command. Please try again.")
		}

		fmt.Println(`
--------------------------------------------------------
1. encode message
2. decode message
3. save encode message
4. load and decode file
5. exit
--------------------------------------------------------
`)
		fmt.Print("Enter command: ")
	}
}
