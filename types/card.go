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
