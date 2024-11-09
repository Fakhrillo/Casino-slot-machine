# Fakha's Casino 🎰

Welcome to **Fakha's Casino**, a console-based slot machine game written in Go! Test your luck by placing bets and spinning the reels. Match symbols to win big and enjoy a stylish, colorful gaming experience in your terminal!

## Features

- **Bet and Win**: Place your bet and spin the reels to see if you can line up matching symbols.
- **Colorful Console Output**: Enhanced with colorful messages and banners to make the game experience visually engaging.
- **Randomized Spins**: Each spin is random, offering a fresh experience every time.
- **Customizable Symbols and Multipliers**: Easily adjust symbol values and multipliers in the code to create your own casino game variants.

## Gameplay

1. **Enter Your Name**: Start by entering your name to join the game.
2. **Place Your Bet**: Decide how much you'd like to bet for each spin. Bets can’t exceed your balance.
3. **Spin the Reels**: Watch the symbols roll in! Line up matching symbols to win according to the multiplier.
4. **Check Your Balance**: After each spin, your balance is updated based on any winnings.
5. **Quit Anytime**: Enter `0` as your bet to leave the game and cash out.

## Setup

### Prerequisites

- **Go**: Ensure Go is installed on your machine. [Download Go here](https://golang.org/dl/).
- **Fatih Color Package**: Used for colorful terminal output. Install it via:

    ```bash
    go get github.com/fatih/color
    ```

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/Fakhrillo/Casino-slot-machine.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Casino-slot-machine
    ```

3. Build the game:

    ```bash
    go build -o casino-game
    ```

4. Run the game:

    ```bash
    ./casino-game
    ```

## Project Structure

- **`main.go`**: Entry point of the game, managing the main loop, betting, and winning calculations.
- **`spin.go`**: Contains functions to generate the symbols, handle random selection for spins, and print the spin results.
- **`utils.go`**: Provides utility functions for user input, such as getting the player's name and handling bets.

## Example Gameplay

```plaintext
==============================
       Welcome to Fakha's Casino!
==============================

Enter your name: Alex
Welcome Alex, let's play!

Enter your bet, or 0 to quit (balance = $200): 10

     SPIN RESULT
--------------------
A | A | A
C | B | A
B | C | C
--------------------

Won $200, (20x) on Line #1
Current Balance: $380

Enter your bet, or 0 to quit (balance = $380): 0
Thank you for playing, Alex! You left with $380.


## Customization
You can modify symbol values and multipliers in main.go:

    ```bash
    symbols := map[string]uint{
        "A": 4,
        "B": 7,
        "C": 12,
        "D": 20,
    }
    multipliers := map[string]uint{
        "A": 20,
        "B": 10,
        "C": 5,
        "D": 2,
    }
    ```

Change these values to adjust the rarity of symbols and the payout multipliers to create your preferred game experience.