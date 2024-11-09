package main

import (
	"github.com/fatih/color"
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}
	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
		lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	// Define symbols and multipliers
	symbols := map[string]uint{
		"A": 4, "B": 7, "C": 12, "D": 20,
	}
	multipliers := map[string]uint{
		"A": 20, "B": 10, "C": 5, "D": 2,
	}

	// Display a styled welcome message
	printBanner("Welcome to Fakha's Casino!")

	symbolArr := GenerateSymbolArray(symbols)
	balance := uint(200)

	playerName := GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet

		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)

		winningLines := checkWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				color.Green("Won $%d, (%dx) on Line #%d", win, multi, i+1)
			}
		}
		color.Cyan("Current Balance: $%d", balance)
	}

	color.Yellow("Thank you for playing, %s! You left with $%d.\n", playerName, balance)
}
