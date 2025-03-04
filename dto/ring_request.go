package dto

type CreateRingRequest struct {
	Size       int    `json:"size" binding:"required"`
	Color      string `json:"color" binding:"required"`
	Connection bool   `json:"connection" binding:"required"` // Field baru
}

type UpdateRingRequest struct {
	Size       int    `json:"size" binding:"required"`
	Color      string `json:"color" binding:"required"`
	Connection bool   `json:"connection" binding:"required"` // Field baru
}
