package main

import (
    "fmt"
    "github.com/fatih/color"
)

func GetName() string {
    var name string
    fmt.Println()
    color.Yellow("Welcome to Fakha's Casino...")
    fmt.Printf("Enter your name: ")
    _, err := fmt.Scanln(&name)

    if err != nil {
        fmt.Println("Error reading name:", err)
        return ""
    }

    color.Cyan("Welcome %s, let's play!\n", name)
    return name
}

func GetBet(balance uint) uint {
    var bet uint
    for {
        color.Blue("Enter your bet, or 0 to quit (balance = $%d): ", balance)
        fmt.Scan(&bet)

        if bet > balance {
            color.Red("Bet cannot be larger than balance!")
        } else {
            break
        }
    }
    return bet
}

func printBanner(message string) {
    fmt.Println()
    color.Magenta("==============================")
    color.Magenta("       %s", message)
    color.Magenta("==============================")
    fmt.Println()
}
