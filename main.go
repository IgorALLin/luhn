package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"

	//WarningColor = "\033[1;33m%s\033[0m"
	//DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Skip empty strings
		if len(scanner.Text()) == 0 {
			continue
		}

		result, err := checkCardNumber(scanner.Text())

		fmt.Printf(NoticeColor, "Input card number is: " + scanner.Text())
		fmt.Println("")

		if err == true {
			fmt.Printf(ErrorColor, result)
		} else {
			fmt.Printf(InfoColor, result)
		}

		fmt.Println("")
		fmt.Println("")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkCardNumber(cardNumber string) (message string, err bool) {
	// Delete all whitespaces from card number string
	cardNumber = strings.Replace(cardNumber, " ", "", -1)
	cardNumberLen := len(cardNumber)

	message = "Card number is correct."
	err = false

	newCardNumber := ""

	cardNumberSum := 0
	parity := cardNumberLen % 2

	for i, number := range cardNumber {
		// Convert ASCII value into int
		newNumber := int(number - '0')

		// If converted value > 9 or < 0 -> cardNumber string contains non digit sign
		if newNumber > 9 || newNumber < 0 {
			err = true
			message = "Incorrect card number. Card must contains only numbers."

			return
		}

		// If len of cardNumber is even -> multiply by 2 each odd char. Othervice multiply by 2 each even digit
		if(parity == 0 && (i + 1) % 2 != 0) || (parity == 1 && (i + 1) % 2 == 0) {
			newNumber *= 2
		}

		if newNumber > 9 {
			newNumber -= 9
		}

		newCardNumber += strconv.Itoa(newNumber)
		cardNumberSum += newNumber
	}

	if cardNumberSum % 10 != 0 {
		err = true
		message = "Wrong card number. Luhn card number sum is: " + strconv.Itoa(cardNumberSum) + ". Luhn card number is: " + newCardNumber
	} else {
		message += " Luhn card number sum is: " + strconv.Itoa(cardNumberSum) + ". Luhn card number is: " + newCardNumber
	}

	return
}
