package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalance(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("file not found.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("file contains invalid information")
	}
	return balance, nil
}

func SaveBalance(balance float64, fileName string) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceString), 0644)
}
