package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Print("Enter Password Length: ")
	length, err := getLengthInput()
	if err != nil {
		log.Fatal(err)
	}
	password := generatePassword(length)
	fmt.Println("Generated Password:", password)
}

func generatePassword(length int64) string {
	passwordChars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-={}[]:;,.<>?")
	passwordCharsLength := len(passwordChars)
	password := make([]rune, length)
	for index := range length {
		password[index] = passwordChars[rand.Intn(passwordCharsLength)]
	}
	return string(password)
}

func getLengthInput() (int64, error) {
	reader := bufio.NewReader(os.Stdin)
	length, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	length = strings.ReplaceAll(length, "\n", "")
	lengthNum, err := strconv.ParseInt(length, 0, 64)
	if err != nil || lengthNum <= 0 {
		return 0, fmt.Errorf("Your length: %s, is not a valid number\n", length)
	}
	return lengthNum, nil
}
