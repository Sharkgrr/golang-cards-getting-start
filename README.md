# Deck of Cards in Go

This is a simple Go application that simulates a deck of playing cards. It allows you to create a new deck, shuffle it, deal hands, save it to a file, and load it from a file. This program demonstrates the use of Go slices, basic file I/O operations, and randomization.

## Getting Started

### Prerequisites

To run this application, you need to have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

## How It Works

- **Deck**: A custom type defined as a slice of strings (`deck []string`) representing a collection of playing cards.
- **Shuffle**: Uses Go's `rand` package to randomly shuffle the deck.
- **File I/O**: Uses the `ioutil` package to save and read the deck from files.

## Features

- **Create a deck**: Generates a deck of cards with suits and values.
- **Print the deck**: Outputs the deck to the console.
- **Deal cards**: Splits the deck into hands.
- **Save to file**: Saves the current state of the deck to a file.
- **Load from file**: Loads a deck from a file.
- **Shuffle the deck**: Randomly shuffles the deck.
