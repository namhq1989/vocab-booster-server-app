package dto

type ChangeExerciseFavoriteRequest struct {
	IsFavorite bool `json:"isFavorite"`
}

type ChangeExerciseFavoriteResponse struct {
	IsFavorite bool `json:"isFavorite"`
}
