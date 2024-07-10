package dto

type GetExercisesRequest struct{}

type GetExercisesResponse struct {
	Exercises []Exercise `json:"exercises"`
}
