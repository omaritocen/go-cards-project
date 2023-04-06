package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLen := 16
	if len(d) != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first element to be %v, but got %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "Four of Clubs"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last element to be %v, but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	// Initial Cleanup
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	// Read the deck from the saved file
	loadedDeck := newDeckFromFile(filename)

	expectedLen := 16
	if len(loadedDeck) != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, len(d))
	}

	// Cleanup after testing
	os.Remove(filename)
}
