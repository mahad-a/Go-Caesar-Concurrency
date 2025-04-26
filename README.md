# Go Caesar Cipher with Concurrency

This project implements a Caesar cipher encoder in Go, with support for concurrent processing of multiple messages using goroutines and channels. The code is modular and demonstrates clean string manipulation, Unicode handling, and concurrency primitives.

## Caesar Cipher Overview

The Caesar cipher is a simple encryption method that shifts each letter in a message by a fixed number of positions in the alphabet. If a letter goes past `'Z'`, it wraps around to the start of the alphabet.

This implementation includes:
- Removal of non-alphabetic characters.
- Conversion to uppercase.
- Shift logic that wraps A–Z.
- Concurrency support for processing messages in parallel.

---

## File Structure

```plaintext
├── cipher.go             # CaesarCipher function (core logic)
├── concurrent_single.go  # Processes the full message list in a single goroutine
├── concurrent_split.go   # Processes the message list split across 3 goroutines
├── main.go               # Entry point with demos for all implementations
```

---

## How to Run

Make sure you have Go installed: [https://go.dev/dl](https://go.dev/dl)

Then run:
```bash
go run main.go
```

You’ll see three demo sections:

1. **CaesarCipher** – Encrypts individual messages with sample outputs.
2. **Single Goroutine** – Encrypts a list of messages concurrently.
3. **Three Goroutines** – Splits the list and processes it in parallel via three goroutines.

---

## Sample Output

```go
fmt.Println(CaesarCipher("I love CS!", 5))
// Output: NQTAJHX
```

Concurrent encrypted outputs (example):
```
ECTNGVQPEC
DTKIJVURCEGECTNGVQPECFNJQOG
UAUE
HWPEVKQPCNKUUV
EQNQPGNDAFT
UAUE
EQPEWTTGPVKUPF
RCTCFKIOU
YKPVGT
```