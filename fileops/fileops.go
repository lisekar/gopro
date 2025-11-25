package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const AccountBalanceFile = "balance.txt"

// function to write balance to a file
func WriteBalanceToFile(balance float64, filename string) {
	balanceText := fmt.Sprintf("Current Balance: %.2f\n", balance) // format balance text
	os.WriteFile(filename, []byte(balanceText), 0644)              // write to file with appropriate permissions
}

func GetBalanceFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 10000, errors.New("fail to find balance file") // return 0 if file does not exist or error occurs
	}
	balanceText := string(data)                                  // float64 cant be directly converted from []byte
	balanceValue, err := strconv.ParseFloat(balanceText[0:], 64) // extract balance value from text
	if err != nil {
		return 10000, errors.New("fail to parse balance value") // return 0 if parsing fails
	}
	return balanceValue, nil
}
