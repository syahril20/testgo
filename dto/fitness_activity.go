package dto

type FitnessActivityDTO struct {
	ID        string `json:"id,omitempty"`
	IDUser    string `json:"id_user"`
	IDFitness string `json:"id_fitness"`
	Finished  bool   `json:"finished"`
}
