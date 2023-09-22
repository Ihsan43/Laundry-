package validasi 

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
	"strconv"
)

func ValidateDateInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt + " : ")
		scanner.Scan()
		dateStr := scanner.Text()
		datePattern := `^\d{4}-\d{2}-\d{2}$`
		if regexp.MustCompile(datePattern).MatchString(dateStr) {
			return dateStr
		}
		fmt.Println("Error: Format tanggal tidak valid. Gunakan format YYYY-MM-DD.")
	}
}

func ValidateNumericInput(prompt string, zeroAllowed bool) string {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(prompt + " : ")
        scanner.Scan()
        inputStr := scanner.Text()
        num, err := strconv.Atoi(inputStr)
        if err == nil {
            if !zeroAllowed && num == 0 {
                fmt.Println("Error:Input Harus lebih dari 0.")
            } else {
                return inputStr
            }
        } else {
            fmt.Println("Error: Harus berupa angka.")
        }
    }
}

func ValidateMaxNumber(prompt string) string {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(prompt + " : ")
        scanner.Scan()
        inputStr := scanner.Text()
        num, err := strconv.Atoi(inputStr)
        if err == nil && num < 5 {
            return inputStr
        }
        fmt.Println("Error: Harus berupa angka kurang dari 5.")
    }
}


func ValidatePhoneNumber(prompt string) string {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(prompt + " : ")
        scanner.Scan()
        phoneNumber := scanner.Text()
        phoneNumberPattern := `^[0-9]{10}$` 
        if regexp.MustCompile(phoneNumberPattern).MatchString(phoneNumber) {
            return phoneNumber
        }
        fmt.Println("Error: Format nomor telepon tidak valid. Gunakan 10 digit angka.")
    }
}

func ValidateEmail(prompt string) string {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(prompt + " : ")
        scanner.Scan()
        email := scanner.Text()
        emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
        if regexp.MustCompile(emailPattern).MatchString(email) {
            return email
        }
        fmt.Println("Error: Format email tidak valid.")
    }
}