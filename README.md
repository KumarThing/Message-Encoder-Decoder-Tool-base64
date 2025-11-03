# Message Encoder & Decoder Tool

A simple command-line tool written in Go to encode, decode, save, and load messages using Base64 encoding. It also has placeholders for custom cipher encoding and decoding.

---

## Features

- Encode messages in Base64
- Decode Base64 messages
- Save encoded messages to a file
- Load and decode messages from a file
- Placeholder for custom cipher encode/decode
- Exit the application

---



### bash
```
go run main.go
```

# Usage 
--------------------------------------------------------
1. encode message
2. decode message
3. save encode message
4. load and decode file
5. custom cipher encode
6. custom cipher decode
7. exit
--------------------------------------------------------

# Example

## Encoding a message:

```
Enter command: encode message
Enter your message: Hello World
Encoded message:
SGVsbG8gV29ybGQ=
```


## Saving the encoded message:

```
Enter command: save encode message
Encoded message saved to encoded_message.txt
```

## Loading and decoding a saved message:

```
Enter command: load and decode file
Decoded message from file:
Hello World
```

