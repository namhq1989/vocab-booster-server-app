package dto

import "github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"

type AnswerExerciseRequest struct {
	IsCorrect      bool  `json:"isCorrect"`
	CompletionTime int   `json:"completionTime"`
	Point          int64 `json:"point"`
}

type AnswerExerciseResponse struct {
	NextReviewAt *httprespond.TimeResponse `json:"nextReviewAt"`
}
