package types

import (
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// Card is a basic container for a Magic the Gathering Card
// Struct fields are based off https://mtgjson.com/documentation.html
type Card struct {
	Names     []string `json:"name" validate:"required"`
	CMC       int8     `json:"cmc" validate:"required"`
	Colors    []string `json:"colors" validate:"required"`
	ManaCost  string   `json:"manaCost" validate:"required"`
	Type      string   `json:"type" validate:"required"`
	Text      string   `json:"text" validate:"required"`
	Power     string   `json:"power,omitempty" `
	Toughness string   `json:"toughness,omitempty"`
	Loyality  int8     `json:"loyality,omitempty"`
}

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func (c *Card) validate() error {
	return validate.Struct(c)
}

// Collection is simply an array of cards
type Collection []*Card

// ManaCurve returns the manacost of the collection of cards
// stored in memory
func (c *Collection) ManaCurve() map[int8]int8 {
	curve := map[int8]int8{}
	for _, card := range *c {
		curve[card.CMC]++
	}
	return curve
}

// CountType will return the number cards of given type
func (c *Collection) CountType(name string) int {
	count := 0
	for _, card := range *c {
		if strings.Contains(card.Type, name) {
			count++
		}
	}
	return count
}

// Colours returns what different colours make up the collection
func (c *Collection) Colours() []string {
	colours := map[string]bool{}
	for _, card := range *c {
		for _, c := range card.Colors {
			colours[c] = true
		}
	}
	ret := []string{}
	for k, _ := range colours {
		ret = append(ret, k)
	}
	return ret
}
