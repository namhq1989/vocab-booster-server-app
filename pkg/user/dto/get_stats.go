package dto

type GetStatsRequest struct{}

type GetStatsResponse struct {
	Point                     int64 `json:"point"`
	CompletionTime            int   `json:"completionTime"`
	MasteredExercises         int   `json:"masteredExercises"`
	WaitingForReviewExercises int   `json:"waitingForReviewExercises"`
}
