package dto

type GetExerciseCollectionsRequest struct{}

type GetExerciseCollectionResponse struct {
	Collections []ExerciseCollection `json:"collections"`
}
