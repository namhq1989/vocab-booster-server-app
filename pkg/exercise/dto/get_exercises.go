package dto

type GetExercisesRequest struct {
	CollectionID string `query:"collectionId" validate:"required" message:"exercise_invalid_collection_id"`
}

type GetExercisesResponse struct {
	Exercises []Exercise `json:"exercises"`
}
