package main

import (
    "fmt"
    "math/rand"
    "github.com/fatih/color"
)

func GenerateSymbolArray(symbols map[string]uint) []string {
    symbolArr := []string{}
    for symbol, count := range symbols {
        for i := uint(0); i < count; i++ {
            symbolArr = append(symbolArr, symbol)
        }
    }
    return symbolArr
}

func getRandomNum(min int, max int) int {
    return rand.Intn(max-min+1) + min
}

func GetSpin(reel []string, rows int, cols int) [][]string {
    result := make([][]string, rows)
    for i := range result {
        result[i] = make([]string, 0)
    }

    for col := 0; col < cols; col++ {
        selected := map[int]bool{}
        for row := 0; row < rows; row++ {
            for {
                randomIndex := getRandomNum(0, len(reel)-1)
                if !selected[randomIndex] {
                    selected[randomIndex] = true
                    result[row] = append(result[row], reel[randomIndex])
                    break
                }
            }
        }
    }
    return result
}

func PrintSpin(spin [][]string) {
    fmt.Println()
    color.Magenta("     SPIN RESULT")
    color.Magenta("--------------------")
    for _, row := range spin {
        for j, symbol := range row {
            fmt.Printf(symbol)
            if j != len(row)-1 {
                fmt.Print(" | ")
            }
        }
        fmt.Println()
    }
    color.Magenta("--------------------")
}
