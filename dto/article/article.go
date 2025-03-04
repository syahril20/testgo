package dto

type DetailArticle struct {
	Description   string `json:"description"`
	Ingredients   string `json:"ingredients"`
	Instructions  string `json:"instructions"`
	CountTime     string `json:"count_time"`
	CountCalories string `json:"count_calories"`
}

type NutritionItem struct {
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Value string `json:"value"`
}

type Article struct {
	ID          string          `json:"id,omitempty"`
	IDCategory  string          `json:"id_category"`
	Title       string          `json:"title"`
	Premium     bool            `json:"premium"`
	ReadingTime int             `json:"reading_time"`
	Image       string          `json:"image"`
	Detail      DetailArticle   `json:"detail"`
	Nutrition   []NutritionItem `json:"nutrition"`
	Index       string          `json:"index"`
	Deleted     bool            `json:"deleted"`
}
