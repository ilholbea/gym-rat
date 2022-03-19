package types

type Exercise struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Video       string `json:"video"`
	Image       string `json:"image"`
}
