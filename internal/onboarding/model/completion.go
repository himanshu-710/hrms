package model
type CompletionDTO struct {
	Sections   map[string]bool `json:"sections"`
	Percentage int             `json:"percentage"`
}