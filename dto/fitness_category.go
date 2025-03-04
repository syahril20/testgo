package dto

type FitnessCategoryDTO struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Deleted     bool   `json:"deleted"`
}
