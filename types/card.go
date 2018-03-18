package types

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

// Collection ...
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
		_ = card
	}
	return count
}
