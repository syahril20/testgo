package dto

type IMSRequest struct {
	Email                string `json:"email" binding:"required,email"`
	Name                 string `json:"name" binding:"required"`
	Old                  int    `json:"old" binding:"required"`
	Phone                int    `json:"phone" binding:"required"`
	Address              string `json:"address" binding:"required"`
	OptiSampleCollection string `json:"opti_sample_collection" binding:"required"`
	CreatedBy            string `json:"created_by"`
	UpdatedBy            string `json:"updated_by"`
}

type IMSResponse struct {
	ID                   string `json:"id"`
	Email                string `json:"email"`
	Name                 string `json:"name"`
	Old                  int    `json:"old"`
	Phone                int    `json:"phone"`
	Address              string `json:"address"`
	OptiSampleCollection string `json:"opti_sample_collection"`
	CreatedAt            string `json:"created_at"`
	CreatedBy            string `json:"created_by"`
	UpdatedAt            string `json:"updated_at"`
	UpdatedBy            string `json:"updated_by"`
}
