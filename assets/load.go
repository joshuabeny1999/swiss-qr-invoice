package assets

import (
	_ "embed"
	"fmt"
)

//go:embed raw/scissors.png
var scissorsImage []byte

//go:embed raw/ch-cross.png
var chCrossImage []byte

func loadFromEmbed(data []byte, description string) ([]byte, error) {
	if data == nil {
		return nil, fmt.Errorf("could not load %s: file not found", description)
	}
	return data, nil
}

// Scissors returns the scissors image as bytes.
func Scissors() ([]byte, error) {
	return loadFromEmbed(scissorsImage, "scissors image")
}

// CHCross returns the CH-cross image as bytes.
func CHCross() ([]byte, error) {
	return loadFromEmbed(chCrossImage, "ch-cross image")
}
