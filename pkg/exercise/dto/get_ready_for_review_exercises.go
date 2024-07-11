package dto

type GetReadyForReviewExercisesRequest struct{}

type GetReadyForReviewExercisesResponse struct {
	Exercises []Exercise `json:"exercises"`
}
