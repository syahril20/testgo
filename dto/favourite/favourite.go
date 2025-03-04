package dto

type FavouriteDTO struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ArticleID string `json:"article_id"`
	Count     int    `json:"count"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type FavouriteUpdateDTO struct {
	UserID    string `json:"user_id" binding:"required"`
	ArticleID string `json:"article_id" binding:"required"`
	UpdatedBy string `json:"updated_by" binding:"required"`
	Count     int    `json:"count"`
}
