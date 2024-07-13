package dto

type GetExercisesRequest struct {
	Level string `query:"level"`
}

type GetExercisesResponse struct {
	Exercises []Exercise `json:"exercises"`
}
